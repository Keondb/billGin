package routers

import "bill_serve/api"

func (router RouterGroup) ImagesRouter() {
	ImagesApi := api.ApiGroupApp.ImagesApi
	router.POST("images", ImagesApi.ImageUploadView)
	router.GET("images", ImagesApi.ImageListView)
	router.DELETE("images", ImagesApi.ImageRemoveView)
}
