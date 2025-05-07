package controllers

import (
	"apotekerBE/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDoctor(ctx *gin.Context) {
	doctors, err := models.GetDoctor()
	if err != nil {
		fmt.Println("Error Get All Doctor", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get doctors"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Get All Doctor",
		"results": doctors,
	})
}
