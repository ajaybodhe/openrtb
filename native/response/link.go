package response

//go:generate ffjson $GOFILE

import "github.com/bsm/openrtb"

type Link struct {
	URL           string            `json:"url"`                // Landing URL of the clickable link
	ClickTrackers []string          `json:"clicktrackers"`      // List of third-party tracker URLs to be fired on click of the URL
	FallbackURL   string            `json:"fallback,omitempty"` // Fallback URL for deeplink. To be used if the URL given in url is not supported by the device.
	Ext           openrtb.Extension `json:"ext,omitempty"`
}

func (l *Link) Reset() {
	l.URL = ""
	if l.ClickTrackers != nil {
		l.ClickTrackers = l.ClickTrackers[:0]
	}
	if l.Ext != nil {
		l.Ext = l.Ext[:0]
	}
	l.FallbackURL = ""
}
