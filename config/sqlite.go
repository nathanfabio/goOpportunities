package config

import (
	"os"

	"github.com/nathanfabio/goOpportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	//Check if db exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("db does not exist, creating...")
		//Create db
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

	//Initialize SQLite
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite open error: %v", err)
		return nil, err
	}

	//Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite migrate error: %v", err)
		return nil, err
	}

	//Return db
	return db, nil
}