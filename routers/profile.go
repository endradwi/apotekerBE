package routers

import (
	"test/controllers"
	"test/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidationToken())
	router.PATCH("/:id", controllers.EditProfile)
	// router.PATCH("/:id", controllers.EditUser)
	// router.GET("", controllers.GetAllProfile)
	router.GET("", controllers.GetProfile)

}
