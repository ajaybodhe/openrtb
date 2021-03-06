package response

//go:generate ffjson $GOFILE

import "github.com/bsm/openrtb"

type Data struct {
	Label string            `json:"label,omitempty"` // The optional formatted string name of the data type to be displayed
	Value string            `json:"value"`           // The formatted string of data to be displayed. Can contain a formatted value such as “5 stars” or “$10” or “3.4 stars out of 5”
	Ext   openrtb.Extension `json:"ext,omitempty"`
}

func (d *Data) Reset() {
	d.Label = ""
	d.Value = ""
	if d.Ext != nil {
		d.Ext = d.Ext[:0]
	}
}
