package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/raenAnand/ad-space-auction/supply/database"
	"github.com/raenAnand/ad-space-auction/supply/api/models"

)

// GetAllAdSpacesHandler handles the GET /ad-spaces endpoint to retrieve all ad spaces
func GetAllAdSpacesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	adSpaces, err := database.GetAllAdSpaces()
	if err != nil {
		log.Println("Error retrieving ad spaces:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(adSpaces)
	if err != nil {
		log.Println("Error marshaling ad spaces:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetAdSpaceHandler handles the GET /ad-spaces/{id} endpoint to retrieve a specific ad space
func GetAdSpaceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/ad-spaces/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	adSpace, err := database.GetAdSpace(id)
	if err != nil {
		log.Println("Error retrieving ad space:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if adSpace == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(adSpace)
	if err != nil {
		log.Println("Error marshaling ad space:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreateAdSpaceHandler handles the POST /ad-spaces endpoint to create a new ad space
func CreateAdSpaceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var adSpace models.AdSpace
	err := json.NewDecoder(r.Body).Decode(&adSpace)
	if err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.CreateAdSpace(&adSpace)
	if err != nil {
		log.Println("Error creating ad space:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateAdSpaceHandler handles the PUT /ad-spaces/{id} endpoint to update an existing ad space
func UpdateAdSpaceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/ad-spaces/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var adSpace models.AdSpace
	err = json.NewDecoder(r.Body).Decode(&adSpace)
	if err != nil {
		log.Println("Error decoding request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if adSpace.ID != id {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.UpdateAdSpace(&adSpace)
	if err != nil {
		log.Println("Error updating ad space:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
