package request

//go:generate ffjson $GOFILE

import "github.com/bsm/openrtb"

type ImageTypeID int

const (
	ImageTypeIcon ImageTypeID = 1 // Icon image
	ImageTypeLogo ImageTypeID = 2 // Logo image for the brand/app
	ImageTypeMain ImageTypeID = 3 // Large image preview for the ad
)

type Image struct {
	TypeID ImageTypeID `json:"type,omitempty"` // Type ID of the image element supported by the publisher

	Width     int `json:"w,omitempty"`    // Width of the image in pixels
	WidthMin  int `json:"wmin,omitempty"` // The minimum requested width of the image in pixels
	Height    int `json:"h,omitempty"`    // Height of the image in pixels
	HeightMin int `json:"hmin,omitempty"` // The minimum requested height of the image in pixels
	// Either h/w or hmin/wmin should be transmitted. If only h/w is included, it
	// should be considered an exact requirement
	Mimes []string          `json:"mimes,omitempty"` // Whitelist of content MIME types supported
	Ext   openrtb.Extension `json:"ext,omitempty"`
}

func (i *Image) Reset() {
	i.TypeID = 0
	i.Width = 0
	i.WidthMin = 0
	i.Height = 0
	i.HeightMin = 0
	if i.Mimes != nil {
		i.Mimes = i.Mimes[:0]
	}
	if i.Ext != nil {
		i.Ext = i.Ext[:0]
	}
}
