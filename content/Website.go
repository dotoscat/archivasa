package content

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/dotoscat/archivasa/theme"

	"github.com/gomarkdown/markdown"
)

// Website contains data about the site
type Website struct {
	Title string
	Pages []*Document
}

type WebsiteDocument struct {
	Website
	Document
	Content string
}

func NewWebsiteDocument(w *Website, d *Document) *WebsiteDocument {
	return &WebsiteDocument{*w, *d, "something"}
}

func NewWebsite(title string) *Website {
	return &Website{title, nil}
}

func (wd *WebsiteDocument) Render(outputDirectory string, document *template.Template) {
	wd.Content = string(markdown.ToHTML(wd.Markdown, nil, nil))
	pageOutputPath := filepath.Join(outputDirectory, wd.URL)
	pageOutput, err := os.Create(pageOutputPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := document.Execute(pageOutput, wd); err != nil {
		log.Fatal(err)
	}
}

// Export generates content
func Export(title string, cwd string) {
	site := NewWebsite(title)
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
		aPage := NewDocument(aPath, "here", "/pages/")
		aPage.Read()
		site.Pages[i] = aPage
		fmt.Println(aPage.URL)
	}
	salidaIndex, err := os.Create(path.Join(outputDirectory, "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	defer salidaIndex.Close()
	index := template.Must(template.ParseFiles("./theme/templates/basic.tmpl", "./theme/templates/index.tmpl"))
	if err := index.Execute(salidaIndex, site); err != nil {
		log.Fatal(err)
	}
	document := template.Must(template.ParseFiles("./theme/templates/basic.tmpl", "./theme/templates/document.tmpl"))
	for _, page := range site.Pages {
		if _, err := os.Stat(pagesDirectory); os.IsNotExist(err) {
			os.Mkdir(pagesDirectory, os.ModePerm)
		}
		websiteDocument := NewWebsiteDocument(site, page)
		websiteDocument.Render(outputDirectory, document)
	}
	theme.Copy(cwd, outputDirectory)
}

func (site *Website) String() string {
	return site.Title
}
