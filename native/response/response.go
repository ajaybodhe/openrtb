package response

//go:generate ffjson $GOFILE

import (
	"github.com/bsm/openrtb"
	"sync"
)

// The native object is the top level JSON object which identifies a native response
type Response struct {
	Ver         openrtb.StringOrNumber `json:"ver,omitempty"`         // Version of the Native Markup
	Assets      []Asset                `json:"assets"`                // An array of Asset Objects
	Link        Link                   `json:"link"`                  // Destination Link. This is default link object for the ad
	ImpTrackers []string               `json:"imptrackers,omitempty"` // Array of impression tracking URLs, expected to return a 1x1 image or 204 response
	JSTracker   string                 `json:"jstracker,omitempty"`   // Optional JavaScript impression tracker. This is a valid HTML, Javascript is already wrapped in <script> tags. It should be executed at impression time where it can be supported
	Ext         openrtb.Extension      `json:"ext,omitempty"`
}

func (nr *Response) Reset() {
	(&nr.Link).Reset()
	if nr.ImpTrackers != nil {
		nr.ImpTrackers = nr.ImpTrackers[:0]
	}
	nr.JSTracker = ""
	if nr.Assets != nil {
		for i := 0; i < len(nr.Assets); i++ {
			(&nr.Assets[i]).Reset()
		}
		nr.Assets = nr.Assets[:0]
	}
	if nr.Ext != nil {
		nr.Ext = nr.Ext[:0]
	}
	nr.Ver = ""
}

var nativeRequestPool = sync.Pool{
	New: func() interface{} {
		return new(Response)
	},
}

func NewNativeRequest() *Response {
	return nativeRequestPool.Get().(*Response)
}

func FreeNativeRequest(nr *Response) {
	nr.Reset()
	nativeRequestPool.Put(nr)
}
