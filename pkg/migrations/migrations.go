package migrations

import (
	"e-invoicing/pkg/models"
)

func AuthMigrationModels() []interface{} {
	return []interface{}{
		&models.Business{},
		&models.Invoice{},
		&models.AccessToken{},
		&models.TokenManager{},
	}

}

func AlterColumnModels() []AlterColumn {
	return []AlterColumn{}
}
