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

func NewDocument(path, base string) *Document {
	// spaces := regexp.MustCompile("\\s") // TODO: Use this line for exporting
	dash := regexp.MustCompile("-|_")
	name := strings.TrimSuffix(dash.ReplaceAllString(base, " "), filepath.Ext(base))
	fmt.Println("name", name)
	return &Document{name, path, nil, ""}
}

func GetDocumentsFromDir(dirname string) []*Document {
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
		documents[i] = NewDocument(path, file.Name())
	}
	return documents
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
