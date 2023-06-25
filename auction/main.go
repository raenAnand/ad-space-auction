// supply-side-service/main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raenAnand/ad-space-auction/auction/api/routes"
	"github.com/raenAnand/ad-space-auction/auction/database"
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
	port := ":7000"
	fmt.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
