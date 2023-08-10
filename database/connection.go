package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Connect will create a connection to our database
func Connect() {
	//Load in the .env
	godotenv.Load(".env")

	//build connection string, note spaces are required
	db_connectionString := "host=" + os.Getenv("db_host") + " port=" + os.Getenv("db_port") + " user=" + os.Getenv("db_user") + " password=" + os.Getenv("db_pass") + " dbname=" + os.Getenv("db_name") + " sslmode=" + os.Getenv("db_sslmode")

	//Initiate DB pointer / possible error
	dbConnection, onConnectAttemptErr := sql.Open("postgres", db_connectionString)

	//Validate sql.Open operation
	if onConnectAttemptErr != nil {
		log.Fatalf("error encountered when attempting to open connection to database: %v", onConnectAttemptErr)
	}

	//Ping database to ensure connection is alive
	onPingErr := dbConnection.Ping()
	if onPingErr != nil {
		log.Fatalf("error encountered when attempting to ping database: %v", onPingErr)
	}

}
