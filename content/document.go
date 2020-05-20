package content

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
)

var spaces *regexp.Regexp = regexp.MustCompile("\\s")
var dash *regexp.Regexp = regexp.MustCompile("-|_")

// Document is mainly generated from a markdown file with content and metadata
// A document is used as both post and page
type Document struct {
	Webpage
	Name     string
	srcPath  string
	markdown []byte
	Content  string
	Date     string
	date     time.Time
}

func NewDocument(website *Website, srcPath, base string) *Document {
	name := strings.TrimSuffix(dash.ReplaceAllString(base, " "), filepath.Ext(base))
	// urlName := strings.TrimSuffix(base, filepath.Ext(base)) + ".html"
	fmt.Println("name", name)
	return &Document{Webpage{website, "", ""}, name, srcPath, nil, "", "", time.Now()}
}

func (post *Document) String() string {
	return post.srcPath
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

func (post *Document) readDate(date string) {
	post.Date = date
	var year int
	var month time.Month
	var day int
	fmt.Sscanf(date, "%d-%d-%d", &year, &month, &day)
	post.date = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func (post *Document) readMeta(meta string) {
	lines := strings.Split(meta, "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		for i, field := range fields {
			fields[i] = strings.TrimSpace(field)
		}
		switch fields[0] {
		case "date":
			post.readDate(fields[1])
		}
	}

}

func (post *Document) Read() {
	fmt.Println("Open srcPath for reading", post.srcPath)
	file, err := os.Open(post.srcPath)
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
	if len(chunks) < 2 {
		log.Println(post.srcPath, "has not '---' delimiter")
		return
	}
	post.readMeta(chunks[0])
	post.markdown = []byte(chunks[1])
	fmt.Println(post.markdown)
}

func GetDocumentsFromDir(dirname, outputDir, prefix string, website *Website) DocumentSlice {
	fmt.Println("Get documents from:", dirname)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatalln(err)
	}
	documents := make(DocumentSlice, len(files))
	for i, file := range files {
		if file.IsDir() {
			continue
		}
		srcPath := filepath.Join(dirname, file.Name())
		urlName := strings.Replace(file.Name(), ".md", ".html", -1)
		outputPath := filepath.Join(outputDir, prefix)
		documents[i] = NewDocument(website, srcPath, file.Name())
		documents[i].BuildURL(prefix, urlName)
		documents[i].BuildOutputPath(outputPath, urlName)
		documents[i].Read()
	}
	return documents
}

type DocumentSlice []*Document

func (d DocumentSlice) Len() int {
	return len(d)
}

func (d DocumentSlice) Less(i, j int) bool {
	return d[i].date.After(d[j].date)
}

func (d DocumentSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d DocumentSlice) String() (output string) {
	for i, document := range d {
		output += fmt.Sprintln(i, document.date)
	}
	return
}
