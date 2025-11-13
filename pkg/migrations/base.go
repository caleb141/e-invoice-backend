package migrations

import (
	"e-invoicing/pkg/database"
	"fmt"

	"gorm.io/gorm"
)

func RunAllMigrations(db *database.Database) {

	MigrateModels(db.Postgresql.DB(), AuthMigrationModels(), AlterColumnModels())

}

func MigrateModels(db *gorm.DB, models []interface{}, AlterColums []AlterColumn) {
	_ = db.AutoMigrate(models...)

	for _, d := range AlterColums {
		err := d.UpdateColumnType(db)
		if err != nil {
			fmt.Println("error migrating ", d.TableName, "for column", d.Column, ": ", err)
		}

	}

}
