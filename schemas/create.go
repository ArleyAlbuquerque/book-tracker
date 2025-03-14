package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Creating struct {
	gorm.Model
	Name  string
	Autor string
	Pages int64
}

type CreatingResponses struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Autor     string    `json:"autor"`
	Pages     int64     `json:"pages"`
}
