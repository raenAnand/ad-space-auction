package routes

import (
	"net/http"

	"github.com/raenAnand/ad-space-auction/supply/api/handlers"
)


// SetupRoutes sets up the routes for the supply side service
func SetupRoutes() {
	http.HandleFunc("/ad-spaces", handlers.GetAllAdSpacesHandler)       // GET
	http.HandleFunc("/ad-spaces/{id}", handlers.GetAdSpaceHandler)     // GET
	http.HandleFunc("/ad-spaces", handlers.CreateAdSpaceHandler)       // POST
	http.HandleFunc("/ad-spaces/{id}", handlers.UpdateAdSpaceHandler)  // PUT
}