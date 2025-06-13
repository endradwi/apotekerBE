package controllers

import (
	"apotekerBE/lib"
	"apotekerBE/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthRegister(ctx *gin.Context) {
	var form models.Users
	err := ctx.ShouldBind(&form)
	if err != nil {
		if strings.Contains(err.Error(), "Field validation for 'Email' failed on the 'email' tag") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Invalid Email Format",
			})
			return
		}
		if strings.Contains(err.Error(), "Field validation for 'Password' failed on the 'min' tag") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Minimum Password 6",
			})
			return
		}
		if strings.Contains(err.Error(), "Field validation for 'Password' failed on the 'containsany' tag") {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Password Must Include 1 Uppercase OR 1 Number",
			})
			return
		}

		log.Println(err)
	}
	findUser := models.FindOneUserByEmail(form.Email)

	if form.Email == findUser.Email {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Email Has Registered",
		})
		return
	}
	hash := lib.CreateHash(form.Password)
	form.Password = hash

	profile := models.RelationProfile{
		Email:        form.Email,
		Password:     form.Password,
		Full_Name:    "",
		Phone_Number: "",
		User_Id:      form.Id,
		Role_Id:      2,
		Image:        "",
	}

	models.AddUsers(profile)

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Register Success"})

}

func AuthLogin(ctx *gin.Context) {
	var form models.Users
	ctx.ShouldBind(&form)

	foundUser := models.FindOneUserByEmail(form.Email)
	if form.Email != foundUser.Email {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Email & Password",
		})
		return
	}

	match := lib.GenerateTokenArgon(form.Password, foundUser.Password)
	if !match {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "Invalid Email & Password",
		})
		return
	}

	token := lib.GeneretedToken(struct {
		UserId int `json:"userId"`
	}{
		UserId: foundUser.Id,
	})
	combine := gin.H{
		"token":   token,
		"role_id": foundUser.Role_Id,
	}

	ctx.SetCookie("token", token, 3600, "/", "", true, true)

	// })

	// fmt.Println("data = ", token)
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "Login Success",
		Results: combine,
	})

}
