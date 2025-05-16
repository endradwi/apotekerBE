package controllers

import (
	"apotekerBE/lib"
	"apotekerBE/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUser(ctx *gin.Context) {
	val, isVail := ctx.Get("userId")
	fmt.Println("Get User =", val)
	if !isVail {
		ctx.JSON(http.StatusUnauthorized, Response{
			Success: false,
			Message: "Unauthorized",
		})
	}
	profile := models.FindOneProfile(val.(int))
	fmt.Println("Get profile =", profile)
	if isVail {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Get User",
			Results: profile,
		})
	}
}

func GetAllUser(ctx *gin.Context) {
	users, err := models.FindAllUsers()
	if err != nil {
		fmt.Println("Error Get All User", err)
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to get users"})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Get All User",
		Results: users,
	})
}

func EditProfile(ctx *gin.Context) {
	val, isAvail := ctx.Get("userId")
	if !isAvail {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "Unauthorized id",
		})
		return
	}

	var profile models.Profile
	err := ctx.ShouldBind(&profile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	fmt.Println("Data profile =", profile)

	// Default kosong, hanya akan diisi jika ada file dikirim
	var storedFile string

	// Coba ambil file
	file, err := ctx.FormFile("image")
	if err == nil && file != nil && file.Filename != "" {
		filename := uuid.New().String()
		splittedFilename := strings.Split(file.Filename, ".")
		ext := strings.ToLower(splittedFilename[len(splittedFilename)-1])

		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "File must be .jpg, .jpeg, or .png",
			})
			return
		}

		// Validasi size
		maxfile := 1 * 1024 * 1024
		if file.Size > int64(maxfile) {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "File is too large (max 1MB)",
			})
			return
		}

		// Simpan file
		storedFile = fmt.Sprintf("%s.%s", filename, ext)
		savePath := fmt.Sprintf("upload/profile/%s", storedFile)
		err = ctx.SaveUploadedFile(file, savePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "Failed to save file",
			})
			return
		}

		profile.Image = storedFile
	}

	// Hash password jika ada
	if profile.Password != "" {
		hash := lib.CreateHash(profile.Password)
		profile.Password = hash
	}

	// Panggil update function (harus mendukung partial update)
	err = models.UpdateDataUser(profile, val.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update User Success",
	})
}
func EditStatusUser(ctx *gin.Context) {
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
	var profile models.Status
	profile.Id = id
	err = ctx.ShouldBind(&profile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	fmt.Println("Data profile =", &profile)

	// Panggil update function (harus mendukung partial update)
	data := models.UpdateDataStatus(profile)
	fmt.Println("error=", err)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update User Success",
		Results: data,
	})
}

func AddAdmin(ctx *gin.Context) {
	var formData models.Profile
	err := ctx.ShouldBind(&formData)
	fmt.Println("Form data 1=", formData.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	var storedFile string
	file, err := ctx.FormFile("image")
	if err == nil && file != nil && file.Filename != "" {
		filename := uuid.New().String()
		splittedFilename := strings.Split(file.Filename, ".")
		ext := strings.ToLower(splittedFilename[len(splittedFilename)-1])

		if ext != "jpg" && ext != "jpeg" && ext != "png" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "File must be .jpg, .jpeg, or .png",
			})
			return
		}

		// Validasi size
		maxfile := 1 * 1024 * 1024
		if file.Size > int64(maxfile) {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "File is too large (max 1MB)",
			})
			return
		}

		// Simpan file
		storedFile = fmt.Sprintf("%s.%s", filename, ext)
		savePath := fmt.Sprintf("upload/admin/%s", storedFile)
		err = ctx.SaveUploadedFile(file, savePath)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "Failed to save file",
			})
			return
		}
		formData.Image = storedFile
	}

	// Hash password jika ada
	// if formData.Password != "" {
	// 	hash := lib.CreateHash(formData.Password)
	// 	formData.Password = hash
	// }

	data, err := models.CreateUser(formData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update User Success",
		Results: data,
	})
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	deleted := models.RemoveUser(id)
	fmt.Println("Deleted user =", deleted)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Deleted Success",
		Results: deleted,
	})

}
