package controllers

import (
	"apotekerBE/lib"
	"apotekerBE/models"
	"fmt"
	"math"
	"net/http"
	"path/filepath"
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
	users, err := models.FindAllUsers(page, limit, search, sortUser)
	if err != nil {
		fmt.Println("Error Get All User", err)
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to get users"})
		return
	}
	// Ambil jumlah total data
	count := models.CountDataAllUser(search)

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
		Message: "Get All User",
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

func EditProfile(ctx *gin.Context) {
	val, isAvail := ctx.Get("userId")
	if !isAvail {
		ctx.JSON(http.StatusUnauthorized, Response{
			Success: false,
			Message: "Unauthorized ID",
		})
		return
	}

	userId := val.(int)

	// Gunakan multipart form karena ada file
	err := ctx.Request.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Failed to parse form",
		})
		return
	}

	var profile models.Profile
	profile.Full_Name = ctx.PostForm("fullname")
	profile.Phone_number = ctx.PostForm("phone_number")
	profile.Email = ctx.PostForm("email")
	profile.Password = ctx.PostForm("password")

	// Handle role_id if diberikan
	roleIdStr := ctx.PostForm("role_id")
	if roleIdStr != "" {
		roleId, convErr := strconv.Atoi(roleIdStr)
		if convErr == nil {
			profile.Role_Id = roleId
		}
	}

	// Optional image upload
	file, err := ctx.FormFile("image")
	if err == nil && file != nil && file.Filename != "" {
		filename := uuid.New().String()
		ext := strings.ToLower(filepath.Ext(file.Filename)) // includes dot

		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "File must be .jpg, .jpeg, or .png",
			})
			return
		}

		if file.Size > 1*1024*1024 {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "File too large (max 1MB)",
			})
			return
		}

		storedFile := fmt.Sprintf("%s%s", filename, ext)
		savePath := fmt.Sprintf("upload/profile/%s", storedFile)
		if saveErr := ctx.SaveUploadedFile(file, savePath); saveErr != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "Failed to save image",
			})
			return
		}

		profile.Image = storedFile
	}

	// Hash password jika ada
	if profile.Password != "" {
		profile.Password = lib.CreateHash(profile.Password)
	}

	// ✅ Panggil update dan dapatkan data hasil update
	profileData, err := models.UpdateDataUser(profile, userId)
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
		Results: profileData,
	})
}

func EditRoleUser(ctx *gin.Context) {
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
	var profile models.Role
	profile.Id = id
	err = ctx.ShouldBind(&profile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	data := models.UpdateDataRole(profile)
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
	var formData models.CreateProfile
	err := ctx.ShouldBind(&formData)
	findEmail, err := models.FindUserByEmail(formData.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "Failed to check email",
		})
		return
	}
	if formData.Email == findEmail.Email {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Email already exists",
		})
		return
	}
	fmt.Println("Find email =", findEmail)
	fmt.Println("Data Email =", formData.Email)
	fmt.Println("Form data 1=", formData.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}
	hash := lib.CreateHash(formData.Password)
	formData.Password = hash

	fmt.Println("Form Password =", formData.Password)

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
