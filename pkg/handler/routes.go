package handler

import "github.com/gin-gonic/gin"

func (h *handler) AuthenUserRoutes(route *gin.RouterGroup) {
	// route.POST("/:username/nickname/edit", routes.ChangeNickname)
	// route.GET("/nickname/edit", routes.EditNickname)
	route.GET("/", h.login)
	// route.PUT("/picture", routes.ChangePicture)

}

func (h *handler) ProfileUserRoutes(route *gin.RouterGroup) {
	route.GET("/", h.getUser)
	route.POST("/editnickname", h.editNickname)
	route.POST("/editpicture", h.editPicture)
}
