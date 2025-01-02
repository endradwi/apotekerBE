package routers

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)

func OrdersRouter(router *gin.RouterGroup) {
	// router.GET("", controllers.GetAllMovies)
	router.GET("/cinema/:id", controllers.GetMovieCinema)
	router.POST("", controllers.OrderMovies)
}
