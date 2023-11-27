package settings_api

import (
	"bill_serve/config"
	"bill_serve/core"
	"bill_serve/global"
	"bill_serve/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsEmailInfoUpdate(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}

	global.Config.Email = cr

	err = core.Setyaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
