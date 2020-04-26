package content

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/dotoscat/archivasa/theme"
)

// Website contains data about the site
type Website struct {
	Title string
	Pages []*Document
	url   string
	Theme *theme.Theme
	Cwd   string
}

func NewWebsite(title, cwd string) *Website {
	theme := theme.New(cwd)
	return &Website{title, nil, "/index.html", theme, cwd}
}

func (site *Website) BaseURL() string {
	return site.url
}

// Export generates content
func Export(title string, cwd string) {
	site := NewWebsite(title, cwd)
	theme := site.Theme
	outputDirectory := filepath.Join(cwd, "output")
	pagesDirectory := filepath.Join(outputDirectory, "pages")
	// postsDirectory := path.Join(outputDirectory, "posts")

	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.Mkdir(outputDirectory, os.ModePerm)
	}
	contentFolder := path.Join(cwd, "content")
	contentPagesDirectory := path.Join(contentFolder, "pages")
	fmt.Println("content folder", contentFolder)
	site.Pages = GetDocumentsFromDir(contentPagesDirectory)
	theme.Render("index", outputDirectory, site)
	site.RenderDocumentsToDir(site.Pages, "document", pagesDirectory, "/pages")
	theme.Copy(outputDirectory)
}

func (site *Website) String() string {
	return site.Title
}

func (site *Website) RenderDocumentsToDir(documents []*Document, templateName, outputDirectory, prefix string) {
	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.MkdirAll(outputDirectory, os.ModePerm)
	}
	for _, document := range documents {
		document.JoinPrefixURL(prefix)
		websiteDocument := NewWebsiteDocument(site, document)
		site.Theme.Render(templateName, outputDirectory, websiteDocument)
	}
}
