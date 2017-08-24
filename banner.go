package openrtb

//go:generate ffjson $GOFILE

// The "banner" object must be included directly in the impression object if the impression offered
// for auction is display or rich media, or it may be optionally embedded in the video object to
// describe the companion banners available for the linear or non-linear video ad.  The banner
// object may include a unique identifier; this can be useful if these IDs can be leveraged in the
// VAST response to dictate placement of the companion creatives when multiple companion ad
// opportunities of the same size are available on a page.
type Banner struct {
	W        int       `json:"w,omitempty"`        // Width
	H        int       `json:"h,omitempty"`        // Height
	Format   []Format  `json:"format,omitempty"`   //Array of format objects representing the banner sizes permitted.
	WMax     int       `json:"wmax,omitempty"`     // Width maximum DEPRECATED
	HMax     int       `json:"hmax,omitempty"`     // Height maximum DEPRECATED
	WMin     int       `json:"wmin,omitempty"`     // Width minimum DEPRECATED
	HMin     int       `json:"hmin,omitempty"`     // Height minimum DEPRECATED
	ID       string    `json:"id,omitempty"`       // A unique identifier
	BType    []int     `json:"btype,omitempty"`    // Blocked creative types
	BAttr    []int     `json:"battr,omitempty"`    // Blocked creative attributes
	Pos      int       `json:"pos,omitempty"`      // Ad Position
	Mimes    []string  `json:"mimes,omitempty"`    // Whitelist of content MIME types supported
	TopFrame int       `json:"topframe,omitempty"` // Default: 0 ("1": Delivered in top frame, "0": Elsewhere)
	ExpDir   []int     `json:"expdir,omitempty"`   // Specify properties for an expandable ad
	Api      []int     `json:"api,omitempty"`      // List of supported API frameworks
	Ext      Extension `json:"ext,omitempty"`
}

func (bn *Banner) Reset() {
	bn.W = 0
	bn.H = 0
	if bn.Format != nil {
		for i := 0; i < len(bn.Format); i++ {
			(&bn.Format[i]).Reset()
		}
		bn.Format = bn.Format[:0]
	}
	bn.WMax = 0
	bn.HMax = 0
	bn.WMin = 0
	bn.HMin = 0
	bn.ID = ""
	bn.TopFrame = 0
	if bn.BType != nil {
		bn.BType = bn.BType[:0]
	}
	if bn.BAttr != nil {
		bn.BAttr = bn.BAttr[:0]
	}
	bn.Pos = 0
	if bn.Mimes != nil {
		bn.Mimes = bn.Mimes[:0]
	}
	if bn.ExpDir != nil {
		bn.ExpDir = bn.ExpDir[:0]
	}
	if bn.Api != nil {
		bn.Api = bn.Api[:0]
	}
	if bn.Ext != nil {
		bn.Ext = bn.Ext[:0]
	}
}

//var bannerPool = sync.Pool{
//	New: func() interface{} {
//		return new(Banner)
//	},
//}
//
//func NewBanner() *Banner{
//	return bannerPool.Get().(*Banner)
//}
//
//func FreeBanner(bn *Banner) {
//	bn.Reset()
//	bannerPool.Put(bn)
//}
