package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DB() *pgx.Conn {
	godotenv.Load()
	config, err := pgx.ParseConfig("")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := pgx.Connect(context.Background(), config.ConnString())
	if err != nil {
		fmt.Println()
	}
	return conn
}
