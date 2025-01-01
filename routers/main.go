package routers

import (
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {

	ProfileRouter(router.Group("/profile"))
	AuthRouters(router.Group("/auth"))
	MovieRouter(router.Group("/movies"))
	OrdersRouter(router.Group("/orders"))
}
