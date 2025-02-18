package schemas

import (
	"gorm.io/gorm"
)

type Creating struct {
	gorm.Model
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Autor string `json:"string"`
	Pages int    `json:"pages"`
}
