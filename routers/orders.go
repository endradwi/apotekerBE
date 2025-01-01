package routers

import (
	"test/controllers"

	"github.com/gin-gonic/gin"
)

func OrdersRouter(router *gin.RouterGroup) {
	router.GET("", controllers.GetAllMovies)
	router.GET("/:id", controllers.GetDetailMoviesById)
	router.POST("", controllers.OrderMovies)
}
