package vk

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sealoftime/adapteris/domain/user"
	"golang.org/x/oauth2"
)

//GetLoginUrl returns the login url to initiate the OAuth2 Access Code authentication flow.
//
//redirectAfter - is a url, that the user must be redirected after the flow is complete.
func (s *Service) GetLoginUrl(ctx context.Context, redirectAfter string) string {
	return s.config.AuthCodeURL(redirectAfter)
}

//ExchangeCode exchanges access code received on a previous step of access code authentication flow
//for an access token.
func (s *Service) exchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return s.config.Exchange(ctx, code)
}

func (s *Service) ExtractExternalUserId(ctx context.Context, token *oauth2.Token) (string, error) {
	exIdRaw := token.Extra("user_id")
	if exIdRaw == nil {
		return "", fmt.Errorf("no user_id extra in VK token")
	}

	exIdInt := int64(exIdRaw.(float64))
	exId := strconv.FormatInt(exIdInt, 10)
	return exId, nil
}

//LoginByAccessCode performs neccessary communication with Vk services to authenticate either
//already signed up user or to create a new account, based on their Vk data.
func (s *Service) LoginByAccessCode(ctx context.Context, code string) (*user.Account, error) {
	token, err := s.exchangeCode(ctx, code)
	if err != nil {
		return nil, ErrExchangeFailed{Cause: err}
	}
	eid, err := s.ExtractExternalUserId(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("error authenticating by VK: %w", err)
	}

	u, err := s.loginByToken(ctx, eid, token)
	if err == nil {
		return u, nil
	}
	if !user.IsNotFoundByExternalAccount(err) {
		return nil, err
	}

	u, err = s.registerByToken(ctx, eid, token)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) loginByToken(ctx context.Context, eid string, t *oauth2.Token) (*user.Account, error) {
	u, err := s.auth.LoginByExternalAccount(ctx, user.ExternalAccount{
		Service:    "vk",
		ExternalId: eid,
	})
	if err != nil {
		return nil, fmt.Errorf("error logging in with VK: %w", err)
	}

	return u, nil
}

func (s *Service) registerByToken(ctx context.Context, eid string, t *oauth2.Token) (*user.Account, error) {
	defaults, err := s.GetUserDefaultsForToken(ctx, t)
	if err != nil {
		return nil, fmt.Errorf("error registering by Vk: %w", err)
	}

	u, err := s.auth.RegisterUserByExternalAccount(ctx, defaults, user.ExternalAccount{
		Service:    "vk",
		ExternalId: eid,
	})
	if err != nil {
		return nil, fmt.Errorf("error registering by Vk: %w", err)
	}

	return u, nil
}
