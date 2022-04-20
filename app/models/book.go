package models

import (
	"time"

	"github.com/google/uuid"
)

type ModelV1 struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	ModelV1
	Title       string
	Author      string
	PublishDate time.Time
}
