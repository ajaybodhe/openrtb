package response

//go:generate ffjson $GOFILE

import "github.com/bsm/openrtb"

// Corresponds to the Asset Object in the request. The main container object for
// each asset requested or supported by Exchange on behalf of the rendering
// client. Any object that is required is to be flagged as such. Only one of the
// {title,img,video,data} objects should be present in each object. All others
// should be null/absent. The id is to be unique within the AssetObject array so
// that the response can be aligned.
type Asset struct {
	ID       int               `json:"id"`                 // Unique asset ID, assigned by exchange, must match one of the asset IDs in request
	Required int               `json:"required,omitempty"` // Set to 1 if asset is required
	Title    *Title            `json:"title,omitempty"`    // Title object for title assets
	Image    *Image            `json:"img,omitempty"`      // Image object for image assets
	Video    *Video            `json:"video,omitempty"`    // Video object for video assets
	Data     *Data             `json:"data,omitempty"`     // Data object for brand name, description, ratings, prices etc.
	Link     *Link             `json:"link,omitempty"`     // Link object for call to actions. The link object applies if the asset item is activated (clicked)
	Ext      openrtb.Extension `json:"ext,omitempty"`
}

func (a *Asset) Reset() {
	a.ID = 0
	a.Required = 0
	if a.Title != nil {
		a.Title.Reset()
	}
	if a.Image != nil {
		a.Image.Reset()
	}
	if a.Ext != nil {
		a.Ext = a.Ext[:0]
	}
	if a.Link != nil {
		a.Link.Reset()
	}
	if a.Data != nil {
		a.Data.Reset()
	}
	if a.Video != nil {
		a.Video.Reset()
	}
}
