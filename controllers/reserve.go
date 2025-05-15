package controllers

import (
	"apotekerBE/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateData(ctx *gin.Context) {
	val, isAvail := ctx.Get("userId")
	if !isAvail {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "Unauthorized id",
		})
		return
	}
	var form models.ReserveData
	err := ctx.ShouldBind(&form)
	fmt.Println("Content-Type:", ctx.ContentType())
	fmt.Println("form data", form)
	fmt.Println("date:", err)
	fmt.Println("date=", form.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Data",
		})
		return
	}
	form.User_id = val.(int)
	data := models.ReserveData{
		Fullname:     form.Fullname,
		Phone_number: form.Phone_number,
		Age:          form.Age,
		Date:         form.Date,
		Doctor:       form.Doctor,
		Complaint:    form.Complaint,
		User_id:      form.User_id,
	}
	fmt.Println("data date=", data.Date)
	reserve, err := models.AddReserve(data)
	if err != nil {
		fmt.Println(err)
	}
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Create Data",
		Results: reserve,
	})
}
