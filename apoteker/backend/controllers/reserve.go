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
		Complaint:    form.Complaint,
		User_id:      form.User_id,
		Doctor_id:    form.Doctor_id,
	}

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
