package models

import (
	"context"
	"fmt"
	"log"
	"test/lib"
	"time"

	"github.com/jackc/pgx/v5"
)

type order struct {
	Id          int
	Profile_id  int    `json:"profile_id" form:"profile_id" db:"profile_id"`
	Movie_id    int    `json:"movie_id" form:"movie_id" db:"movie_id"`
	Cinema_id   int    `json:"cinema_id" form:"cinema_id" db:"cinema_id"`
	Seat_id     int    `json:"seat_id" form:"seat_id"`
	Payment_id  int    `json:"payment_id" form:"payment_id"`
	Tittle      string `json:"title" form:"title" db:"title"`
	Image       string `json:"image" form:"image" db:"image"`
	Genre       string `json:"genre" form:"genre" db:"genre"`
	Quantity    int    `json:"quantity" form:"quantity" db:"quantity"`
	TotalPrice  int    `json:"total_price" form:"total_price" db:"total_price"`
	Cinema_name string `json:"cinema_name" form:"cinema_name" db:"cinema_name"`
	Location    string `json:"location" form:"location" db:"location"`
}

type OrderBody struct {
	order
	Date string `json:"date" form:"date" `
	Time string `json:"time" form:"time" `
}

type OrderData struct {
	order
	Date time.Time `db:"date"`
	Time time.Time `db:"time"`
}

type Orders struct {
	Id         int    `json:"id"`
	Profile_Id int    `json:"profile_id" form:"profile_id"`
	Movie_Id   int    `json:"movie_id" form:"movie_id"`
	Tittle     string `json:"tittle" form:"tittle"`
	Genre      string `json:"genre" form:"genre"`
	Images     string `json:"image" form:"image"`
	Qty        int    `json:"qty" form:"qty"`
	Seat       string `json:"seat" form:"seat"`
	Cinema     string `json:"cinema" form:"cinema"`
}

type Payment struct {
	No_Rekening   int       `json:"no_rekening" form:"no_rekening"`
	Total_Payment int       `json:"total_payment" form:"total_payment"`
	Limit_Payment time.Time `json:"limit_payment" form:"limit_paymnet"`
}

type MoviesCinema struct {
	Id              int    `json:"id"`
	Tittle          string `json:"tittle" form:"tittle"`
	Genre           string `json:"genre" form:"genre"`
	Images          string `json:"image" form:"image"`
	Cinema          string `json:"cinema" form:"cinema"`
	Cinema_date     string `json:"cinema_date" form:"cinema_date"`
	Cinema_time     string `json:"cinema_time" form:"cinema_time"`
	Cinema_location string `json:"cinema_location" form:"cinema_location"`
}

type seat string

type SeatCinema struct {
	Id              int    `json:"id"`
	Tittle          string `json:"tittle" form:"tittle"`
	Genre           string `json:"genre" form:"genre"`
	Images          string `json:"image" form:"image"`
	Cinema          string `json:"cinema" form:"cinema"`
	Cinema_date     string `json:"cinema_date" form:"cinema_date"`
	Cinema_time     string `json:"cinema_time" form:"cinema_time"`
	Cinema_location string `json:"cinema_location" form:"cinema_location"`
	Price           int    `json:"price" from:"price"`
	Seats           []seat `json:"seat" from:"seat[]"`
}

type TotalSeatCinema struct {
	SeatCinema
	Total_Seat int
}

type ListCinema []MoviesCinema

type ListOrders []Orders

func OrderTicket(data OrderBody) OrderData {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var order OrderData
	cinema := struct {
		price int `db:"price"`
	}{}

	conn.QueryRow(context.Background(),
		`SELECT price FROM cinema WHERE id = $1`, data.Cinema_id).Scan(&cinema.price)

	totalPrice := cinema.price * data.Quantity

	log.Println("data = total price", totalPrice)
	conn.QueryRow(context.Background(), `
		INSERT INTO orders (profile_id, movie_id, cinema_id, seat_id, date, qty, total_price, payment_id ) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
		RETURNING id, profile_id, movie_id, cinema_id, seat_id, date, qty, total_price, payment_id`,
		data.Profile_id, data.Movie_id, data.Cinema_id, data.Seat_id, data.Date, data.Quantity, totalPrice, data.Payment_id).Scan(
		&order.Id, &order.Profile_id, &order.Movie_id, &order.Cinema_id, &order.Seat_id, &order.Date, &order.Quantity,
		&order.TotalPrice, &order.Payment_id,
	)

	return order
}

func BookingCinema(paramId int, searchName string, searchTime string, searchDate string, searchLocation string) ListCinema {
	conn := lib.DB()
	defer conn.Close(context.Background())
	// var movie ListCinema

	// searchingName := fmt.Sprintf("%%%s%%", searchName)
	searchingTime := fmt.Sprintf("%%%s%%", searchTime)
	// searchingDate := fmt.Sprintf("%%%s%%", searchDate)
	// searchingLocation := fmt.Sprintf("%%%s%%", searchLocation)

	rows, _ := conn.Query(context.Background(), `
	SELECT movies.id, movies.tittle, movies.genre,
	movies.images, cinema.name, cinema.cinema_date,
	cinema.cinema_time, cinema.cinema_location
	FROM cinema 
    JOIN movies ON cinema.movies_id = movies.id
    WHERE movies.id = $1 AND cinema.cinema_time LIKE $2
	`, paramId, searchingTime)

	cinema, _ := pgx.CollectRows(rows, pgx.RowToStructByName[MoviesCinema])
	log.Println("data = ", rows)
	return cinema
}

func SeatOrder() {
	conn := lib.DB()
	defer conn.Close(context.Background())

	conn.QueryRow(context.Background(), `
	`)
}
