package models

import (
	"apotekerBE/lib"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
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
	RecMedic string `json:"rec_medic" form:"rec_medic"`
	Status   string `json:"status" form:"status"`
}

func AddReserve(reserve StatusRegister) (StatusRegister, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())
	fmt.Println("data baru=", reserve)
	fmt.Println("date", reserve.Date)
	var reserveAdd StatusRegister
	// var tempDate time.Time
	err := conn.QueryRow(context.Background(), `
	INSERT INTO reserve (fullname, phone_number, age, date,doctor, complaint, user_id, status, rec_medic ) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, fullname, phone_number, age, date, doctor, complaint, user_id, status, rec_medic
	`, reserve.Fullname, reserve.Phone_number, reserve.Age, reserve.Date, reserve.Doctor, reserve.Complaint, reserve.User_id, reserve.Status, reserve.RecMedic).Scan(&reserveAdd.Id, &reserveAdd.Fullname, &reserveAdd.Phone_number, &reserveAdd.Age, &reserveAdd.Date, &reserveAdd.Doctor, &reserveAdd.Complaint, &reserveAdd.User_id, &reserveAdd.Status, &reserveAdd.RecMedic)

	// reserveAdd.Date = CustomDate(tempDate)
	fmt.Println("Reserve Add = ", reserveAdd)
	if err != nil {
		return StatusRegister{}, err
	}

	return reserveAdd, nil

}

func GetAllReserve(page int, limit int, search string, sort string) ([]StatusRegister, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())
	// var getAll []StatusRegister
	offset := (page - 1) * limit
	search = fmt.Sprintf("%%%s%%", search)
	query := fmt.Sprintf(`SELECT id,  fullname, phone_number, age, date, doctor, complaint, user_id, status, rec_medic FROM reserve
	WHERE fullname ILIKE $1
	ORDER BY date %s
	LIMIT $2 OFFSET $3`, sort)
	rows, err := conn.Query(context.Background(), query, search, limit, offset)
	if err != nil {
		fmt.Println("Error Find All Users", err)
		return nil, err
	}
	reserve, _ := pgx.CollectRows(rows, pgx.RowToStructByName[StatusRegister])
	// for rows.Next() {
	// 	var data StatusRegister
	// 	if err := rows.Scan(&data.Id, &data.Fullname, &data.Phone_number, &data.Age, &data.Date, &data.Doctor, &data.Complaint, &data.User_id, &data.Status); err != nil {
	// 		return nil, err
	// 	}
	// 	getAll = append(getAll, data)
	// }
	return reserve, err

}

func GetAllReserveByUser(userId int, page int, limit int, search string, sort string) ([]StatusRegister, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())
	// var getAll []StatusRegister
	offset := (page - 1) * limit
	search = fmt.Sprintf("%%%s%%", search)
	query := fmt.Sprintf(`SELECT id,  fullname, phone_number, age, date, doctor, complaint, user_id, status, rec_medic FROM reserve
	WHERE user_id = $1 AND fullname ILIKE $2 ORDER BY date %s
	LIMIT $3 OFFSET $4 `, sort)
	rows, err := conn.Query(context.Background(), query, userId, search, limit, offset)
	if err != nil {
		fmt.Println("Error Find All Users", err)
		return nil, err
	}
	reserve, err := pgx.CollectRows(rows, pgx.RowToStructByName[StatusRegister])
	if err != nil {
		fmt.Println("Error Collect Rows", err)
		return nil, err
	}
	// for rows.Next() {
	// var data StatusRegister
	// if err := rows.Scan(&data.Id, &data.Fullname, &data.Phone_number, &data.Age, &data.Date, &data.Doctor, &data.Complaint, &data.User_id, &data.Status); err != nil {
	// return nil, err
	// }
	// getAll = append(getAll, data)
	//
	// }
	return reserve, err

}
func CountDataAll(search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var count int
	search = fmt.Sprintf("%%%s%%", search)

	conn.QueryRow(context.Background(), `
	SELECT COUNT(id)
	FROM reserve
	WHERE fullname ILIKE $1
	`, search).Scan((&count))
	return count
}

func CountDataAllPasien(userID int, search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var count int

	query := `
	SELECT COUNT(id)
	FROM reserve
	WHERE user_id = $1
	`

	// Jika kamu ingin pakai search juga
	if search != "" {
		query += ` AND (nama ILIKE '%' || $2 || '%')`
		conn.QueryRow(context.Background(), query, userID, search).Scan(&count)
	} else {
		conn.QueryRow(context.Background(), query, userID).Scan(&count)
	}

	return count
}

func UpdateStatus(status StatusRegister) ([]StatusRegister, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	query := `UPDATE reserve SET `
	param := []interface{}{}
	paramIndex := 1

	if status.Fullname != "" {
		query += fmt.Sprintf("fullname = $%d,", paramIndex)
		param = append(param, status.Fullname)
		paramIndex++
	}
	if status.Phone_number != "" {
		query += fmt.Sprintf("phone_number = $%d,", paramIndex)
		param = append(param, status.Phone_number)
		paramIndex++
	}
	if status.Age != "" {
		query += fmt.Sprintf("age = $%d,", paramIndex)
		param = append(param, status.Age)
		paramIndex++
	}
	if !status.Date.IsZero() {
		query += fmt.Sprintf("date = $%d,", paramIndex)
		param = append(param, status.Date)
		paramIndex++
	}
	if status.Doctor != "" {
		query += fmt.Sprintf("doctor = $%d,", paramIndex)
		param = append(param, status.Doctor)
		paramIndex++
	}
	if status.Complaint != "" {
		query += fmt.Sprintf("complaint = $%d,", paramIndex)
		param = append(param, status.Complaint)
		paramIndex++
	}
	if status.Status != "" {
		query += fmt.Sprintf("status = $%d,", paramIndex)
		param = append(param, status.Status)
		paramIndex++
	}
	if status.RecMedic != "" {
		status.RecMedic = lib.ToOrderList(status.RecMedic)
		query += fmt.Sprintf("rec_medic = $%d,", paramIndex)
		param = append(param, status.RecMedic)
		paramIndex++
	}

	if len(param) == 0 {
		return nil, fmt.Errorf("tidak ada data yang diubah")
	}

	// Finalisasi query
	query = strings.TrimSuffix(query, ",")
	query += fmt.Sprintf(" WHERE id = $%d", paramIndex)
	param = append(param, status.Id)

	// Eksekusi update
	_, err := conn.Exec(context.Background(), query, param...)
	if err != nil {
		return nil, fmt.Errorf("gagal update data: %v", err)
	}
	fmt.Println("query", query)

	// Ambil semua reservasi milik user tersebut
	row := conn.QueryRow(context.Background(), `
	SELECT id, fullname, phone_number, age, date, doctor, complaint, user_id, rec_medic, status
	FROM reserve
	WHERE id = $1
`, status.Id)

	var result StatusRegister
	err = row.Scan(&result.Id, &result.Fullname, &result.Phone_number, &result.Age, &result.Date, &result.Doctor, &result.Complaint, &result.User_id, &result.RecMedic, &result.Status)
	if err != nil {
		return nil, fmt.Errorf("gagal ambil data setelah update: %v", err)
	}

	return []StatusRegister{result}, nil

}
