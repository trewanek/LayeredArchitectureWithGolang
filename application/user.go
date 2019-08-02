package application

import (
	"context"
	"github.com/trewanek/LayeredArchitectureWithGolang/domain"
	"golang.org/x/xerrors"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := domain.GetUserByID(ctx, userID)
	if err != nil {
		return nil, xerrors.Errorf("error in package application#GetUserByID: %w", err)
	}
	return user, nil
}
