package content

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
)

// Website contains data about the site
type Website struct {
	Title string
	Pages []*Document
}

func NewWebsite(title string) *Website {
	return &Website{title, nil}
}

// Export generates content
func Export(title string, cwd string) {
	site := NewWebsite(title)
	outputDirectory := path.Join(cwd, "output")

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
}

func (site *Website) String() string {
	return site.Title
}
