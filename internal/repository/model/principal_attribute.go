package model

import (
	"time"

	"github.com/google/uuid"
)

type PrincipalAttributeType int

const (
	Email PrincipalAttributeType = iota
	MacAddress
)

var principalAttributes = map[PrincipalAttributeType]string{
	Email:      "email",
	MacAddress: "mac_address",
}

func (p PrincipalAttributeType) String() string {
	return principalAttributes[p]
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
