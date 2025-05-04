package model

import (
	"time"

	"github.com/google/uuid"
)

type AccessToken struct {
	Id             uuid.UUID
	PrincipalId    uuid.UUID
	OrganizationId uuid.UUID
	Roles          []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	RevokedAt      time.Time
	RevokedBy      uuid.UUID
}
