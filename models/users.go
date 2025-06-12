package models

import (
	"apotekerBE/lib"
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id           int    `json:"id" form:"id"`
	Full_Name    string `json:"fullname" form:"fullname"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Role_Id      int    `json:"role_id" form:"role_id"`
	Image        string `json:"image" form:"-"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
}

type CreateProfile struct {
	Id           int    `json:"id" form:"id"`
	Full_Name    string `json:"fullname" form:"fullname"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Role_Id      int    `json:"role_id" form:"role_id"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
}

type PointProfile struct {
	Profile
	Point int
}

type RemoveUserData struct {
	Id           int    `json:"id" form:"id"`
	Full_Name    string `json:"fullname" form:"fullname"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Role_Id      int    `json:"role_id" form:"role_id"`
	Image        string `json:"image" form:"-"`
	Email        string `json:"email" form:"email"`
}

type ListProfile []Profile

func FindOneProfile(paramId int) Profile {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var profile Profile

	conn.QueryRow(context.Background(), `
	SELECT id,  fullname, phone_number, role_id, image, email, password 
	FROM users
	WHERE id = $1
	`, paramId).Scan(&profile.Id, &profile.Full_Name, &profile.Phone_number, &profile.Role_Id, &profile.Image, &profile.Email, &profile.Password)

	return profile
}

func FindAllUsers(page int, limit int, search string, sort string) ([]Profile, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	// var profiles []Profile
	offset := (page - 1) * limit
	search = fmt.Sprintf("%%%s%%", search)
	query := fmt.Sprintf(`SELECT id,  fullname, phone_number, role_id, COALESCE(image, '-') as image, email, password FROM users WHERE fullname ILIKE $1 ORDER BY fullname %s LIMIT $2 OFFSET $3`, sort)
	rows, err := conn.Query(context.Background(), query, search, limit, offset)
	if err != nil {
		fmt.Println("Error Find All Users", err)
		return nil, err
	}
	reserve, err := pgx.CollectRows(rows, pgx.RowToStructByName[Profile])
	if err != nil {
		fmt.Println("Error Collect Rows", err)
		return nil, err
	}
	// for rows.Next() {
	// var profile Profile
	// if err := rows.Scan(&profile.Id, &profile.Full_Name, &profile.Phone_number, &profile.Role_Id, &profile.Image, &profile.Email, &profile.Password); err != nil {
	// return nil, err
	// }
	// profiles = append(profiles, profile)
	// }

	return reserve, nil
}

func UpdateDataUser(user Profile, userId int) error {
	conn := lib.DB()
	defer conn.Close(context.Background())

	query := `UPDATE users SET `
	params := []interface{}{}
	paramIndex := 1

	if user.Full_Name != "" {
		query += fmt.Sprintf("fullname = $%d,", paramIndex)
		params = append(params, user.Full_Name)
		paramIndex++
	}

	if user.Phone_number != "" {
		query += fmt.Sprintf("phone_number = $%d,", paramIndex)
		params = append(params, user.Phone_number)
		paramIndex++
	}

	if user.Role_Id != 0 {
		// validasi role_id tetap perlu
		var roleExists bool
		err := conn.QueryRow(context.Background(), `SELECT EXISTS(SELECT 1 FROM role WHERE id = $1)`, user.Role_Id).Scan(&roleExists)
		if err != nil || !roleExists {
			return fmt.Errorf("invalid role_id")
		}
		query += fmt.Sprintf("role_id = $%d,", paramIndex)
		params = append(params, user.Role_Id)
		paramIndex++
	}

	if user.Image != "" {
		query += fmt.Sprintf("image = $%d,", paramIndex)
		params = append(params, user.Image)
		paramIndex++
	}

	if user.Email != "" {
		query += fmt.Sprintf("email = $%d,", paramIndex)
		params = append(params, user.Email)
		paramIndex++
	}

	if user.Password != "" {
		query += fmt.Sprintf("password = $%d,", paramIndex)
		params = append(params, user.Password)
		paramIndex++
	}

	// hapus koma terakhir
	query = strings.TrimSuffix(query, ",")

	query += fmt.Sprintf(" WHERE id = $%d", paramIndex)
	params = append(params, userId)

	_, err := conn.Exec(context.Background(), query, params...)
	return err
}

type Role struct {
	Id           int     `form:"id"` // ID tetap wajib, bukan pointer
	Full_Name    *string `form:"fullname"`
	Phone_number *string `form:"phone_number"`
	Role_Id      *int    `form:"role_id"`
	Image        *string `form:"image"`
	Email        *string `form:"email"`
	Password     *string `form:"password"`
}

func UpdateDataRole(user Role) error {
	conn := lib.DB()
	defer conn.Close(context.Background())

	query := `UPDATE users SET `
	params := []interface{}{}
	paramIndex := 1

	if user.Full_Name != nil {
		query += fmt.Sprintf("fullname = $%d,", paramIndex)
		params = append(params, *user.Full_Name)
		paramIndex++
	}

	if user.Phone_number != nil {
		query += fmt.Sprintf("phone_number = $%d,", paramIndex)
		params = append(params, *user.Phone_number)
		paramIndex++
	}

	if user.Role_Id != nil {
		query += fmt.Sprintf("role_id = $%d,", paramIndex)
		params = append(params, *user.Role_Id)
		paramIndex++
	}

	if user.Image != nil {
		query += fmt.Sprintf("image = $%d,", paramIndex)
		params = append(params, *user.Image)
		paramIndex++
	}

	if user.Email != nil {
		query += fmt.Sprintf("email = $%d,", paramIndex)
		params = append(params, *user.Email)
		paramIndex++
	}

	if user.Password != nil {
		query += fmt.Sprintf("password = $%d,", paramIndex)
		params = append(params, *user.Password)
		paramIndex++
	}

	// hapus koma terakhir
	query = strings.TrimSuffix(query, ",")

	query += fmt.Sprintf(" WHERE id = $%d", paramIndex)
	params = append(params, user.Id)

	_, err := conn.Exec(context.Background(), query, params...)
	return err
}

func CreateUser(user CreateProfile) (CreateProfile, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var userData CreateProfile
	err := conn.QueryRow(context.Background(), `
	INSERT INTO users (fullname, phone_number, role_id, email, password)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, fullname, phone_number, role_id, email, password
	`, user.Full_Name, user.Phone_number, user.Role_Id, user.Email, user.Password).Scan(&userData.Id, &userData.Full_Name, &userData.Phone_number, &userData.Role_Id, &userData.Email, &userData.Password)

	if err != nil {
		return userData, err
	}

	return userData, nil
}

func FindUserByEmail(email string) (CreateProfile, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var user CreateProfile
	err := conn.QueryRow(context.Background(), `
	SELECT id, fullname, phone_number, role_id, email, password 
	FROM users WHERE email = $1
	`, email).Scan(&user.Id, &user.Full_Name, &user.Phone_number, &user.Role_Id, &user.Email, &user.Password)

	if err != nil {
		if err == pgx.ErrNoRows {
			return user, nil // tidak ditemukan
		}
		return user, err // error lain
	}

	return user, nil
}

func RemoveUser(id int) RemoveUserData {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var user RemoveUserData
	conn.QueryRow(context.Background(), `DELETE FROM users WHERE id = $1 RETURNING id, fullname, phone_number, role_id, image, email`, id).Scan(&user.Id, &user.Full_Name, &user.Phone_number, &user.Role_Id, &user.Image, &user.Email)
	// _, err := conn.Exec(context.Background(), `DELETE FROM users WHERE id = $1 Re`, id)
	// if err != nil {
	// return err
	// }

	return user
}

func CountDataAllUser(search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var count int
	search = fmt.Sprintf("%%%s%%", search)

	conn.QueryRow(context.Background(), `
	SELECT COUNT(id)
	FROM users
	WHERE fullname ILIKE $1
	`, search).Scan((&count))
	return count
}
