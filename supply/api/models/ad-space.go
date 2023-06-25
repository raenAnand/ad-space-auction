package models

// AdSpace represents an ad space available for auction
type AdSpace struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	BasePrice  float64 `json:"base_price"`
}

// Bidder represents a bidder interested in bidding for ad spaces
type Bidder struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Auction represents the details of an auction
type Auction struct {
	ID         int     `json:"id"`
	AdSpaceID  int     `json:"ad_space_id"`
	BidderID   int     `json:"bidder_id"`
	BidAmount  float64 `json:"bid_amount"`
	StartTime  string  `json:"start_time"`
	EndTime    string  `json:"end_time"`
}
