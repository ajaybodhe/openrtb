package response

//go:generate ffjson $GOFILE

import "github.com/bsm/openrtb"

type Image struct {
	URL    string            `json:"url,omitempty"` // URL of the image asset
	Width  int               `json:"w,omitempty"`   // Width of the image in pixels
	Height int               `json:"h,omitempty"`   // Height of the image in pixels
	Ext    openrtb.Extension `json:"ext,omitempty"`
}

func (i *Image) Reset() {
	i.URL = ""
	i.Width = 0
	i.Height = 0
	if i.Ext != nil {
		i.Ext = i.Ext[:0]
	}
}
