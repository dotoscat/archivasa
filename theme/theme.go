package theme

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/dotoscat/archivasa/context"
)

type Theme struct {
	templates   map[string]*template.Template
	cwd, folder string
}

func New(cwd string) *Theme {
	themePath := filepath.Join(cwd, "theme")
	templatePath := filepath.Join(themePath, "templates")
	basicTemplatePath := filepath.Join(templatePath, "basic.tmpl")
	indexTemplatePath := filepath.Join(templatePath, "index.tmpl")
	documentTemplatePath := filepath.Join(templatePath, "document.tmpl")

	folder := filepath.Join(cwd, "theme")

	templates := map[string]*template.Template{
		"index":    template.Must(template.ParseFiles(basicTemplatePath, indexTemplatePath)),
		"document": template.Must(template.ParseFiles(basicTemplatePath, documentTemplatePath)),
	}

	return &Theme{templates, cwd, folder}
}

func (t *Theme) Index() *template.Template {
	return t.templates["index"]
}

func (t *Theme) Document() *template.Template {
	return t.templates["document"]
}

func (t *Theme) Copy(outputFolder string) {
	CSSFolder := filepath.Join(t.folder, "css")
	outputFolderCSSFolder := filepath.Join(outputFolder, "css")
	fmt.Println(CSSFolder, outputFolderCSSFolder)
	copyFolder(CSSFolder, outputFolderCSSFolder)
}

func (t *Theme) Render(templateName string, ctx context.Context) {
	pageOutput, err := os.Create(ctx.OutputPath())
	defer pageOutput.Close()
	if err != nil {
		log.Fatal(err)
	}
	template := t.templates[templateName]
	if template == nil {
		log.Fatalf("%v template is nil", templateName)
	}
	if err := template.Execute(pageOutput, ctx); err != nil {
		log.Fatal(err)
	}
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
