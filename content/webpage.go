package content

import "path/filepath"

// Webpage represents a internet document with its URL
// A webpage belongs to a website
type Webpage struct {
	*Website
	Url        string
	outputPath string
}

func NewWebpage(site *Website, url, outputPath string) *Webpage {
	return &Webpage{site, url, outputPath}
}

func (w *Webpage) BuildURL(prefix, base string) {
	w.Url = filepath.Join(prefix, base)
}

func (w *Webpage) BuildOutputPath(prefix, base string) {
	w.outputPath = filepath.Join(prefix, base)
}

func (w *Webpage) URL() string {
	return w.Url
}

func (w *Webpage) OutputPath() string {
	return w.outputPath
}
