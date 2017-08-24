package openrtb

import "sync"

//go:generate ffjson $GOFILE

type Inventory struct {
	ID            string     `json:"id,omitempty"` // ID on the exchange
	Name          string     `json:"name,omitempty"`
	Domain        string     `json:"domain,omitempty"`
	Cat           []string   `json:"cat,omitempty"`           // Array of IAB content categories
	SectionCat    []string   `json:"sectioncat,omitempty"`    // Array of IAB content categories for subsection
	PageCat       []string   `json:"pagecat,omitempty"`       // Array of IAB content categories for page
	PrivacyPolicy *int       `json:"privacypolicy,omitempty"` // Default: 1 ("1": has a privacy policy)
	Publisher     *Publisher `json:"publisher,omitempty"`     // Details about the Publisher
	Content       *Content   `json:"content,omitempty"`       // Details about the Content
	Keywords      string     `json:"keywords,omitempty"`      // Comma separated list of keywords about the site.
	Ext           Extension  `json:"ext,omitempty"`
}

// GetPrivacyPolicy returns the privacy policy value
func (a *Inventory) GetPrivacyPolicy() int {
	if a.PrivacyPolicy != nil {
		return *a.PrivacyPolicy
	}
	return 1
}

// An "app" object should be included if the ad supported content is part of a mobile application
// (as opposed to a mobile website).  A bid request must not contain both an "app" object and a
// "site" object.
type App struct {
	Inventory
	Bundle   string `json:"bundle,omitempty"`   // App bundle or package name
	StoreURL string `json:"storeurl,omitempty"` // App store URL for an installed app
	Ver      string `json:"ver,omitempty"`      // App version
	Paid     int    `json:"paid,omitempty"`     // "1": Paid, "2": Free
}

func (app *App) Reset() {
	if app.Ext != nil {
		app.Ext = app.Ext[:0]
	}
	if app.Content != nil {
		app.Content.Reset()
	}
	if app.Publisher != nil {
		app.Publisher.Reset()
	}
	if app.SectionCat != nil {
		app.SectionCat = app.SectionCat[:0]
	}
	if app.PageCat != nil {
		app.PageCat = app.PageCat[:0]
	}
	if app.Cat != nil {
		app.Cat = app.Cat[:0]
	}
	app.PrivacyPolicy = nil
	app.Keywords = ""
	app.Domain = ""
	app.Name = ""
	app.ID = ""
	app.Bundle = ""
	app.StoreURL = ""
	app.Ver = ""
	app.Paid = 0
}

var appPool = sync.Pool{
	New: func() interface{} {
		return new(App)
	},
}

func NewApp() *App {
	return appPool.Get().(*App)
}

func FreeApp(app *App) {
	app.Reset()
	appPool.Put(app)
}

// A site object should be included if the ad supported content is part of a website (as opposed to
// an application).  A bid request must not contain both a site object and an app object.
type Site struct {
	Inventory
	Page   string `json:"page,omitempty"`   // URL of the page
	Ref    string `json:"ref,omitempty"`    // Referrer URL
	Search string `json:"search,omitempty"` // Search string that caused naviation
	Mobile int    `json:"mobile,omitempty"` // Mobile ("1": site is mobile optimised)
}

func (s *Site) Reset() {
	if s.Ext != nil {
		s.Ext = s.Ext[:0]
	}
	if s.Content != nil {
		s.Content.Reset()
	}
	if s.Publisher != nil {
		s.Publisher.Reset()
	}
	if s.SectionCat != nil {
		s.SectionCat = s.SectionCat[:0]
	}
	if s.PageCat != nil {
		s.PageCat = s.PageCat[:0]
	}
	if s.Cat != nil {
		s.Cat = s.Cat[:0]
	}
	s.PrivacyPolicy = nil
	s.Keywords = ""
	s.Domain = ""
	s.Name = ""
	s.ID = ""
	s.Page = ""
	s.Ref = ""
	s.Mobile = 0
	s.Search = ""
}

var sitePool = sync.Pool{
	New: func() interface{} {
		return new(Site)
	},
}

func NewSite() *Site {
	return sitePool.Get().(*Site)
}

func FreeSite(s *Site) {
	s.Reset()
	sitePool.Put(s)
}
