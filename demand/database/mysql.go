package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/raenAnand/ad-space-auction/demand/api/models"
)

// ConnectDB establishes a connection to the MySQL database
func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/database")
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

// GetAllBids retrieves all bids from the database
func GetAllBids() ([]model.Bid, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM bids")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bids []model.Bid
	for rows.Next() {
		var bid model.Bid
		err := rows.Scan(&bid.ID, &bid.AdSpaceID, &bid.BidderID, &bid.BidAmount)
		if err != nil {
			return nil, err
		}
		bids = append(bids, bid)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bids, nil
}

// GetBid retrieves a specific bid by ID from the database
func GetBid(bidID int) (*model.Bid, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM bids WHERE id = ?", bidID)

	var bid model.Bid
	err = row.Scan(&bid.ID, &bid.AdSpaceID, &bid.BidderID, &bid.BidAmount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Bid not found
		}
		return nil, err
	}

	return &bid, nil
}

// CreateBid creates a new bid in the database
func CreateBid(bid *model.Bid) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO bids (ad_space_id, bidder_id, bid_amount) VALUES (?, ?, ?)",
		bid.AdSpaceID, bid.BidderID, bid.BidAmount)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBid updates an existing bid in the database
func UpdateBid(bid *model.Bid) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE bids SET ad_space_id = ?, bidder_id = ?, bid_amount = ? WHERE id = ?",
		bid.AdSpaceID, bid.BidderID, bid.BidAmount, bid.ID)
	if err != nil {
		return err
	}

	return nil
}
