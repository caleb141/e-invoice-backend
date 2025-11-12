package token

import (
	"einvoice-access-point/pkg/database"
	"einvoice-access-point/pkg/models"
	"time"
)

func FindToken(db database.DatabaseManager, provider, orgID string) (*models.TokenManager, error) {
	var token models.TokenManager
	if err := db.DB().Where("provider = ? AND organization_id = ?", provider, orgID).
		First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func SaveNewZohoToken(db database.DatabaseManager, provider, orgID, accessToken, refreshToken string, expiresIn int) error {
	zohoToken := models.TokenManager{
		Provider:       provider,
		OrganizationID: orgID,
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		ExpiresAt:      time.Now().Add(time.Duration(expiresIn) * time.Second),
	}

	return db.DB().Create(&zohoToken).Error
}

func UpdateZohoToken(db database.DatabaseManager, provider, orgID, accessToken, refreshToken string, expiresIn int) error {
	var zohoToken models.TokenManager
	if err := db.DB().Where("provider = ? AND organization_id = ?", provider, orgID).First(&zohoToken).Error; err != nil {
		return err
	}

	zohoToken.AccessToken = accessToken
	zohoToken.RefreshToken = refreshToken
	zohoToken.ExpiresAt = time.Now().Add(time.Duration(expiresIn) * time.Second)

	return db.DB().Save(&zohoToken).Error
}
