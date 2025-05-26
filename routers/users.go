package routers

import (
	"apotekerBE/controllers"
	"apotekerBE/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	// router.Use(middlewares.ValidationToken())
	router.PATCH("", middlewares.ValidationToken(), controllers.EditProfile)
	router.POST("admin", middlewares.ValidationToken(), controllers.AddAdmin)
	router.PATCH("/:id", controllers.EditRoleUser)
	router.DELETE("/:id", controllers.DeleteUser)
	// router.PATCH("/:id", controllers.EditUser)
	router.GET("", middlewares.ValidationToken(), controllers.GetUser)
	router.GET("all", controllers.GetAllUser)
}
