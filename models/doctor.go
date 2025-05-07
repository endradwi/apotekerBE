package models

import (
	"apotekerBE/lib"
	"context"
)

type Doctor struct {
	Id         int    `json:"id" form:"id"`
	DoctorName string `json:"doctor_name" form:"doctor_name"`
	Speciality string `json:"speciality" form:"speciality"`
}

type ListDoctor []Doctor

func GetDoctor() ([]Doctor, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var d Doctor

	rows, err := conn.Query(context.Background(), `
		SELECT id, name, spesialis 
		FROM doctor
	`)
	if err != nil {
		return nil, err
	}

	var doctors []Doctor
	for rows.Next() {
		if err := rows.Scan(&d.Id, &d.DoctorName, &d.Speciality); err != nil {
			return nil, err
		}
		doctors = append(doctors, d)
	}

	return doctors, nil
}
