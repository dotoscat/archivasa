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
	Name       string
	Path       string
	OutputPath string
	Markdown   string
	URL        string
}

func NewDocument(path string, outputPath string, prefixURL string) *Document {
	base := filepath.Base(path)
	spaces := regexp.MustCompile("\\s")
	URLbase := strings.ReplaceAll(spaces.ReplaceAllString(base, "-"), ".md", ".html")
	URL := filepath.Join(prefixURL, URLbase)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	return &Document{name, path, outputPath, "", URL}
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
	post.Markdown = chunks[1]
	fmt.Println(post.Path)
	fmt.Println(post.Markdown)

}
