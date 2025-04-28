package model

import (
	"time"

	"github.com/google/uuid"
)

type LoginEvent struct {
	Id          uuid.UUID
	IpAddress   string
	Code        string
	Attempts    int
	CreatedAt   time.Time
	SucceedAt   time.Time
	LastTriedAt time.Time
}
