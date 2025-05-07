package routers

import "github.com/gin-gonic/gin"

func Routers(router *gin.Engine) {
	AuthRouters(router.Group("/auth"))
	DoctorRouter(router.Group("/doctor"))
	UserRouter(router.Group("/users"))
	ReserveRouter(router.Group("/reserve"))
}
