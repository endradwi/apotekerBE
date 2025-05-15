package models

import (
	"apotekerBE/lib"
	"context"
	"fmt"
	"time"
)

type ReserveData struct {
	Id           int       `json:"id" form:"id"`
	Fullname     string    `json:"fullname" form:"fullname"`
	Phone_number string    `json:"phone_number" form:"phone_number"`
	Age          string    `json:"age" form:"age"`
	Date         time.Time `json:"date" form:"date"`
	Doctor       string    `json:"doctor" form:"doctor"`
	Complaint    string    `json:"complaint" form:"complaint"`
	User_id      int       `json:"user_id" form:"user_id"`
}

type StatusRegister struct {
	ReserveData
	Status string `json:"status" from:"status"`
}

func AddReserve(reserve StatusRegister) (StatusRegister, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())
	fmt.Println("data baru=", reserve)
	fmt.Println("date", reserve.Date)
	var reserveAdd StatusRegister
	// var tempDate time.Time
	err := conn.QueryRow(context.Background(), `
	INSERT INTO reserve (fullname, phone_number, age, date,doctor, complaint, user_id, status ) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, fullname, phone_number, age, date, doctor, complaint, user_id, status
	`, reserve.Fullname, reserve.Phone_number, reserve.Age, reserve.Date, reserve.Doctor, reserve.Complaint, reserve.User_id, reserve.Status).Scan(&reserveAdd.Id, &reserveAdd.Fullname, &reserveAdd.Phone_number, &reserveAdd.Age, &reserveAdd.Date, &reserveAdd.Doctor, &reserveAdd.Complaint, &reserveAdd.User_id, &reserveAdd.Status)

	// reserveAdd.Date = CustomDate(tempDate)
	fmt.Println("Reserve Add = ", reserveAdd)
	if err != nil {
		return StatusRegister{}, err
	}

	return reserveAdd, nil

}

func GetAllReserve() ([]StatusRegister, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var getAll []StatusRegister
	rows, err := conn.Query(context.Background(), `
	SELECT id,  fullname, phone_number, age, date, doctor, complaint, user_id, status
	FROM reserve
	`)
	if err != nil {
		fmt.Println("Error Find All Users", err)
		return nil, err
	}

	for rows.Next() {
		var data StatusRegister
		if err := rows.Scan(&data.Id, &data.Fullname, &data.Phone_number, &data.Age, &data.Date, &data.Doctor, &data.Complaint, &data.User_id, &data.Status); err != nil {
			return nil, err
		}
		getAll = append(getAll, data)
	}
	return getAll, err

}

func GetAllReserveByUser(userId int) ([]StatusRegister, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var getAll []StatusRegister
	rows, err := conn.Query(context.Background(), `
	SELECT id,  fullname, phone_number, age, date, doctor, complaint, user_id , status
	FROM reserve
	WHERE user_id = $1
	`, userId)
	if err != nil {
		fmt.Println("Error Find All Users", err)
		return nil, err
	}

	for rows.Next() {
		var data StatusRegister
		if err := rows.Scan(&data.Id, &data.Fullname, &data.Phone_number, &data.Age, &data.Date, &data.Doctor, &data.Complaint, &data.User_id, &data.Status); err != nil {
			return nil, err
		}
		getAll = append(getAll, data)

	}
	return getAll, err

}
