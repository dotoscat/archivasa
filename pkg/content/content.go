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
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
)

type Content struct {
	Pages Documents
	Posts Documents
}

func Read(cwd string) Content {
	pagesPath := filepath.Join(cwd, "content", "pages")
	postsPath := filepath.Join(cwd, "content", "posts")
	posts := readDir(postsPath)
	sort.Sort(posts)
	return Content{Pages: readDir(pagesPath), Posts: posts}
}

func readDir(dir string) Documents {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln("Error reading ", dir, " ; ", err)
	}
	documents := make(Documents, 0)
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".md" {
			continue
		}
		filePath := filepath.Join(dir, file.Name())
		document := NewDocument(filePath)
		documents = append(documents, document)
	}
	return documents
}
