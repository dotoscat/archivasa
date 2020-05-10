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
	Title      string
	Pages      []*Document
	Postspages []*Postspage
	Theme      *theme.Theme
	cwd        string
}

func NewWebsite(title, cwd string) *Website {
	return &Website{title, nil, nil, nil, cwd}
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
	pagesDirectory := filepath.Join(outputDirectory, "pages")
	postsDirectory := path.Join(outputDirectory, "posts")
	MakeOutputdirIfNotExists(pagesDirectory)
	MakeOutputdirIfNotExists(postsDirectory)

	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.Mkdir(outputDirectory, os.ModePerm)
	}
	contentFolder := filepath.Join(cwd, "content")
	contentPagesDirectory := filepath.Join(contentFolder, "pages")
	contentPostsDirectory := filepath.Join(contentFolder, "posts")
	fmt.Println("content folder", contentFolder)
	site.Pages = GetDocumentsFromDir(contentPagesDirectory, outputDirectory, "/pages", site)
	posts := GetDocumentsFromDir(contentPostsDirectory, outputDirectory, "/posts", site)
	site.Postspages = CreatePostspages(posts, 2, outputDirectory, "/postspage", site)
	fmt.Println("postspages", len(site.Postspages))
	site.RenderPostspages("postspage")
	site.RenderPosts("document")
	site.RenderDocuments(site.Pages, "document")
	theme.Copy(outputDirectory)
}

func (site *Website) String() string {
	return site.Title
}

func (site *Website) RenderDocuments(documents []*Document, templateName string) {
	for _, document := range documents {
		document.Read()
		document.BuildContent()
		site.Theme.Render(templateName, document)
	}
}

func (site *Website) RenderPostspages(templateName string) {
	fmt.Println("Postspages: ", len(site.Postspages))
	for _, page := range site.Postspages {
		fmt.Println("Prev | next", page, page.Prev, page.Next)
		site.Theme.Render(templateName, page)
	}
}

func (site *Website) RenderPosts(templateName string) {
	fmt.Println("Postspages: ", len(site.Postspages))
	for _, page := range site.Postspages {
		fmt.Println("Posts of page: ", len(page.Posts))
		for _, webpage := range page.Posts {
			fmt.Println("Render webpage", webpage)
			site.Theme.Render(templateName, webpage)
		}
	}
}

func MakeOutputdirIfNotExists(outputDirectory string) {
	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.MkdirAll(outputDirectory, os.ModePerm)
	}
}
