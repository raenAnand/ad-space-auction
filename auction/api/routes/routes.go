package routes

import (
	"net/http"

	"github.com/raenAnand/ad-space-auction/auction/api/handlers"
)

// SetupRoutes sets up the routes for the auction service
func SetupRoutes() {
	http.HandleFunc("/auctions/{id}", handlers.GetAuctionHandler)    // GET
	http.HandleFunc("/auctions", handlers.CreateAuctionHandler)     // POST
	http.HandleFunc("/auctions/{id}", handlers.UpdateAuctionHandler) // PUT
}
