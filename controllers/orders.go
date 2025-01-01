package controllers

import (
	"fmt"
	"net/http"
	"test/models"

	"github.com/gin-gonic/gin"
)

// Add Movie godoc
// @Schemes
// @Description Add Order tikers
// @Tags Order Tiket
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Edit_Movie formData models.OrderBody true "Add Order Tiket"
// @Success 200 {object} Response{results=models.OrderBody}
// @Router /orders [post]

func OrderMovies(ctx *gin.Context) {

	var orderMovie models.OrderBody
	err := ctx.ShouldBind(&orderMovie)

	if err != nil {
		fmt.Println(err)
		return
	}

	models.OrderTicket(orderMovie)
	// log.Println("data apa =", order)
	var orderPayment models.Payment

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Order tiket sukses",
		Results: orderPayment,
	})

}
