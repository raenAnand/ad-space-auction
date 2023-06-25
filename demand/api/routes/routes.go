package routes

import (
	"net/http"

	"github.com/raenAnand/ad-space-auction/demand/api/handlers"
)

// SetupRoutes sets up the routes for the demand side service
func SetupRoutes() {
	http.HandleFunc("/bids", handlers.GetAllBidsHandler)      // GET
	http.HandleFunc("/bids/{id}", handlers.GetBidHandler)     // GET
	http.HandleFunc("/bids", handlers.CreateBidHandler)       // POST
	http.HandleFunc("/bids/{id}", handlers.UpdateBidHandler)  // PUT
}
