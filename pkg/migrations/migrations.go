package migrations

import (
	"einvoice-access-point/pkg/models"
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
