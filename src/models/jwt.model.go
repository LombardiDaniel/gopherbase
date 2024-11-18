package models

import (
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	UserId         uint32  `json:"userId" binding:"required"`
	Email          string  `json:"email" binding:"required"`
	OrganizationId *string `json:"organizationId" binding:"required"`

	jwt.StandardClaims
}

// only here because swaggo cant expand the above example (but same thing, KEEP IN SYNC!!)
type JwtClaimsOutput struct {
	UserId         uint32  `json:"userId" binding:"required"`
	Email          string  `json:"email" binding:"required"`
	OrganizationId *string `json:"organizationId" binding:"required"`

	Audience  string `json:"aud"`
	ExpiresAt int64  `json:"exp"`
	Id        string `json:"jti"`
	IssuedAt  int64  `json:"iat"`
	Issuer    string `json:"iss"`
	NotBefore int64  `json:"nbf"`
	Subject   string `json:"sub"`
}
