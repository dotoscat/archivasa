package builder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/dotoscat/archivasa/config"
	"github.com/dotoscat/archivasa/content"
	"github.com/dotoscat/archivasa/theme"
)

func Run(config config.Config, content *content.Content, theme *theme.Theme) {
	outputPath := filepath.Join(config.Cwd, "output")
	pagesPath := filepath.Join(config.Cwd, "output", "pages")
	postsPath := filepath.Join(config.Cwd, "output", "posts")
	MakeOutputdirIfNotExists(pagesPath)
	MakeOutputdirIfNotExists(postsPath)
	fmt.Println("nPostspages", len(content.Posts), config.PostsPerPage)
	nPostspages := len(content.Posts) / config.PostsPerPage
	if len(content.Posts)%config.PostsPerPage != 0 {
		nPostspages++
	}
	fmt.Println("nPostspages(2)", nPostspages)
	website := NewWebsite(config.Title, len(content.Pages), nPostspages)
	// Build pages
	for i := 0; i < len(website.Pages); i++ {
		website.Pages[i] = NewDocument(website, content.Pages[i], "pages")
	}
	fmt.Println(website.Pages)
	posts := make([]*Document, len(content.Posts))
	for i := 0; i < len(posts); i++ {
		posts[i] = NewDocument(website, content.Posts[i], "posts")
	}
	fmt.Println(posts)
	FillPostspages(website, posts, config.PostsPerPage)
	documentTemplate, ok := theme.Templates["document"]
	if !ok {
		log.Fatalln(documentTemplate, " do not exists")
	}
	for _, page := range website.Pages {
		Render(documentTemplate, page, outputPath)
	}
	//fmt.Println(website.Postspages)
	//fmt.Println(website.Postspages[0].Posts)
	//fmt.Println(config)
	// Build posts
	// Buils postspages
	// Render posts
	// Render postspages

}

func Render(template *template.Template, webpage Urler, outputDirectory string) {
	outputPath := filepath.Join(outputDirectory, webpage.Url())
	file, err := os.Create(outputPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if template == nil {
		log.Fatalf("%v template is nil", template)
	}
	if err := template.Execute(file, webpage); err != nil {
		log.Fatal(err)
	}
}

func MakeOutputdirIfNotExists(outputDirectory string) {
	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.MkdirAll(outputDirectory, os.ModePerm)
	}
}
