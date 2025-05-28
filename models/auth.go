package models

import (
	"apotekerBE/lib"
	"context"
	"fmt"
)

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"email"`
	Password string `json:"password" form:"password" binding:"min=6,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"`
}
type UserAdmin struct {
	Id           int    `json:"id" form:"id"`
	Full_Name    string `json:"fullname" form:"fullname"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
}

type RelationProfile struct {
	Id           int    `json:"id"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"` // password
	Full_Name    string `json:"fullname" form:"fullname"`
	Last_Name    string `json:"last_name" form:"last_name"`
	Phone_Number string `json:"phone_number" form:"phone_number"`
	User_Id      int    `json:"user_id" form:"user_id"`
	Role_Id      int    `json:"role_id" form:"role_id"`
	Image        string `json:"image" form:"image"`
}

type ListUser []Users

func FindOneUserByEmail(email string) RelationProfile {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var user RelationProfile
	conn.QueryRow(context.Background(), `
	SELECT id, email, password, role_id 
	FROM users WHERE email = $1
	`, email).Scan(&user.Id, &user.Email, &user.Password, &user.Role_Id)
	return user
}

func AddUsers(profile RelationProfile) RelationProfile {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var profileAdd RelationProfile
	var user_id int

	// Query pertama: Insert data ke tabel users dan dapatkan user_id
	err := conn.QueryRow(context.Background(), `
		INSERT INTO users (email, password, fullname, phone_number, role_id, image) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`, profile.Email, profile.Password, profile.Full_Name, profile.Phone_Number, profile.Role_Id, profile.Image).Scan(&user_id)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return profileAdd
	}
	fmt.Println("err =", err)

	fmt.Println("user_id =", user_id)

	return profileAdd
}

func AddProfile(profile RelationProfile) RelationProfile {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var new RelationProfile
	conn.QueryRow(context.Background(),
		`INSERT INTO profile (fullname, phone_number, user_id, role_id)
		VALUES ($1, $2, $3, 2)
		RETURNING id, fullname, phone_number, user_id, role_id`,
		profile.Full_Name, profile.Phone_Number, profile.User_Id, profile.Role_Id).Scan(
		&new.Id, &new.Full_Name, &new.Phone_Number, &new.User_Id, &new.Role_Id,
	)

	return new
}
