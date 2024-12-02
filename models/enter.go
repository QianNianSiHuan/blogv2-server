package models

import (
	"time"
)

type Model struct {
	UUID      uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
