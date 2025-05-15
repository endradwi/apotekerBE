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

// type CustomDate time.Time

// func (cd *CustomDate) UnmarshalJSON(b []byte) error {
// 	s := strings.Trim(string(b), `"`)
// 	t, err := time.Parse("02/01/2006", s)
// 	if err != nil {
// 		fmt.Println("error parse date", err)
// 		return err
// 	}
// 	*cd = CustomDate(t)
// 	return nil
// }

// func (cd *CustomDate) UnmarshalText(text []byte) error {
// 	s := string(text)
// 	t, err := time.Parse("02/01/2006", s)
// 	if err != nil {
// 		fmt.Println("error parse date from form-data", err)
// 		return err
// 	}
// 	*cd = CustomDate(t)
// 	return nil
// }

// func (cd CustomDate) ToTime() time.Time {
// 	return time.Time(cd)
// }

func AddReserve(reserve ReserveData) (ReserveData, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())
	fmt.Println("data baru=", reserve)
	fmt.Println("date", reserve.Date)
	var reserveAdd ReserveData
	// var tempDate time.Time
	err := conn.QueryRow(context.Background(), `
	INSERT INTO reserve (fullname, phone_number, age, date,doctor, complaint, user_id ) 
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, fullname, phone_number, age, date, doctor, complaint, user_id
	`, reserve.Fullname, reserve.Phone_number, reserve.Age, reserve.Date, reserve.Doctor, reserve.Complaint, reserve.User_id).Scan(&reserveAdd.Id, &reserveAdd.Fullname, &reserveAdd.Phone_number, &reserveAdd.Age, &reserveAdd.Date, &reserveAdd.Doctor, &reserveAdd.Complaint, &reserveAdd.User_id)

	// reserveAdd.Date = CustomDate(tempDate)
	fmt.Println("Reserve Add = ", reserveAdd)
	if err != nil {
		return ReserveData{}, err
	}

	return reserveAdd, nil

}
