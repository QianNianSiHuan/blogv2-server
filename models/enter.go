package models

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type IDRequest struct {
	ID uint `json:"id" uri:"id"`
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
