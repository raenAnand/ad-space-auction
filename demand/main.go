package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raenAnand/ad-space-auction/demand/api/routes"
	"github.com/raenAnand/ad-space-auction/demand/database"
)

func main() {
	// Connect to the database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Set up the routes
	routes.SetupRoutes()

	// Start the HTTP server
	port := ":9000"
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
