package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"test/models"

	"github.com/gin-gonic/gin"
)

// Orders godoc
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

// Orders godoc
// @Schemes
// @Description Add Choose Cinema tikers
// @Tags Order Tiket
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path int true "Select Movie Tiket"
// @Param searchName query string true "Search Name Cinema"
// @Param searchDate query string true "Search Date Cinema"
// @Param searchTime query string true "Search Time Cinema"
// @Param searchLocation query string true "Search Location Cinema"
// @Success 200 {object} Response{results=models.MoviesCinema}
// @Router /orders/cinema/{id} [post]
func GetMovieCinema(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.DefaultQuery("id", ""))
	searchName := ctx.DefaultQuery("searchName", "")
	searchDate := ctx.DefaultQuery("searchDate", "")
	searchTime := ctx.DefaultQuery("searchTime", "")
	searchLocation := ctx.DefaultQuery("searchLocation", "")

	// var cinema models.ListCinema

	cinema := models.BookingCinema(id, searchName, searchDate, searchTime, searchLocation)
	log.Print(cinema)
	// log.Println(find.Cinema_time)

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Cinema Selected And Choose Yout seat",
		Results: cinema,
	})

}
