package business

import (
	"e-invoicing/pkg/database"
	"e-invoicing/pkg/models"
	"e-invoicing/pkg/utility"

	"gorm.io/gorm"
)

func CreateBusiness(b *models.Business, db database.DatabaseManager) error {

	err := db.CreateOneRecord(&b)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAUser(b *models.Business, db database.DatabaseManager) error {
	_, err := db.SaveAllFields(&b)
	return err
}

func DeleteAUser(b *models.Business, db database.DatabaseManager) error {

	err := db.DeleteRecordFromDb(&b)

	if err != nil {
		return err
	}

	return nil
}

func GenerateUniqueServiceID(db *gorm.DB) string {
	var existingCount int64
	serviceID := utility.GenerateRandomServiceID()

	for {
		db.Table("businesses").Where("service_id = ?", serviceID).Count(&existingCount)
		if existingCount == 0 {
			break
		}
		serviceID = utility.GenerateRandomServiceID()
	}

	return serviceID
}
