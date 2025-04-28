package model

import (
	"time"

	"github.com/google/uuid"
)

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
