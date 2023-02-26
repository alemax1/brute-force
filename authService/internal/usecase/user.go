package usecase

import (
	"authService/internal/domain"
	"authService/internal/repository"
	"authService/internal/usecase/helper"
	"context"
	"fmt"
)

type user struct {
	uRepo repository.User
}

type User interface {
	CreateUser(ctx context.Context, user domain.User) error
}

func New(uRepo repository.User) User {
	return &user{
		uRepo: uRepo,
	}
}

func (u user) CreateUser(ctx context.Context, user domain.User) error {
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("hash: %v", err)
	}

	user.Password = hashedPassword

	if err := u.uRepo.SaveUser(ctx, user); err != nil {
		return fmt.Errorf("save: %v", err)
	}

	return nil
}
