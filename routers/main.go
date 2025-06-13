package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://apoteker-fe-production.up.railway.app"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
			"User_ID", // <- tambahkan custom header ini
		},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://apoteker-fe-production.up.railway.app"
		},
	}))

	AuthRouters(router.Group("/auth"))
	DoctorRouter(router.Group("/doctor"))
	UserRouter(router.Group("/users"))
	ReserveRouter(router.Group("/reserve"))
}
