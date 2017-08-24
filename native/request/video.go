package request

//go:generate ffjson $GOFILE

import "github.com/bsm/openrtb"

// TODO unclear if its the same as imp.video https://github.com/openrtb/OpenRTB/issues/26
type Video struct {
	Mimes       []string          `json:"mimes,omitempty"`       // Whitelist of content MIME types supported
	MinDuration int               `json:"minduration,omitempty"` // Minimum video ad duration in seconds
	MaxDuration int               `json:"maxduration,omitempty"` // Maximum video ad duration in seconds
	Protocols   []int             `json:"protocols,omitempty"`   // Video bid response protocols
	Ext         openrtb.Extension `json:"ext,omitempty"`
}

func (v *Video) Reset() {
	v.MinDuration = 0
	v.MaxDuration = 0
	if v.Mimes != nil {
		v.Mimes = v.Mimes[:0]
	}
	if v.Protocols != nil {
		v.Protocols = v.Protocols[:0]
	}
	if v.Ext != nil {
		v.Ext = v.Ext[:0]
	}
}
