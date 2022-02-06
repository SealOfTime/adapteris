package domain

import (
	"context"
	"time"
)

const (
	USER  UserRole = "USER"
	ADMIN UserRole = "ADMIN"
)

type UserRole string

type User struct {
	Id           int64
	RegisteredAt time.Time
	Role         UserRole

	FullName  *string
	ShortName string

	Email       string
	Telegram    *string
	Vk          *string
	PhoneNumber *string
}

type UserStorage interface {
	FindById(ctx context.Context, id int64) (*User, error)
	UpsertByEmail(ctx context.Context, user User) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
}
