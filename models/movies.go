package models

import (
	"context"
	"fmt"
	"log"
	"test/lib"
	"time"

	"github.com/jackc/pgx/v5"
)

type Movies struct {
	Id       int    `json:"id"`
	Tittle   string `json:"tittle" form:"tittle" example:"Spiderman"`
	Genre    string `json:"genre" form:"genre" example:"Action"`
	Image    string `json:"image" form:"image" example:"Spiderman.jpg"`
	Synopsis string `json:"synopsis" form:"synopsis" example:"film action universal"`
	Author   string `json:"author" form:"author"`
	Actors   string `json:"actors" form:"actors"`
	// Release_date time.Time `json:"release_date" form:"release_date"`
	// Duration     time.Time `json:"duration" form:"duration"`
	Tag string `json:"tag" form:"tag"`
}

type Movie_body struct {
	Movies
	Release_date string `json:"release_date" form:"release_date"`
	Duration     string `json:"duration" form:"duration" `
}

type Movie_Data struct {
	Movies
	Release_date time.Time `db:"release_date"`
	Duration     time.Time `db:"duration"`
}
type MoviesAll struct {
	Id              int       `json:"id"`
	Tittle          string    `json:"tittle" form:"tittle"`
	Genre           string    `json:"genre" form:"genre"`
	Images          string    `json:"image" form:"image"`
	Synopsis        string    `json:"synopsis" form:"synopsis"`
	Author          string    `json:"author" form:"author"`
	Actors          string    `json:"actors" form:"actors"`
	Release_date    time.Time `json:"release_date" form:"release_date"`
	Duration        time.Time `json:"duration" form:"duration"`
	Price           int       `json:"price" form:"price"`
	Cinema          string    `json:"cinema" form:"cinema"`
	Cinema_date     time.Time `json:"cinema_date" form:"cinema_date"`
	Cinema_time     time.Time `json:"cinema_time" form:"cinema_time"`
	Cinema_location string    `json:"cinema_location" form:"cinema_location"`
}

type ListMovie []Movies
type ListMovieAll []MoviesAll

func FindAllMovie(page int, limit int, search string, sort string) ListMovie {
	conn := lib.DB()
	defer conn.Close(context.Background())
	offset := (page - 1) * limit

	searching := fmt.Sprintf("%%%s%%", search)

	query := fmt.Sprintf(`
	SELECT id, tittle, genre,images, synopsis,
	author, actors, release_date,
	duration, tag
	FROM movies
	WHERE tittle ILIKE $1
 	ORDER BY id %s
 	LIMIT $2 OFFSET $3
 `, sort)
	rows, _ := conn.Query(context.Background(), query, searching, limit, offset)
	log.Println("data = ", rows)
	movie, _ := pgx.CollectRows(rows, pgx.RowToStructByName[Movies])
	log.Println("data = ", movie)
	return movie
}
func CountData(search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var count int
	search = fmt.Sprintf("%%%s%%", search)

	conn.QueryRow(context.Background(), `
	SELECT COUNT(id)
	FROM movies
	WHERE tittle ILIKE $1
	`, search).Scan((&count))
	return count
}

func FindOneMovie(paramId int) Movie_Data {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var movie Movie_Data

	conn.QueryRow(context.Background(), `
	SELECT id, tittle, genre, synopsis, 
	author, actors, release_date, duration, tag
	FROM movies 
    WHERE id = $1
	`, paramId).Scan(&movie.Id, &movie.Tittle, &movie.Genre,
		&movie.Synopsis, &movie.Author, &movie.Actors,
		&movie.Release_date, &movie.Duration, &movie.Tag)
	return movie
}

func FindDetailMovie(paramId int) MoviesAll {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var movie MoviesAll

	conn.QueryRow(context.Background(), `
	SELECT movies.id, movies.tittle, movies.genre,
	movies.synopsis, movies.author, movies.actors,
	movies.release_date, movies.duration, cinema.name, 
	cinema_date.name_date, cinema_time.name_time,
	cinema_location.name_location
	FROM movies 
    JOIN cinema ON cinema.movies_id = movies.id
    JOIN cinema_date ON cinema_date.cinema_id = cinema.id
    JOIN cinema_time ON cinema_time.cinema_id = cinema.id
    JOIN cinema_location ON cinema_location.cinema_id = cinema.id
    WHERE movies.id = $1
	`, paramId).Scan(&movie.Id, &movie.Tittle, &movie.Genre,
		&movie.Synopsis, &movie.Author, &movie.Actors,
		&movie.Release_date, &movie.Duration, &movie.Cinema,
		&movie.Cinema_date, &movie.Cinema_time, &movie.Cinema_location)
	return movie
}

func CountMovie(search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var total int

	search = fmt.Sprintf("%%%s%%", search)

	conn.QueryRow(context.Background(), `
	SELECT COUNT(id) FROM movie WHERE title ILIKE $1
	`, search).Scan(&total)

	return total
}

func InsertMovie(data Movie_body) (Movie_Data, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieInsert Movie_Data

	if data.Release_date == "" {
		return Movie_Data{}, fmt.Errorf("release date is required")
	}

	movieDate, _ := time.Parse(time.DateOnly, data.Release_date)

	log.Println("movieDate =", movieDate)
	movieDuration, _ := time.Parse(time.TimeOnly, data.Duration)

	log.Println("data duration =", movieDuration)

	text := conn.QueryRow(context.Background(), `
	INSERT INTO movies (tittle, genre, images, synopsis,
	author, actors, release_date, duration, tag) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id, tittle, genre, images, synopsis,
	author, actors, release_date, duration, tag
	`, data.Tittle, data.Genre, data.Image, data.Synopsis,
		data.Author, data.Actors, movieDate, movieDuration, data.Tag).
		Scan(
			&movieInsert.Id,
			&movieInsert.Tittle,
			&movieInsert.Genre,
			&movieInsert.Image,
			&movieInsert.Synopsis,
			&movieInsert.Author,
			&movieInsert.Actors,
			&movieInsert.Release_date,
			&movieInsert.Duration,
			&movieInsert.Tag,
		)
	log.Println("data text =", text)
	return movieInsert, nil
}

func UpdateMovie(movie Movie_body) Movie_Data {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieUpdate Movie_Data

	// Check if Release_date is empty
	if movie.Release_date == "" {
		return Movie_Data{}
	}

	movieDate, err := time.Parse(time.DateOnly, movie.Release_date)
	if err != nil {
		return Movie_Data{}
	}

	log.Println(movie)

	movieDuration, err := time.Parse(time.TimeOnly, movie.Duration)
	if err != nil {
		return Movie_Data{}
	}

	conn.QueryRow(context.Background(), `
		UPDATE movie
		SET title = $1, genre = $2, images = $3, synopsis = $5,
		author = $6, actors = $7, release_date = $8, duration = $9,
		tag = $10  
		WHERE id = $11
		RETURNING id, title, genre, images, synopsis, author, actors, release_date, duration, tag  
	`, movie.Tittle, movie.Genre, movie.Image, movie.Synopsis,
		movie.Author, movie.Actors, movieDate, movieDuration, movie.Tag, movie.Id).Scan(
		&movieUpdate.Id,
		&movieUpdate.Tittle,
		&movieUpdate.Genre,
		&movieUpdate.Image,
		&movieUpdate.Synopsis,
		&movieUpdate.Author,
		&movieUpdate.Actors,
		&movieUpdate.Release_date,
		&movieUpdate.Duration,
		&movieUpdate.Tag,
	)

	return movieUpdate
}

func DeleteMovie(iddb int) Movies {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieDelete Movies

	conn.QueryRow(context.Background(), `
	DELETE FROM movie WHERE id = $1
	RETURNING  id, title, genre, images, synopsis,
	author, actors, release_date, duration, tag
	`, iddb).Scan(
		&movieDelete.Id,
		&movieDelete.Tittle,
		&movieDelete.Genre,
		&movieDelete.Image,
		&movieDelete.Synopsis,
		&movieDelete.Author,
		&movieDelete.Actors,
		// &movieDelete.Release_date,
		// &movieDelete.Duration,
		&movieDelete.Tag,
	)

	return movieDelete
}
