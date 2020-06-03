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
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// Webpage represents a internet document with its URL
// A webpage belongs to a website
type Webpage struct {
	*Website
	URL string
}

func NewWebpage(site *Website, url string) *Webpage {
	return &Webpage{site, url}
}

func (w *Webpage) Render(template *template.Template, outputDirectory string) {
	outputPath := filepath.Join(outputDirectory, w.URL)
	file, err := os.Create(outputPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if template == nil {
		log.Fatalf("%v template is nil", template)
	}
	if err := template.Execute(file, w); err != nil {
		log.Fatal(err)
	}
}
