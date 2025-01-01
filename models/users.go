package models

import (
	"context"
	"test/lib"
)

type Users struct {
	Id       int    `json:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ListUser []Users

func FindOneUserByEmail(email string) Users {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var user Users
	conn.QueryRow(context.Background(), `
	SELECT id, email, password 
	FROM users WHERE email = $1
	`, email).Scan(&user.Id, &user.Email, &user.Password)
	return user
}

func InsertUser(user Users) Users {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var new Users

	conn.QueryRow(context.Background(), `
	INSERT INTO users(email, password) VALUES
	($1, $2)
	RETURNING id, email, password
	`, user.Email, user.Password).Scan(&new.Id,
		&new.Email, &new.Password)
	return new
}
