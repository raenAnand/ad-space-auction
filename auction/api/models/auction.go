package models

// Auction represents the details of an auction
type Auction struct {
	ID         int     `json:"id"`
	AdSpaceID  int     `json:"ad_space_id"`
	BidderID   int     `json:"bidder_id"`
	BidAmount  float64 `json:"bid_amount"`
	StartTime  string  `json:"start_time"`
	EndTime    string  `json:"end_time"`
}
