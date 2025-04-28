package model

import (
	"time"

	"github.com/google/uuid"
)

type PrincipalType int

const (
	USER PrincipalType = iota
	SERVICE
)

var principalType = map[PrincipalType]string{
	USER:    "user",
	SERVICE: "service",
}

func (p PrincipalType) String() string {
	return principalType[p]
}

type Principal struct {
	Id             uuid.UUID
	Type           PrincipalType
	OrganizationId uuid.UUID
	CreatedAt      time.Time
	DeletedAt      time.Time
	DeletedBy      uuid.UUID
}
