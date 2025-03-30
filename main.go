package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	// MySQL connection string: user:password@tcp(host:port)/dbname
	dsn := "testuser:testpass@tcp(localhost:3306)/testdb"

	// Open connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("✅ Connected to MySQL!")
	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}
	defer db.Close()
}
