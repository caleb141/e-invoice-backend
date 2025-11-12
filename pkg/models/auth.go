package models

type CreateUserRequestModel struct {
	Name            string              `json:"name" validate:"required,min=2,max=250"`
	Email           string              `json:"email" validate:"required,email"`
	Password        string              `json:"password" validate:"required,min=6"`
	PlatformConfigs PlatformConfigsAuth `json:"platform_configs" validate:"dive"`
}

type UpdateUserRequestModel struct {
	Name string `json:"name" validate:"required"`
}

type LoginRequestModel struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type PlatformConfigsAuth map[string]AccountingPlatformConfigAuth
type AccountingPlatformConfigAuth struct {
	OrgID      string `json:"org_id"`
	AuthToken  string `json:"auth_token"`
	HMACSecret string `json:"hmac_secret"`
	APIKey     string `json:"api_key"`
	APISecret  string `json:"api_secret"`
}
