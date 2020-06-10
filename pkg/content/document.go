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

package content

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Document struct {
	Date  time.Time
	Tags  []string
	Title string
	Path  string
}

func (d *Document) String() string {
	output := fmt.Sprintf("\n%v\n===\n%v\n%v\n", d.Path, d.Date, d.Tags)
	return output
}

func NewDocument(path string) *Document {
	document := Document{}
	document.Path = path
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Println("Error reading ", path, " ; ", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			break
		}
		keyValue := strings.Split(line, "=")
		if len(keyValue) != 2 {
			continue
		}
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		switch key {
		case "date":
			document.Date = getDate(value)
		case "tags":
			document.Tags = getTags(value)
		case "title":
			document.Title = value
		}
	}
	return &document
}

func getTags(line string) (tags []string) {
	chunks := strings.Split(line, ",")
	for _, chunk := range chunks {
		tags = append(tags, strings.TrimSpace(chunk))
	}
	return
}

func getDate(date string) time.Time {
	var year int
	var month time.Month
	var day int
	fmt.Sscanf(date, "%d-%d-%d", &year, &month, &day)
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

type Documents []*Document

func (d Documents) Len() int {
	return len(d)
}

func (d Documents) Less(i, j int) bool {
	return d[i].Date.After(d[j].Date)
}

func (d Documents) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}
