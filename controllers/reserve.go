package controllers

import (
	"apotekerBE/models"
	"fmt"
	"math"
	"net/http"
	"strconv"

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
	var form models.StatusRegister
	err := ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Data",
		})
		return
	}
	form.User_id = val.(int)
	data := models.StatusRegister{
		ReserveData: models.ReserveData{
			Fullname:     form.Fullname,
			Phone_number: form.Phone_number,
			Age:          form.Age,
			Date:         form.Date,
			Doctor:       form.Doctor,
			Complaint:    form.Complaint,
			User_id:      form.User_id,
		},
		Status:   "pending",
		RecMedic: "-",
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

func GetAllReserveAdmin(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		fmt.Println("Invalid page number:", err)
	}
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	if err != nil || limit < 1 {
		fmt.Println("Invalid limit number:", err)
	}
	sortUser := ctx.DefaultQuery("sort", "ASC")
	if sortUser != "ASC" {
		sortUser = "DESC"
	}

	// Ambil data reservasi langsung dari database
	users, err := models.GetAllReserve(page, limit, search, sortUser)
	if err != nil {
		fmt.Println("Error Get All User", err)
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to get users",
		})
		return
	}

	// Ambil jumlah total data
	count := models.CountDataAll(search)

	// Hitung total halaman
	totalPage := int(math.Ceil(float64(count) / float64(limit)))

	nextPage := totalPage - page
	if nextPage < 0 {
		nextPage = 0
	}

	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	// Return response ke client
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Get All Reserve User",
		PageInfo: PageInfo{
			CurentPage: page,
			NextPage:   nextPage,
			PrevPage:   prevPage,
			TotalPage:  totalPage,
			TotalData:  count,
		},
		Results: users,
	})
}

func GetAllReserve(ctx *gin.Context) {
	val, isAvail := ctx.Get("userId")
	if !isAvail {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "Unauthorized id",
		})
		return
	}
	search := ctx.DefaultQuery("search", "")
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		fmt.Println("Invalid page number:", err)
	}
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	if err != nil || limit < 1 {
		fmt.Println("Invalid limit number:", err)
	}
	sortUser := ctx.DefaultQuery("sort", "ASC")
	if sortUser != "ASC" {
		sortUser = "DESC"
	}

	users, err := models.GetAllReserveByUser(val.(int), page, limit, search, sortUser)
	if err != nil {
		fmt.Println("Error Get All User", err)
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to get users"})
		return
	}
	// Ambil jumlah total data
	count := models.CountDataAll(search)

	// Hitung total halaman
	totalPage := int(math.Ceil(float64(count) / float64(limit)))

	nextPage := totalPage - page
	if nextPage < 0 {
		nextPage = 0
	}

	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Get All Reserve User By ID",
		PageInfo: PageInfo{
			CurentPage: page,
			NextPage:   nextPage,
			PrevPage:   prevPage,
			TotalPage:  totalPage,
			TotalData:  count,
		},
		Results: users,
	})
}

func UpdateStatus(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}
	fmt.Println("ID param =", id)
	var form models.StatusRegister
	err = ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Data",
		})
		return
	}

	form.Id = id
	data, err := models.UpdateStatus(form)
	fmt.Println("Data =", data)
	fmt.Println("Form Data =", form)
	fmt.Println("Form User ID =", form.User_id)
	fmt.Println("Form Status =", form.Status)
	fmt.Println("Form RecMedic =", form.RecMedic)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to update status",
		})
		return
	}
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update Status Success",
		Results: data,
	})

}
