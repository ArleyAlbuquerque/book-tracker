package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	//Initialize DB(SQLite)
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("erro to initialize sqlite: %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Create a New Logger
	logger = NewLogger(p)
	return logger
}
