package config

import (
	"os"

	"github.com/ArleyAlbuquerque/book-tracker/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	//Check if the DB file exists
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//Create DB & connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrF("sqlite opening error:%v", err)
		return nil, err
	}
	//Migrate the schemas to creating struct
	err = db.AutoMigrate(&schemas.Creating{})
	if err != nil {
		logger.ErrF("sqlite auto migration error:%v", err)
		return nil, err
	}
	// Return DB
	return db, nil
}
