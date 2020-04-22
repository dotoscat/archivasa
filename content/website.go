package content

import (
	"fmt"
	"log"
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

	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.Mkdir(outputDirectory, os.ModePerm)
	}
	contentFolder := path.Join(cwd, "content")
	fmt.Println("content folder", contentFolder)
	pagePaths, error := PathsFromPages(contentFolder)
	site.Pages = make([]*Document, len(pagePaths))
	if error != nil {
		log.Fatalln(error)
	}
	for i, aPath := range pagePaths {
		aPage := NewDocument(aPath, "/pages/")
		aPage.Read()
		site.Pages[i] = aPage
	}
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
