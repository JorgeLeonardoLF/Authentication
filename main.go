package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JorgeLeonardoLF/Authentication/database"
	"github.com/JorgeLeonardoLF/Authentication/routes"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	//Load .env file
	godotenv.Load(".env")

	//Initialize connection to database
	database.Connect()

	/*initialize a router using the chi library
	Chi returns to us a Mux that implements the router interface*/
	router := chi.NewRouter()

	/*Call routes.Setup routes
	- To run configurations
	- Mount the vXRouter on to it
		- vXRouter will have the paths setup on to it
	- Pass in our parent database struct so that the routes have access to the database
	*/
	routes.Setup(router)

	//Setup Server
	portString := os.Getenv("go_server_port")
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
