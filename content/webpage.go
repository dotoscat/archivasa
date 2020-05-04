package content

import "path/filepath"

// Webpage represents a internet document with its URL
// A webpage belongs to a website
type Webpage struct {
	*Website
	Url string
}

func NewWebpage(site *Website, url string) *Webpage {
	return &Webpage{site, url}
}

func (w *Webpage) BuildURL(prefix, base string) {
	w.Url = filepath.Join(prefix, base)
}

func (w *Webpage) URL() string {
	return w.Url
}
