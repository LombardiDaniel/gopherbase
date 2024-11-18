package services

import (
	"context"

	"github.com/LombardiDaniel/go-gin-template/models"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
	CreateUnconfirmedUser(ctx context.Context, unconfirmedUser models.UnconfirmedUser) error
	ConfirmUser(ctx context.Context, otp string) error
	GetUser(ctx context.Context, email string) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}
