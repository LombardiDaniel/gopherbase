package services

import (
	"context"

	"github.com/LombardiDaniel/go-gin-template/models"
	"github.com/LombardiDaniel/go-gin-template/schemas"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
	CreateUnconfirmedUser(ctx context.Context, unconfirmedUser models.UnconfirmedUser) error
	ConfirmUser(ctx context.Context, otp string) error
	GetUser(ctx context.Context, email string) (models.User, error)
	GetUserFromId(ctx context.Context, id uint32) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserOrgs(ctx context.Context, userId uint32) ([]schemas.OrganizationOutput, error)
}
