package controllers

import (
	"fmt"
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
// @Param Profile formData models.Profile true "Add New Movie"
// @Success 200 {object} Response{results=models.Profile}
// @Router /profile [patch]
func EditProfile(ctx *gin.Context) {
	// paramId, _ := strconv.Atoi(ctx.Param("id"))
	// if paramId != profile.Id {
	// 	ctx.JSON(http.StatusNotFound, Response{
	// 		Success: false,
	// 		Message: "ID Not Found",
	// 	})
	// 	return
	// }
	val, _ := ctx.Get("userId")
	userId := int(val.(float64))

	profile := models.FindOneProfile(userId)
	// handling body form without file
	if err := ctx.ShouldBind(&profile); err != nil {
		ctx.Status(http.StatusInternalServerError)
	}
	f, _ := ctx.MultipartForm()
	file, _ := ctx.FormFile("image")

	profile.Email = f.Value["email"][0]

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

	updated := models.UpdatedProfile(profile)
	// profile.Password = f.Value["password"][0]

	hash := lib.CreateHash(profile.Password)
	if profile.Password != "" {
		profile.Password = hash
	}
	fmt.Println("data upload =", updated)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Update User Success",
		Results: profile,
	})
}

// func EditUser(ctx *gin.Context) {

// 	file, err := ctx.FormFile("image")
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, Response{
// 			Success: false,
// 			Message: "No file provided",
// 		})
// 		return
// 	}

// 	var fileName string
// 	if file.Filename != "" {

// 		splitFile := strings.Split(file.Filename, ".")
// 		if len(splitFile) < 2 {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Invalid file format",
// 			})
// 			return
// 		}

// 		fileExt := strings.ToLower(splitFile[len(splitFile)-1])

// 		allowedExtensions := map[string]bool{
// 			"jpg": true,
// 			"png": true,
// 		}

// 		if !allowedExtensions[fileExt] {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Only .jpg and .png files are allowed",
// 			})
// 			return
// 		}

// 		fileName = uuid.New().String()
// 		filePath := fmt.Sprintf("upload/users/%s.%s", fileName, fileExt)

// 		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
// 			ctx.JSON(http.StatusInternalServerError, Response{
// 				Success: false,
// 				Message: "Error saving the file",
// 			})
// 			return
// 		}
// 	}

// 	firstname := ctx.PostForm("firstname")
// 	lastname := ctx.PostForm("lastname")
// 	email := ctx.PostForm("email")
// 	password := ctx.PostForm("password")

// 	if email == "" {
// 		ctx.JSON(http.StatusBadRequest, Response{
// 			Success: false,
// 			Message: "Email is required",
// 		})
// 		return
// 	}

// 	user := models.FindOneUserByEmail(email)
// 	if user == (models.Users{}) {
// 		ctx.JSON(http.StatusNotFound, Response{
// 			Success: false,
// 			Message: "Email not found",
// 		})
// 		return
// 	}

// 	if password != "" && !strings.Contains(password, "$argon2i$v=19$m=65536,t=1,p=2$") {
// 		password = lib.CreateHash(password)
// 	}

// 	updatedUser := models.Profile{
// 		First_Name: firstname,
// 		Last_Name:  lastname,
// 		Image:      fileName + "." + strings.Split(file.Filename, ".")[1],
// 		Email:      email,
// 		Password:   password,
// 	}

// 	result := models.UpdateUser(updatedUser)

//		ctx.JSON(http.StatusOK, Response{
//			Success: true,
//			Message: "User updated successfully",
//			Results: result,
//		})
//	}

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
	userId := int(val.(float64))

	profile := models.FindOneProfile(userId)

	if isAvail {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "See All Profile",
			Results: profile,
		})
	}
}
