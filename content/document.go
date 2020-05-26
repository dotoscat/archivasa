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
	Date time.Time
	Tags []string
	Path string
}

func (d *Document) String() string {
	output := fmt.Sprintf("\n%v\n===\n%v\n%v\n%v\n", d.Path, d.Date, d.Tags)
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
