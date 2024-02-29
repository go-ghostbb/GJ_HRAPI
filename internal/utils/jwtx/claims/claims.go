package claims

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"hrapi/internal/types/enum"
)

type ClaimType string

const (
	Access  ClaimType = "access"
	Refresh ClaimType = "refresh"
)

type AccessClaims struct {
	Type enum.MenuShow
	Key  uuid.UUID
	jwt.RegisteredClaims
	BaseClaims
}

type BaseClaims struct {
	EmployeeID uint
	RealName   string
	Account    string
	Roles      []string
}

type RefreshClaims struct {
	Key uuid.UUID
	jwt.RegisteredClaims
}
