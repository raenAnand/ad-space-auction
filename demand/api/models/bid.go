package model

// Bid represents a bid for an ad space in the auction
type Bid struct {
	ID         int     `json:"id"`
	AdSpaceID  int     `json:"ad_space_id"`
	BidderID   int     `json:"bidder_id"`
	BidAmount  float64 `json:"bid_amount"`
}
