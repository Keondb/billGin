package api

import (
	"bill/bill_serve/api/images_api"
	"bill/bill_serve/api/settings_api"
	"bill/bill_serve/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	UserApi     user_api.UserApi
	ImagesApi   images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
