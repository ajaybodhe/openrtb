package openrtb

//go:generate ffjson $GOFILE

// This object describes the content in which the impression will appear, which may be syndicated or nonsyndicated
// content. This object may be useful when syndicated content contains impressions and does
// not necessarily match the publisher's general content. The exchange might or might not have
// knowledge of the page where the content is running, as a result of the syndication method. For
// example might be a video impression embedded in an iframe on an unknown web property or device.
type Content struct {
	ID                 string    `json:"id,omitempty"`                 // ID uniquely identifying the content.
	Episode            int       `json:"episode,omitempty"`            // Episode number (typically applies to video content).
	Title              string    `json:"title,omitempty"`              // Content title.
	Series             string    `json:"series,omitempty"`             // Content series.
	Season             string    `json:"season,omitempty"`             // Content season.
	Artist             string    `json:"artist,omitempty"`             // Artist credited with the content.
	Genre              string    `json:"genre,omitempty"`              // Genre that best describes the content
	Album              string    `json:"album,omiyempty"`              // Album to which the content belongs; typically for audio.
	ISRC               string    `json:"isrc,omitempty"`               // International Standard Recording Code conforming to ISO - 3901.
	Producer           *Producer `json:"producer,omitempty"`           // The producer.
	URL                string    `json:"url,omitempty"`                // URL of the content, for buy-side contextualization or review.
	Cat                []string  `json:"cat,omitempty"`                // Array of IAB content categories that describe the content.
	ProdQuality        int       `json:"prodq,omitempty"`              // Production quality per IAB's classification.
	VideoQuality       int       `json:"videoquality,omitempty"`       // Video quality per IAB's classification.
	Context            int       `json:"context,omitempty"`            // Type of content (game, video, text, etc.).
	ContentRating      string    `json:"contentrating,omitempty"`      // Content rating (e.g., MPAA).
	UserRating         string    `json:"userrating,omitempty"`         // User rating of the content (e.g., number of stars, likes, etc.).
	QAGMediaRating     int       `json:"qagmediarating,omitempty"`     // Media rating per QAG guidelines.
	Keywords           string    `json:"keywords,omitempty"`           // Comma separated list of keywords describing the content.
	LiveStream         int       `json:"livestream,omitempty"`         // 0 = not live, 1 = content is live (e.g., stream, live blog).
	SourceRelationship int       `json:"sourcerelationship,omitempty"` // 0 = indirect, 1 = direct.
	Len                int       `json:"len,omitempty"`                // Length of content in seconds; appropriate for video or audio.
	Language           string    `json:"language,omitempty"`           // Content language using ISO-639-1-alpha-2.
	Embeddable         int       `json:"embeddable,omitempty"`         // Indicator of whether or not the content is embeddable (e.g., an embeddable video player), where 0 = no, 1 = yes.
	Data               []Data    `json:"data,omitempty"`               // Additional content data.
	Ext                Extension `json:"ext,omitempty"`
}

func (c *Content) Reset() {
	if c.Data != nil {
		for i := 0; i < len(c.Data); i++ {
			(&c.Data[i]).Reset()
		}
		c.Data = c.Data[:0]
	}
	if c.Ext != nil {
		c.Ext = c.Ext[:0]
	}
	c.Embeddable = 0
	c.Language = ""
	c.Len = 0
	c.SourceRelationship = 0
	c.LiveStream = 0
	c.Keywords = ""
	c.QAGMediaRating = 0
	c.UserRating = ""
	c.ContentRating = ""
	c.VideoQuality = 0
	c.ProdQuality = 0
	if c.Cat != nil {
		c.Cat = c.Cat[:0]
	}
	c.URL = ""
	if c.Producer != nil {
		c.Producer.Reset()
	}
	c.ISRC = ""
	c.Genre = ""
	c.Artist = ""
	c.Album = ""
	c.Series = ""
	c.Season = ""
	c.Title = ""
	c.Episode = 0
	c.ID = ""
}
