package memory

import (
	"context"
	"fmt"
	"time"

	domain "github.com/sealoftime/adapteris"
)

type UserStore struct {
	lastId  int64
	storage []domain.User
}

func (s *UserStore) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	user.Id = s.lastId
	s.lastId++
	s.storage = append(s.storage, user)
	return &user, nil
}

func (s *UserStore) FindById(ctx context.Context, id int64) (*domain.User, error) {
	for _, u := range s.storage {
		if u.Id == id {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("Not found") //TODO: User not found by id error
}

func (s *UserStore) UpsertByEmail(ctx context.Context, user domain.User) (*domain.User, error) {
	for _, u := range s.storage {
		if u.Email == user.Email {
			u.FullName = user.FullName
			u.ShortName = user.ShortName
			u.Vk = user.Vk
			return &u, nil
		}
	}

	newUser := domain.User{
		Id:           s.lastId,
		RegisteredAt: time.Now(),
		Role:         "USER",
		FullName:     &*user.FullName,
		ShortName:    user.ShortName,
		Email:        user.Email,
		Vk:           &*user.Vk,
	}

	fmt.Println(newUser)
	s.storage = append(s.storage, newUser)
	return &newUser, nil
}
