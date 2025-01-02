package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"test/lib"
	"test/models"

	"github.com/gin-gonic/gin"
)

// Auth godoc
// @Schemes
// @Description Registrasi Account
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Input Email"
// @Param password formData string true "Input Password"
// @Success 200 {object} Response
// @Router /auth/register [post]
func AuthRegister(ctx *gin.Context) {
	var form models.Users
	ctx.ShouldBind(&form)

	findUser := models.FindOneUserByEmail(form.Email)
	// fmt.Println("error = ", findUser)
	if form.Email == findUser.Email {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Email Has Registered",
		})
		return
	}
	// if !strings.Contains(form.Email, "@") {
	// 	ctx.JSON(http.StatusBadRequest, Response{
	// 		Success: false,
	// 		Message: "Email Must Include @",
	// 	})
	// 	return
	// }
	if len(form.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Password Must be 6 Character",
		})
		return
	}
	if !strings.ContainsAny(form.Password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Password Must Include Uppercase Character",
		})
		return
	}
	if !strings.ContainsAny(form.Password, "0123456789") {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Password Must Include One Number",
		})
		return
	}
	if !strings.ContainsAny(form.Password, "!@#$%^&*()-_=+[]{}|;:,.<>?") {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Password Must Include Unique Character",
		})
		return
	}
	hash := lib.CreateHash(form.Password)
	form.Password = hash
	newUser := models.InsertUser(form)

	if newUser.Id > 0 {
		profile := models.RelationProfile{
			First_Name: "",
			Last_Name:  "",
			Image:      "",
			User_Id:    newUser.Id,
		}

		models.AddProfile(profile)

		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "Register Success"})
	} else {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "User Registration Failed",
		})
	}
}

// Auth godoc
// @Schemes
// @Description Login Account
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param email formData string true "Input Email"
// @Param password formData string true "Input Password"
// @Success 200 {object} Response
// @Router /auth/login [post]
func AuthLogin(ctx *gin.Context) {
	var form models.Users
	ctx.ShouldBind(&form)

	foundUser := models.FindOneUserByEmail(form.Email)
	fmt.Println("data = ", foundUser)
	if form.Email != foundUser.Email {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Email Not Found",
		})
		return
	}

	match := lib.GenerateTokenArgon(form.Password, foundUser.Password)
	if !match {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Password",
		})
		return
	}

	token := lib.GeneretedToken(struct {
		UserId int `json:"userId"`
	}{
		UserId: foundUser.Id,
	})
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Login Success",
		Results: token,
	})

}
