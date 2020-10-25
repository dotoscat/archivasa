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

	"github.com/dotoscat/archivasa/pkg/content"
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
	fmt.Println("Whut??!!")
	rawContent, err := ioutil.ReadFile(contentDocument.Path)
	if err != nil {
		log.Fatalln("error: ", contentDocument, " ; ", err)
	}
	baseName := filepath.Base(contentDocument.Path)
	URLBaseName := strings.TrimSuffix(baseName, ".md") + ".html"
	URL := fmt.Sprintf("/%v/%v", prefix, URLBaseName)
	var name string
	if contentDocument.Title != "" {
		name = contentDocument.Title
	} else {
		name = strings.TrimSuffix(dash.ReplaceAllString(baseName, " "), ".md")
	}
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
