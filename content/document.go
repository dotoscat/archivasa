package content

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
)

var spaces *regexp.Regexp = regexp.MustCompile("\\s")
var dash *regexp.Regexp = regexp.MustCompile("-|_")

// Document is mainly generated from a markdown file with content and metadata
// A document is used as both post and page
type Document struct {
	Webpage
	Name     string
	path     string
	markdown []byte
	Content  string
}

func NewDocument(website *Website, path, base string) *Document {
	name := strings.TrimSuffix(dash.ReplaceAllString(base, " "), filepath.Ext(base))
	// urlName := strings.TrimSuffix(base, filepath.Ext(base)) + ".html"
	fmt.Println("name", name)
	return &Document{Webpage{website, ""}, name, path, nil, ""}
}

func (post *Document) String() string {
	return post.path
}

func (post *Document) Markdown() []byte {
	return post.markdown
}

func (post *Document) BuildContent() bool {
	if post.markdown == nil {
		return false
	}
	post.Content = string(markdown.ToHTML(post.markdown, nil, nil))
	return true
}

func (post *Document) Read() {
	file, err := os.Open(post.path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	rawContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	content := string(rawContent)
	chunks := strings.Split(content, "---")
	post.markdown = []byte(chunks[1])
	fmt.Println(post.path)
	fmt.Println(post.markdown)

}

func GetDocumentsFromDir(dirname, prefix string, website *Website) []*Document {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatalln(err)
	}
	documents := make([]*Document, len(files))
	for i, file := range files {
		if file.IsDir() {
			continue
		}
		path := filepath.Join(dirname, file.Name())
		urlName := strings.Replace(file.Name(), ".md", ".html", -1)
		documents[i] = NewDocument(website, path, file.Name())
		documents[i].BuildURL(prefix, urlName)
	}
	return documents
}
