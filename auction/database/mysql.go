package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/raenAnand/ad-space-auction/auction/api/models"
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

// GetAuction retrieves a specific auction by ID from the database
func GetAuction(auctionID int) (*models.Auction, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM auctions WHERE id = ?", auctionID)

	var auction models.Auction
	err = row.Scan(&auction.ID, &auction.AdSpaceID, &auction.BidderID, &auction.BidAmount, &auction.StartTime, &auction.EndTime)
	if err == sql.ErrNoRows {
		return nil, nil // Auction not found
	}
	if err != nil {
		return nil, err
	}

	return &auction, nil
}

// CreateAuction creates a new auction in the database
func CreateAuction(auction *models.Auction) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO auctions(ad_space_id, bidder_id, bid_amount, start_time, end_time) VALUES (?, ?, ?, ?, ?)", auction.AdSpaceID, auction.BidderID, auction.BidAmount, auction.StartTime, auction.EndTime)
	if err != nil {
		return err
	}

	// Get the ID of the newly inserted auction
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	auction.ID = int(id)
	return nil
}

// UpdateAuction updates an existing auction in the database
func UpdateAuction(auction *models.Auction) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE auctions SET ad_space_id = ?, bidder_id = ?, bid_amount = ?, start_time = ?, end_time = ? WHERE id = ?", auction.AdSpaceID, auction.BidderID, auction.BidAmount, auction.StartTime, auction.EndTime, auction.ID)
	if err != nil {
		return err
	}

	return nil
}
