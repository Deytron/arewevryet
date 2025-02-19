package models

type Account struct {
	Name string`form:"Name" json:"Name"`
	Password string `form:"Password" json:"Password"`
}

type UserConfig struct {
	Id string `json:"Id"`
	EnableAllFolders bool `json:"EnableAllFolders"`
	EnabledFolders []string `json:"EnabledFolders"`
	AuthenticationProviderId string `json:"AuthenticationProviderId"`
	PasswordResetProviderId string `json:"PasswordResetProviderId"`
}