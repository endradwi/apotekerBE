package models

import (
	"apotekerBE/lib"
	"context"
	"fmt"
)

type ReserveData struct {
	Id           int    `json:"id" form:"id"`
	Fullname     string `json:"fullname" form:"fullname"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Age          string `json:"age" form:"age"`
	Complaint    string `json:"complaint" form:"complaint"`
	User_id      int    `json:"user_id" form:"user_id"`
	Doctor_id    int    `json:"doctor_id" form:"doctor_id"`
}

func AddReserve(reserve ReserveData) (ReserveData, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())
	fmt.Println("data baru=", reserve)

	// data := conn.QueryRow(context.Background(), `SELECT id FROM users WHERE id = $1`, reserve.User_id).Scan(&reserve.User_id)
	var reserveAdd ReserveData
	err := conn.QueryRow(context.Background(), `
		INSERT INTO reserve (fullname, phone_number, age, complaint, user_id, doctor_id) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, fullname, phone_number, age, complaint, user_id, doctor_id
	`, reserve.Fullname, reserve.Phone_number, reserve.Age, reserve.Complaint, reserve.User_id, reserve.Doctor_id).Scan(&reserveAdd.Id, &reserveAdd.Fullname, &reserveAdd.Phone_number, &reserveAdd.Age, &reserveAdd.Complaint, &reserveAdd.User_id, &reserveAdd.Doctor_id)

	fmt.Println("Reserve Add = ", reserveAdd)
	if err != nil {
		return ReserveData{}, err
	}

	return reserveAdd, nil

}
