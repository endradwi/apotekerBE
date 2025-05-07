package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		// AllowAllOrigins: true,
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowCredentials: true,
	}))
	AuthRouters(router.Group("/auth"))
	DoctorRouter(router.Group("/doctor"))
	UserRouter(router.Group("/users"))
	ReserveRouter(router.Group("/reserve"))
}
