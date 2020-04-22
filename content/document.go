package content

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Document struct {
	Name     string
	Path     string
	Markdown []byte
	Url      string
}

func NewDocument(path string, prefixURL string) *Document {
	base := filepath.Base(path)
	spaces := regexp.MustCompile("\\s")
	dash := regexp.MustCompile("-|_")
	URLbase := strings.ReplaceAll(spaces.ReplaceAllString(base, "-"), ".md", ".html")
	URL := filepath.Join(prefixURL, URLbase)
	name := strings.TrimSuffix(dash.ReplaceAllString(base, " "), filepath.Ext(base))
	fmt.Println("name", name)
	return &Document{name, path, nil, URL}
}

func (post *Document) String() string {
	return post.Path
}

func (post *Document) Read() {
	file, err := os.Open(post.Path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	rawContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	//delimiter := regexp.MustCompile("---")
	content := string(rawContent)
	chunks := strings.Split(content, "---")
	post.Markdown = []byte(chunks[1])
	fmt.Println(post.Path)
	fmt.Println(post.Markdown)

}
