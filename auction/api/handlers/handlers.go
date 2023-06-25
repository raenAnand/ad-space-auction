package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/raenAnand/ad-space-auction/auction/api/models"
	"github.com/raenAnand/ad-space-auction/auction/database"
)

// GetAuctionHandler handles the GET /auctions/{id} endpoint to retrieve a specific auction
func GetAuctionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/auctions/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auction, err := database.GetAuction(id)
	if err != nil {
		log.Println("Error retrieving auction:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if auction == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(auction)
	if err != nil {
		log.Println("Error marshaling auction:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreateAuctionHandler handles the POST /auctions endpoint to create a new auction
func CreateAuctionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var auction models.Auction
	err := json.NewDecoder(r.Body).Decode(&auction)
	if err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.CreateAuction(&auction)
	if err != nil {
		log.Println("Error creating auction:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateAuctionHandler handles the PUT /auctions/{id} endpoint to update an existing auction
func UpdateAuctionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/auctions/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var auction models.Auction
	err = json.NewDecoder(r.Body).Decode(&auction)
	if err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if auction.ID != id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.UpdateAuction(&auction)
	if err != nil {
		log.Println("Error updating auction:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
