package vk

import (
	"context"
	"fmt"

	"github.com/sealoftime/adapteris/domain/user"
	"golang.org/x/oauth2"
)

func (vk *Service) GetUserDefaultsForToken(
	ctx context.Context,
	token *oauth2.Token,
) (user.Account, error) {
	return vk.extractUserDefaults(ctx, token)
}

func (vk *Service) extractUserDefaults(
	ctx context.Context,
	t *oauth2.Token,
) (user.Account, error) {
	vkUser, err := vk.GetUserProfile(ctx, t.AccessToken, []string{"domain"})
	if err != nil {
		return user.Account{}, err
	}

	fullName := fmt.Sprintf("%s %s", vkUser.FirstName, vkUser.LastName)
	return user.Account{
		FullName:  &fullName,
		ShortName: vkUser.FirstName,
		Email:     t.Extra("email").(string),
		Vk:        vkUser.Domain,
	}, nil
}
