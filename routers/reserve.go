package routers

import (
	"apotekerBE/controllers"
	"apotekerBE/middlewares"

	"github.com/gin-gonic/gin"
)

func ReserveRouter(router *gin.RouterGroup) {
	router.POST("", middlewares.ValidationToken(), controllers.CreateData)
	router.GET("/all/reserve/admin", controllers.GetAllReserveAdmin)
	router.GET("/all/reserve/users", middlewares.ValidationToken(), controllers.GetAllReserve)
	// router.PATCH("/:id", middlewares.ValidationToken(), controllers.EditReserve)
	// router.DELETE("/:id", middlewares.ValidationToken(), controllers.DeleteReserve)
}
