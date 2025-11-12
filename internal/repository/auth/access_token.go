package auth

import (
	"einvoice-access-point/pkg/database"
	"einvoice-access-point/pkg/models"
	"fmt"
	"net/http"
)

func GetAccessTokens(a *models.AccessToken, db database.DatabaseManager) error {
	err := db.SelectFirstFromDb(&a)
	if err != nil {
		return fmt.Errorf("token selection failed: %v", err.Error())
	}
	return nil
}

func GetByOwnerID(a *models.AccessToken, db database.DatabaseManager) (int, error) {
	err, nilErr := db.SelectOneFromDb(db, &a, "owner_id = ? ", a.OwnerID)
	if nilErr != nil {
		return http.StatusBadRequest, nilErr
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func GetByID(ID string, db database.DatabaseManager) (models.AccessToken, error) {
	var accessT models.AccessToken

	query := db.DB().Where("id = ?", ID)

	if err := query.First(&accessT).Error; err != nil {
		return accessT, err
	}

	return accessT, nil
}

func GetByIDBoolean(a *models.AccessToken, db database.DatabaseManager) (int, error) {
	err, nilErr := db.SelectOneFromDb(&a, "id = ? ", a.ID)
	if nilErr != nil {
		return http.StatusBadRequest, nilErr
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func GetLatestByOwnerIDAndIsLive(a *models.AccessToken, db database.DatabaseManager) (int, error) {
	err, nilErr := db.SelectLatestFromDb(&a, "owner_id = ? and is_live = ? ", a.OwnerID, a.IsLive)
	if nilErr != nil {
		return http.StatusBadRequest, nilErr
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func CreateAccessToken(a *models.AccessToken, db database.DatabaseManager, tokenData interface{}) error {
	if a.OwnerID == "" {
		return fmt.Errorf("owner id not provided to create access token")
	}

	if a.ID == "" {
		return fmt.Errorf("access id not provided to create access token")
	}

	var (
		access_token = tokenData.(map[string]string)["access_token"]
		exp          = tokenData.(map[string]string)["exp"]
	)

	a.IsLive = true
	a.LoginAccessToken = access_token
	a.LoginAccessTokenExpiresIn = exp
	err := db.CreateOneRecord(&a)
	if err != nil {
		return fmt.Errorf("user creation failed: %v", err.Error())
	}
	return nil
}

func RevokeAccessToken(a *models.AccessToken, db database.DatabaseManager) error {
	if a.ID == "" {
		return fmt.Errorf("access token id not provided to revoke access token")
	}
	a.IsLive = false
	_, err := db.SaveAllFields(&a)
	return err
}
