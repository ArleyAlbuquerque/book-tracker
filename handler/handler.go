package handler

import (
	"github.com/ArleyAlbuquerque/book-tracker/config"
	"gorm.io/gorm"
)

// Global Variables package handler

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
