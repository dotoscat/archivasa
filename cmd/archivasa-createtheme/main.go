package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const mainCSS = `
body {
    font-family: Verdana, Geneva, Tahoma, sans-serif;
    color: #333;
    background-color: azure;
    margin-left: auto;
    margin-right: auto;
    max-width: 60rem;
    line-height: 1.5rem;
    padding: 1rem;
}

header h1, section nav {
    text-align: center;
}

footer {
    font-size: small;
}
`

const basicTemplate = `<!doctype html>
<html>
    <head>
        <meta charset='utf-8'>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>{{ block "title" . }}{{ .Title }}{{ end }}</title>
        <link rel="stylesheet" href="/css/normalize.css">
        <link rel="stylesheet" href="/css/main.css">
        {{ block "styles" . }}
        {{ end }}
        {{ block "scripts" . }}
        {{ end }}
    </head>
    <body>
        <header>
            <h1><a href="/">{{.Title}}</a></h1>
            <nav>
                <ul>
                    {{ range .Pages }}<li><a href="{{ .URL }}">{{ .Name }}</a></li>
                    {{ end }}
                </ul>
            </nav>
        </header>
        {{ block "content" . }}
        {{ end }}
        <footer>
            {{ block "footer" . }}
            {{ end }}
            <small>This site is synthetized by <a href="https://github.com/dotoscat/archivasa">archivasa</a></small>
        </footer>
    </body>
</html>
`

const documentTemplate = `{{ template "basic.tmpl" }}
{{ block "content" . }}
<article>
    {{ .Content }}
</article>
{{ end }}
`

const postspageTemplate = `{{ template "basic.tmpl" }}
{{ block "content" . }}
    <section>
        <ul>
        {{ range .Posts }}
            <li>
                <a href="{{ .URL }}">{{ .Date }} - {{ .Name }}</a>
            </li>
        {{ end }}
        </ul>
        <nav>
            {{ if .Prev }}
                <a href="{{ .Prev.URL }}">Prev</a>
            {{ else if and (not .Prev) .Next }}
                <span>First</span>
            {{ end }}
             | 
            {{ if .Next}}
                <a href="{{ .Next.URL }}">Next</a>
            {{ else if and .Prev (not .Next) }}
                <span>End</span>
            {{ end }}
        </nav>
    </section>
{{ end }}
`

var structure = [...]string{
	"/theme/templates",
	"/theme/css"}

var files = map[string]string{
	"/theme/css/main.css":             mainCSS,
	"/theme/css/normalize.css":        normalizeCSS,
	"/theme/templates/basic.tmpl":     basicTemplate,
	"/theme/templates/document.tmpl":  documentTemplate,
	"/theme/templates/postspage.tmpl": postspageTemplate}

const version = "0.4.0"

func copyFile(src, dst string) {
	content, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatalln(err)
	}
	if err := ioutil.WriteFile(dst, content, os.ModePerm); err != nil {
		log.Fatalln(err)
	}
}

func buildPaths(dst string, files map[string]string) map[string]string {
	newPaths := make(map[string]string)
	for path, content := range files {
		newPath := filepath.Join(dst, path)
		newPaths[newPath] = content
	}
	return newPaths
}

func mkDirs(files map[string]string) error {
	for path := range files {
		err := os.MkdirAll(path, os.ModeDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func copyFiles(files map[string]string) {

}

func main() {
	knowVersion := flag.Bool("version", false, "-version")
	ask := flag.Bool("noask", false, "No ask to overwrite each file (politely).")
	flag.Parse()
	if *knowVersion {
		fmt.Println(version)
		return
	}
	fmt.Println("ask: ", *ask)
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	for _, path := range structure {
		fmt.Println(path)
		continue
		err := os.MkdirAll(filepath.Join(pwd, path), os.ModeDir)
		if err != nil {
			log.Fatalln(err)
		}
	}
	os.Exit(0)
	for path, content := range files {
		err := ioutil.WriteFile(filepath.Join(pwd, path), []byte(content), os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
