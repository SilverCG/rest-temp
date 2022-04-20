package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelV1 struct {
	ID        uuid.UUID `gorm:"primary_key;unique;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TODO do this for now until we add the postgress function for uuid_generate_v4()
func (m *ModelV1) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

type Book struct {
	ModelV1
	Title       string
	Author      string
	PublishDate time.Time
}
