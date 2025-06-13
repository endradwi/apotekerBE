package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "*"},
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
	}))

	AuthRouters(router.Group("/auth"))
	DoctorRouter(router.Group("/doctor"))
	UserRouter(router.Group("/users"))
	ReserveRouter(router.Group("/reserve"))
}
