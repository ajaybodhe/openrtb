package openrtb

//go:generate ffjson $GOFILE

// Private Marketplace Object
type Pmp struct {
	Private int       `json:"private_auction,omitempty"`
	Deals   []Deal    `json:"deals,omitempty"`
	Ext     Extension `json:"ext,omitempty"`
}

// PMP Deal
type Deal struct {
	ID               string    `json:"id,omitempty"` // Unique deal ID
	BidFloor         float64   `json:"bidfloor,omitempty"`
	BidFloorCurrency string    `json:"bidfloorcur,omitempty"` // Currency of bid floor
	WSeat            []string  `json:"wseat,omitempty"`       // Array of buyer seats allowed to bid on this Direct Deal.
	WAdvDomain       []string  `json:"wadomain,omitempty"`    // Array of advertiser domains allowed to bid on this Direct Deal
	AuctionType      int       `json:"at,omitempty"`          // Optional override of the overall auction type of the bid request, where 1 = First Price, 2 = Second Price Plus, 3 = the value passed in bidfloor is the agreed upon deal price. Additional auction types can be defined by the exchange.
	Ext              Extension `json:"ext,omitempty"`

	Seats []string `json:"seats,omitempty"` // DEPRECATED: kept for backwards compatibility
	Type  int      `json:"type,omitempty"`  // DEPRECATED: kept for backwards compatibility
}

func (d *Deal) Reset() {
	d.ID = ""
	d.BidFloor = 0.0
	d.BidFloorCurrency = ""
	if d.WSeat != nil {
		d.WSeat = d.WSeat[:0]
	}
	if d.WAdvDomain != nil {
		d.WAdvDomain = d.WAdvDomain[:0]
	}
	d.Type = 0
	d.AuctionType = 0
	if d.Ext != nil {
		d.Ext = d.Ext[:0]
	}
	if d.Seats != nil {
		d.Seats = d.Seats[:0]
	}
}

//var dealPool = sync.Pool{
//	New: func() interface{} {
//		return new(Deal)
//	},
//}
//
//func NewDeal() *Deal{
//	return dealPool.Get().(*Deal)
//}
//
//func FreeDeal(d *Deal) {
//	d.Reset()
//	dealPool.Put(d)
//}

func (p *Pmp) Reset() {
	p.Private = 0
	if p.Deals != nil {
		for i := 0; i < len(p.Deals); i++ {
			(&p.Deals[i]).Reset()
		}
		p.Deals = p.Deals[:0]
	}
	if p.Ext != nil {
		p.Ext = p.Ext[:0]
	}
}

//var pmpPool = sync.Pool{
//	New: func() interface{} {
//		return new(Pmp)
//	},
//}
//
//func NewPmp() *Pmp{
//	return pmpPool.Get().(*Pmp)
//}
//
//func FreePmp(p *Pmp) {
//	p.Reset()
//	pmpPool.Put(p)
//}

//type jsonDeal Deal

// MarshalJSON custom marshalling with normalization
//func (d *Deal) MarshalJSON() ([]byte, error) {
//	d.normalize()
//	return json.Marshal((*jsonDeal)(d))
//}
//
//// UnmarshalJSON custom unmarshalling with normalization
//func (d *Deal) UnmarshalJSON(data []byte) error {
//	var h jsonDeal
//	if err := json.Unmarshal(data, &h); err != nil {
//		return err
//	}
//
//	*d = (Deal)(h)
//	d.normalize()
//	return nil
//}
//
//func (d *Deal) normalize() {
//	if d.AuctionType == 0 {
//		d.AuctionType = 2
//	}
//}
