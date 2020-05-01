package content

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dotoscat/archivasa/theme"
)

// Website contains data about the site
type Website struct {
	Title string
	Pages []*Document
	Theme *theme.Theme
	cwd   string
}

func NewWebsite(title, cwd string) *Website {
	return &Website{title, nil, nil, cwd}
}

func (site *Website) LoadTheme() *theme.Theme {
	site.Theme = theme.New(site.cwd)
	return site.Theme
}

// Export generates content
func Export(title string, cwd string) {
	site := NewWebsite(title, cwd)
	theme := site.LoadTheme()
	outputDirectory := filepath.Join(cwd, "output")
	// pagesDirectory := filepath.Join(outputDirectory, "pages")
	// postsDirectory := path.Join(outputDirectory, "posts")

	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.Mkdir(outputDirectory, os.ModePerm)
	}
	contentFolder := filepath.Join(cwd, "content")
	contentPagesDirectory := filepath.Join(contentFolder, "pages")
	fmt.Println("content folder", contentFolder)
	site.Pages = GetDocumentsFromDir(contentPagesDirectory, "/pages", site)
	index := NewWebpage(site, "/index.html")
	theme.Render("index", outputDirectory, index)
	site.RenderDocumentsToDir(site.Pages, "document", outputDirectory)
	theme.Copy(outputDirectory)
}

func (site *Website) String() string {
	return site.Title
}

func (site *Website) RenderDocumentsToDir(documents []*Document, templateName, outputDirectory string) {
	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.MkdirAll(outputDirectory, os.ModePerm)
	}
	for _, document := range documents {
		site.Theme.Render(templateName, outputDirectory, document)
	}
}
