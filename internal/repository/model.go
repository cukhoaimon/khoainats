package repository

import (
	"time"

	"github.com/cukhoaimon/khoainats/internal/auth"
	"github.com/google/uuid"
)

type AccessToken struct {
	Id             uuid.UUID
	PrincipalId    uuid.UUID
	OrganizationId uuid.UUID
	Roles          []auth.PrincipalRoleType
	CreatedAt      time.Time
	UpdatedAt      time.Time
	RevokedAt      time.Time
	RevokedBy      uuid.UUID
}

type LoginEvent struct {
	Id          uuid.UUID
	IpAddress   string
	Code        string
	Attempts    int
	CreatedAt   time.Time
	SucceedAt   time.Time
	LastTriedAt time.Time
}

type Organization struct {
	Id         uuid.UUID
	OwnerEmail string
	Domain     string
	CreatedAt  time.Time
	CreatedBy  uuid.UUID
	UpdatedAt  time.Time
	UpdatedBy  uuid.UUID
	DeletedAt  time.Time
	DeletedBy  uuid.UUID
}

type Principal struct {
	Id             uuid.UUID
	Type           PrincipalType
	OrganizationId uuid.UUID
	CreatedAt      time.Time
	DeletedAt      time.Time
	DeletedBy      uuid.UUID
}

type PrincipalAttribute struct {
	Id             uuid.UUID
	PrincipalId    uuid.UUID
	Attribute      PrincipalAttributeType
	AttributeValue string
	CreatedAt      time.Time
	CreatedBy      uuid.UUID
	DeletedAt      time.Time
	DeletedBy      time.Time
}
