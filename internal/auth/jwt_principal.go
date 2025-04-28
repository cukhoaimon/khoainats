package auth

import (
	"github.com/google/uuid"
)

type JwtPrincipal struct {
	OrgId     uuid.UUID
	UserAgent string
	roles     []PrincipalRoleType
}
