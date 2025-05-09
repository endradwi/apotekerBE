package main

import (
	"apotekerBE/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.Routers(router)
	router.Static("/profile/image", "./upload/profile")
	router.Run(":8889")
}
