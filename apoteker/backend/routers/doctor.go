package routers

import (
	"apotekerBE/controllers"

	"github.com/gin-gonic/gin"
)

func DoctorRouter(router *gin.RouterGroup) {

	router.GET("", controllers.GetAllDoctor)

}
