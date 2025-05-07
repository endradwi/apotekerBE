package routers

import (
	"apotekerBE/controllers"
	"apotekerBE/middlewares"

	"github.com/gin-gonic/gin"
)

func ReserveRouter(router *gin.RouterGroup) {
	router.POST("", middlewares.ValidationToken(), controllers.CreateData)
	// router.GET("", middlewares.ValidationToken(), controllers.GetAllReserve)
	// router.GET("/:id", middlewares.ValidationToken(), controllers.GetOneReserve)
	// router.PATCH("/:id", middlewares.ValidationToken(), controllers.EditReserve)
	// router.DELETE("/:id", middlewares.ValidationToken(), controllers.DeleteReserve)
}
