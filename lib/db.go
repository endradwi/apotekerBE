package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DB() *pgx.Conn {
	godotenv.Load()
	// connstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
	// 	os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	// conn, _ := pgx.Connect(context.Background(), connstring)
	// conn, _ := pgx.Connect(context.Background(), connstring)
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

// package lib

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/joho/godotenv"
// )

// func DB() *pgx.Conn {
// 	// Load environment variables from .env file
// 	if err := godotenv.Load(); err != nil {
// 		fmt.Println("❌ Failed to load .env file:", err)
// 		return nil
// 	}

// 	// Get DB URL from env
// 	connString := os.Getenv("SUPABASE_DB_URL")
// 	if connString == "" {
// 		fmt.Println("❌ SUPABASE_DB_URL is empty")
// 		return nil
// 	}

// 	// Connect directly using the connection string
// 	conn, err := pgx.Connect(context.Background(), connString)
// 	if err != nil {
// 		fmt.Println("❌ Failed to connect to Supabase DB:", err)
// 		return nil
// 	}

// 	fmt.Println("✅ Connected to Supabase DB")
// 	return conn
// }
