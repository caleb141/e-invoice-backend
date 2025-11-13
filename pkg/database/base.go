package database

import (
	"e-invoicing/pkg/config"
	"e-invoicing/pkg/utility"

	"gorm.io/gorm"
)

type DbConnection interface {
	NewDatabaseConnection(db *gorm.DB, logger *utility.Logger, config *config.Database) *Database
}

type Database struct {
	Postgresql DatabaseManager
	Redis      CacheManager
}

var DB *Database = &Database{}

func Connection() *Database {
	return DB
}
