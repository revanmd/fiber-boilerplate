package database

import (
	"api-ipms/config"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Connect() {
	connectionString := "sqlserver://" +
		config.Cfg.DatabaseUser + ":" +
		config.Cfg.DatabasePassword + "@" +
		config.Cfg.DatabaseHost + ":" +
		config.Cfg.DatabasePort + "?" + "database=" +
		config.Cfg.DatabaseName

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Error connected to GORM connection Database")
	}

	DB = db
}
