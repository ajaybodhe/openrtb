package response

//go:generate ffjson $GOFILE

type Video struct {
	VASTTag string `json:"vasttag"` // VAST XML
}

func (v *Video) Reset() {
	v.VASTTag = ""
}
