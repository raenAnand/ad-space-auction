package database

import (
	"database/sql"
	"log"
	"fmt"
	"os"



	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/raenAnand/ad-space-auction/supply/api/models"

)


// ConnectDB establishes a connection to the MySQL database
func ConnectDB() (*sql.DB, error) {
	// Get the database IP address and port from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	if dbHost == "" {
		dbHost = "localhost" // Default to localhost
	}
	if dbPort == "" {
		dbPort = "3306" // Default port
	}

	dbURI := fmt.Sprintf("root:secret@tcp(%s:%s)/ad_auction", dbHost, dbPort)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database")
	return db, nil
}

// GetAllAdSpaces retrieves all ad spaces from the database
func GetAllAdSpaces() ([]models.AdSpace, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ad_spaces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var adSpaces []models.AdSpace
	for rows.Next() {
		var adSpace models.AdSpace
		err := rows.Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice)
		if err != nil {
			return nil, err
		}
		adSpaces = append(adSpaces, adSpace)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return adSpaces, nil
}

// GetAdSpace retrieves a specific ad space by ID from the database
func GetAdSpace(adSpaceID int) (*models.AdSpace, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM ad_spaces WHERE id = ?", adSpaceID)

	var adSpace models.AdSpace
	err = row.Scan(&adSpace.ID, &adSpace.Name, &adSpace.BasePrice)
	if err == sql.ErrNoRows {
		return nil, nil // Ad space not found
	}
	if err != nil {
		return nil, err
	}

	return &adSpace, nil
}

// CreateAdSpace creates a new ad space in the database
func CreateAdSpace(adSpace *models.AdSpace) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO ad_spaces(name, base_price) VALUES (?, ?)", adSpace.Name, adSpace.BasePrice)
	if err != nil {
		return err
	}

	// Get the ID of the newly inserted ad space
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	adSpace.ID = int(id)
	return nil
}

// UpdateAdSpace updates an existing ad space in the database
func UpdateAdSpace(adSpace *models.AdSpace) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE ad_spaces SET name = ?, base_price = ? WHERE id = ?", adSpace.Name, adSpace.BasePrice, adSpace.ID)
	if err != nil {
		return err
	}

	return nil
}
