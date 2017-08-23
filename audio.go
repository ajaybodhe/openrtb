package openrtb

//go:generate ffjson $GOFILE

import (
	"encoding/json"
	"errors"
)

// Validation errors
var (
	ErrInvalidAudioNoMimes = errors.New("openrtb: audio has no mimes")
)

// The "audio" object must be included directly in the impression object
type Audio struct {
	Mimes         []string  `json:"mimes"`                 // Content MIME types supported.
	MinDuration   int       `json:"minduration,omitempty"` // Minimum video ad duration in seconds
	MaxDuration   int       `json:"maxduration,omitempty"` // Maximum video ad duration in seconds
	Protocols     []int     `json:"protocols,omitempty"`   // Video bid response protocols
	StartDelay    int       `json:"startdelay,omitempty"`  // Indicates the start delay in seconds
	Sequence      int       `json:"sequence,omitempty"`    // Default: 1
	BAttr         []int     `json:"battr,omitempty"`       // Blocked creative attributes
	MaxExtended   int       `json:"maxextended,omitempty"` // Maximum extended video ad duration
	MinBitrate    int       `json:"minbitrate,omitempty"`  // Minimum bit rate in Kbps
	MaxBitrate    int       `json:"maxbitrate,omitempty"`  // Maximum bit rate in Kbps
	Delivery      []int     `json:"delivery,omitempty"`    // List of supported delivery methods
	CompanionAd   []Banner  `json:"companionad,omitempty"`
	API           []int     `json:"api,omitempty"`
	CompanionType []int     `json:"companiontype,omitempty"`
	MaxSequence   int       `json:"maxseq,omitempty"`   // The maximumnumber of ads that canbe played in an ad pod.
	Feed          int       `json:"feed,omitempty"`     // Type of audio feed.
	Stitched      int       `json:"stitched,omitempty"` // Indicates if the ad is stitched with audio content or delivered independently
	NVol          int       `json:"nvol,omitempty"`     // Volume normalization mode.
	Ext           Extension `json:"ext,omitempty"`
}

func (au *Audio) Reset() {
	if au.Ext != nil {
		au.Ext = au.Ext[:0]
	}
	if au.API != nil {
		au.API = au.API[:0]
	}
	if au.CompanionType != nil {
		au.CompanionType = au.CompanionType[:0]
	}
	if au.BAttr != nil {
		au.BAttr = au.BAttr[:0]
	}
	if au.CompanionAd != nil {
		for i:=0; i<len(au.CompanionAd); i ++ {
			(&au.CompanionAd[i]).Reset()
		}
		au.CompanionAd = au.CompanionAd[:0]
	}
	if au.Delivery != nil {
		au.Delivery = au.Delivery[:0]
	}
	au.Feed = 0
	au.Stitched = 0
	au.NVol = 0
	au.StartDelay = 0
	au.Sequence = 0
	au.MaxBitrate = 0
	au.MaxExtended = 0
	au.MaxSequence = 0
	if au.Protocols != nil {
		au.Protocols = au.Protocols[:0]
	}
	au.MinBitrate = 0
	au.MaxDuration = 0
	if au.Mimes != nil {
		au.Mimes = au.Mimes[:0]
	}
}

//var audioPool = sync.Pool{
//	New: func() interface{} {
//		return new(Audio)
//	},
//}
//
//func NewAudio() *Audio{
//	return audioPool.Get().(*Audio)
//}
//
//func FreeAudio(au *Audio) {
//	au.Reset()
//	audioPool.Put(au)
//}

type jsonAudio Audio

// Validates the object
func (a *Audio) Validate() error {
	if len(a.Mimes) == 0 {
		return ErrInvalidAudioNoMimes
	}
	return nil
}

// MarshalJSON custom marshalling with normalization
func (a *Audio) MarshalJSON() ([]byte, error) {
	a.normalize()
	return json.Marshal((*jsonAudio)(a))
}

// UnmarshalJSON custom unmarshalling with normalization
func (a *Audio) UnmarshalJSON(data []byte) error {
	var h jsonAudio
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}

	*a = (Audio)(h)
	a.normalize()
	return nil
}

func (a *Audio) normalize() {
	if a.Sequence == 0 {
		a.Sequence = 1
	}
}
