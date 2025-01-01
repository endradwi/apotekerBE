package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DB() *pgx.Conn {
	godotenv.Load()
	connstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	// conn, _ := pgx.Connect(context.Background(), connstring)
	conn, _ := pgx.Connect(context.Background(), connstring)

	return conn
}
