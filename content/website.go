package content

import (
	"fmt"
	"os"
	"path"

	"github.com/dotoscat/archivasa/theme"
)

// Website contains data about the site
type Website struct {
	Title string
	Pages []*Document
	url   string
}

func NewWebsite(title string) *Website {
	return &Website{title, nil, "/index.html"}
}

func (site *Website) URL() string {
	return site.url
}

// Export generates content
func Export(title string, cwd string) {
	site := NewWebsite(title)
	theme := theme.New(cwd)
	outputDirectory := path.Join(cwd, "output")
	pagesDirectory := path.Join(outputDirectory, "pages")
	// postsDirectory := path.Join(outputDirectory, "posts")

	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.Mkdir(outputDirectory, os.ModePerm)
	}
	contentFolder := path.Join(cwd, "content")
	contentPagesDirectory := path.Join(contentFolder, "pages")
	fmt.Println("content folder", contentFolder)
	site.Pages = GetDocumentsFromDir(contentPagesDirectory)
	theme.Render("index", outputDirectory, site)
	for _, page := range site.Pages {
		if _, err := os.Stat(pagesDirectory); os.IsNotExist(err) {
			os.Mkdir(pagesDirectory, os.ModePerm)
		}
		websiteDocument := NewWebsiteDocument(site, page)
		theme.Render("document", outputDirectory, websiteDocument)
	}
	theme.Copy(outputDirectory)
}

func (site *Website) String() string {
	return site.Title
}
