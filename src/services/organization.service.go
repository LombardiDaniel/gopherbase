package services

import (
	"context"

	"github.com/LombardiDaniel/go-gin-template/models"
)

type OrganizationService interface {
	GetOrganization(ctx context.Context, orgId string) (models.Organization, error)
	CreateOrganization(ctx context.Context, org models.Organization) error
	CreateOrganizationInvite(ctx context.Context, invite models.OrganizationInvite) error
}