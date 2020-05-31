/*
archivasa - a static web generator, and only that
Copyright (C) 2020 Oscar Triano García

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package builder

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dotoscat/archivasa/content"
	"github.com/gomarkdown/markdown"
)

var spaces *regexp.Regexp = regexp.MustCompile("\\s")
var dash *regexp.Regexp = regexp.MustCompile("-|_")

// Document is mainly generated from a markdown file with content and metadata
// A document is used as both post and page
type Document struct {
	Webpage
	Name    string
	Content string
	Date    string
}

func NewDocument(website *Website, contentDocument *content.Document, prefix string) *Document {
	rawContent, err := ioutil.ReadFile(contentDocument.Path)
	if err != nil {
		log.Fatalln("error: ", contentDocument, " ; ", err)
	}
	baseName := filepath.Base(contentDocument.Path)
	URLBaseName := strings.TrimSuffix(baseName, ".md") + ".html"
	URL := filepath.Join("/", prefix, URLBaseName)
	name := strings.TrimSuffix(dash.ReplaceAllString(baseName, " "), ".md")
	date := contentDocument.Date
	dateString := fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day())
	contentString := string(rawContent)
	chunks := strings.Split(contentString, "---")
	if len(chunks) < 2 {
		log.Fatalln(contentDocument.Path, " has not '---' delimiter")
	}
	markdownChunk := chunks[1]
	content := string(markdown.ToHTML([]byte(markdownChunk), nil, nil))
	document := Document{Webpage{website, URL}, name, content, dateString}
	return &document
}

func (d *Document) String() string {
	return fmt.Sprintf("%v\n===\n%v\n%v\n", d.Name, d.URL, d.Date)
}

/*

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
*/
