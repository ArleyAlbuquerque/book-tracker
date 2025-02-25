package schemas

import (
	"gorm.io/gorm"
)

type Creating struct {
	gorm.Model
	Id    int64
	Name  string
	Autor string
	Pages int64
}
