package images_api

import (
	"bill/bill_serve/models"
	"bill/bill_serve/models/res"
	"bill/bill_serve/service/common"

	"github.com/gin-gonic/gin"
)

// ImageListView 图片列表
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, total, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})

	res.OkWithList(list, total, c)
	return
}
