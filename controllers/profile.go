package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"test/lib"
	"test/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Profile godoc
// @Schemes
// @Description Update Movies
// @Tags Profile
// @Accept mpfd
// @Produce json
// @Param First_Name formData string true "Update First Name"
// @Param Last_Name formData string true "Update Last Name"
// @Param Image formData file false "Update Image"
// @Param Phone_Number formData string true "Update Phone_Number"
// @Param Email formData string true "Update Email"
// @Param Password formData string true "Update Password"
// @Success 200 {object} Response{results=models.Profile}
// @Router /profile [patch]
func EditProfile(ctx *gin.Context) {
	val, isAvail := ctx.Get("userId")
	if !isAvail {
		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "Unauthorized id",
		})
		return
	}
	log.Println("data val =", val)
	// userId := int(val.(float64))
	// log.Println("data =", userId)
	var profile models.Profile
	// handling body form without file
	ctx.ShouldBind(&profile)
	f, _ := ctx.MultipartForm()
	file, err := ctx.FormFile("image")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Unauthorized file name",
		})
	}

	profile.Email = f.Value["email"][0]
	profile.Password = f.Value["password"][0]

	if file.Filename != "" {
		filename := uuid.New().String()

		// handling extentioin .jpg dll
		splitedFilename := strings.Split(file.Filename, ".")
		ext := splitedFilename[len(splitedFilename)-1]
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		if ext != "jpg" && ext != "png" && ext != "jpeg" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Must Fill .jpg, .jpeg, .png",
			})
			return
		}

		// handling name file
		ctx.SaveUploadedFile(file, fmt.Sprintf("updload/profile/%s", storedFile))
		profile.Image = storedFile
	}

	// Validation Size File
	maxfile := 1 * 1024 * 1024
	if file.Size > int64(maxfile) {
		ctx.JSON(400, Response{
			Success: false,
			Message: "File to Large",
		})
		return
	}
	if profile.Password != "" {
		hash := lib.CreateHash(profile.Password)
		profile.Password = hash
	}

	updated := models.UpdatedProfile(profile, val.(int))

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update User Success",
		Results: updated,
	})
}

// Profile godoc
// @Summary Profile
// @Description  Get Profile
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} Response{results=models.PointProfile}
// @Security ApiKeyAuth
// @Router /profile [get]
func GetProfile(ctx *gin.Context) {
	val, isAvail := ctx.Get("userId")
	// userId := int(val.(float64))

	profile := models.FindOneProfile(val.(int))

	if isAvail {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Detail Profile",
			Results: profile,
		})
	}
}
