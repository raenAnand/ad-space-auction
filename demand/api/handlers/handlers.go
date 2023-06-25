package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/raenAnand/ad-space-auction/demand/database"
	"github.com/raenAnand/ad-space-auction/demand/api/models"

)

// GetAllBidsHandler handles the GET /bids endpoint to retrieve all bids
func GetAllBidsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bids, err := database.GetAllBids()
	if err != nil {
		log.Println("Error retrieving bids:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(bids)
	if err != nil {
		log.Println("Error marshaling bids:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetBidHandler handles the GET /bids/{id} endpoint to retrieve a specific bid
func GetBidHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/bids/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bid, err := database.GetBid(id)
	if err != nil {
		log.Println("Error retrieving bid:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if bid == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(bid)
	if err != nil {
		log.Println("Error marshaling bid:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreateBidHandler handles the POST /bids endpoint to create a new bid
func CreateBidHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var bid model.Bid
	err := json.NewDecoder(r.Body).Decode(&bid)
	if err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.CreateBid(&bid)
	if err != nil {
		log.Println("Error creating bid:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateBidHandler handles the PUT /bids/{id} endpoint to update an existing bid
func UpdateBidHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/bids/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var bid model.Bid
	err = json.NewDecoder(r.Body).Decode(&bid)
	if err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if bid.ID != id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.UpdateBid(&bid)
	if err != nil {
		log.Println("Error updating bid:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
