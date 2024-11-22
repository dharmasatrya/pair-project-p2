package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// ConnectDB connects to the database
func ConnectDB() *sql.DB {
	// Connection details
	db_name := "postgres"
	db_user := "postgres.nqujnuwnntdamhodrbrj"
	db_pass := "akusukangoding"
	db_host := "aws-0-ap-southeast-1.pooler.supabase.com"

	// Create connection string
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require", db_name, db_user, db_pass, db_host)

	// Open the connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
		return nil
	}

	// Check if the connection is actually working
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
		return nil
	}

	// Connection was successful
	fmt.Println("Successfully connected to the database!")

	// Return the DB object for further use
	return db
}
