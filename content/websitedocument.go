package content

import "github.com/gomarkdown/markdown"

type WebsiteDocument struct {
	Website
	Document
	Content string
}

func NewWebsiteDocument(w *Website, d *Document) *WebsiteDocument {
	content := string(markdown.ToHTML(d.Markdown, nil, nil))
	return &WebsiteDocument{*w, *d, content}
}

func (ws *WebsiteDocument) BaseURL() string {
	return ws.Document.BaseURL()
}
