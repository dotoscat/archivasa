package content

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// PathsFrom Returns a list of path from a path
func pathsFrom(directoryPath string) ([]string, error) {
	directory, err := os.Open(directoryPath)
	defer directory.Close()
	if err != nil {
		log.Fatalln(err)
	}
	dirnames, err := directory.Readdirnames(0)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("info dirname", dirnames, len(dirnames))
	paths := make([]string, len(dirnames))
	for i, dirname := range dirnames {
		path := filepath.Join(directoryPath, dirname)
		fmt.Println("path from dirname", path, len(path))
		paths[i] = path
	}
	fmt.Println("paths len", len(paths), paths)
	return paths, err
}

// PathsFromPages Returns a list from the folder 'pages' inside 'content'
func PathsFromPages(cwd string) ([]string, error) {
	path := filepath.Join(cwd, "pages")
	println("path", path)
	return pathsFrom(path)
}

func PathsFromPosts(cwd string) ([]string, error) {
	path := filepath.Join(cwd, "posts")
	println("path", path)
	return pathsFrom(path)
}
