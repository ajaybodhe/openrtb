package openrtb

//go:generate ffjson $GOFILE

import (
	"errors"
	"sync"
)

// Validation errors
var (
	ErrInvalidRespNoID       = errors.New("openrtb: response missing ID")
	ErrInvalidRespNoSeatBids = errors.New("openrtb: response missing seatbids")
)

// ID and at least one "seatbid” object is required, which contains a bid on at least one impression.
// Other attributes are optional since an exchange may establish default values.
// No-Bids on all impressions should be indicated as a HTTP 204 response.
// For no-bids on specific impressions, the bidder should omit these from the bid response.
type BidResponse struct {
	ID         string             `json:"id"`                   // Reflection of the bid request ID for logging purposes
	SeatBid    []SeatBid          `json:"seatbid"`              // Array of seatbid objects
	BidID      string             `json:"bidid,omitempty"`      // Optional response tracking ID for bidders
	Currency   string             `json:"cur,omitempty"`        // Bid currency
	CustomData string             `json:"customdata,omitempty"` // Encoded user features
	NBR        int                `json:"nbr,omitempty"`        // Reason for not bidding, where 0 = unknown error, 1 = technical error, 2 = invalid request, 3 = known web spider, 4 = suspected Non-Human Traffic, 5 = cloud, data center, or proxy IP, 6 = unsupported device, 7 = blocked publisher or site, 8 = unmatched user
	Ext        Extension          `json:"ext,omitempty"`        // Custom specifications in JSon
	TD         map[string]float64 `json:"-"`                    // time detail logging for local use
}

func (br *BidResponse) Reset() {
	br.NBR = 0
	br.CustomData = ""
	br.Currency = ""
	br.BidID = ""
	br.TD = nil
	if br.Ext != nil {
		br.Ext = br.Ext[:0]
	}
	br.ID = ""
	if br.SeatBid != nil {
		for i := 0; i < len(br.SeatBid); i++ {
			(&br.SeatBid[i]).Reset()
		}
		br.SeatBid = br.SeatBid[:0]
	}
}

var bidResponsePool = sync.Pool{
	New: func() interface{} {
		return new(BidRequest)
	},
}

func NewBidResponse() *BidResponse {
	return bidResponsePool.Get().(*BidResponse)
}

func FreeBidResponse(br *BidResponse) {
	if br == nil {
		return
	}
	br.Reset()
	bidResponsePool.Put(br)
}

// Validate required attributes
func (res *BidResponse) Validate() error {
	if res.ID == "" {
		return ErrInvalidRespNoID
	} else if len(res.SeatBid) == 0 {
		return ErrInvalidRespNoSeatBids
	}

	for _, sb := range res.SeatBid {
		if err := sb.Validate(); err != nil {
			return err
		}
	}

	return nil
}
