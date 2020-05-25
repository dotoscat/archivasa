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
		if file.IsDir() {
			continue
		}
		filePath := filepath.Join(dir, file.Name())
		document := NewDocument(filePath)
		documents = append(documents, document)
	}
	return documents
}
