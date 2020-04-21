package theme

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type Theme struct {
	index, document *template.Template
	cwd, folder     string
}

func New(cwd string) *Theme {
	themePath := filepath.Join(cwd, "theme")
	templatePath := filepath.Join(themePath, "templates")
	basicTemplatePath := filepath.Join(templatePath, "basic.tmpl")
	indexTemplatePath := filepath.Join(templatePath, "index.tmpl")
	documentTemplatePath := filepath.Join(templatePath, "document.tmpl")

	index := template.Must(template.ParseFiles(basicTemplatePath, indexTemplatePath))
	document := template.Must(template.ParseFiles(basicTemplatePath, documentTemplatePath))

	folder := filepath.Join(cwd, "theme")

	return &Theme{index, document, cwd, folder}
}

func (t *Theme) Index() *template.Template {
	return t.index
}

func (t *Theme) Document() *template.Template {
	return t.document
}

func (t *Theme) Copy(outputFolder string) {
	CSSFolder := filepath.Join(t.folder, "css")
	outputFolderCSSFolder := filepath.Join(outputFolder, "css")
	fmt.Println(CSSFolder, outputFolderCSSFolder)
	copyFolder(CSSFolder, outputFolderCSSFolder)
}

func copyFolder(src, dst string) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		os.Mkdir(dst, os.ModePerm)
	}
	for _, file := range files {
		fmt.Println("copy file", file.Name())
		if file.IsDir() {
			continue
		}
		fileSrc := filepath.Join(src, file.Name())
		fileDst := filepath.Join(dst, file.Name())
		fmt.Println("copy:", fileSrc, fileDst)
		copyFile(fileSrc, fileDst)
	}
}

func copyFile(src, dst string) {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatalln(err)
	}
	dstFile, err := os.Create(dst)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := io.Copy(dstFile, srcFile); err != nil {
		log.Fatalln(err)
	}
}
