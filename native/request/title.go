package request

//go:generate ffjson $GOFILE

import "github.com/bsm/openrtb"

type Title struct {
	Length int               `json:"len"` // Maximum length of the text in the title element
	Ext    openrtb.Extension `json:"ext,omitempty"`
}

func (t *Title) Reset() {
	t.Length = 0
	if t.Ext != nil {
		t.Ext = t.Ext[:0]
	}
}
