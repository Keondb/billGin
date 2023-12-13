package settings_api

import (
	"bill/bill_serve/global"
	"bill/bill_serve/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {
	res.OkWithData(global.Config.Email, c)
}
