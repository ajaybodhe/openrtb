package request

//go:generate ffjson $GOFILE

import "github.com/bsm/openrtb"

type Title struct {
	Length int               `json:"len"` // Maximum length of the text in the title element
	Ext    openrtb.Extension `json:"ext,omitempty"`
}
