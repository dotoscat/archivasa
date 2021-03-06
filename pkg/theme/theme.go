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

// Theme is a collection of templates and other resources from the cwd
type Theme struct {
    templates map[string]*template.Template
    folder    string
}

// New create a new theme from the folder "theme"
func Load(cwd string) *Theme {
    themePath := filepath.Join(cwd, "theme")
    templatePath := filepath.Join(themePath, "templates")
    basicTemplatePath := filepath.Join(templatePath, "basic.tmpl")
    documentTemplatePath := filepath.Join(templatePath, "document.tmpl")
    postspageTemplatePath := filepath.Join(templatePath, "postspage.tmpl")

    templates := map[string]*template.Template{
        "document":  template.Must(template.ParseFiles(basicTemplatePath, documentTemplatePath)),
        "postspage": template.Must(template.ParseFiles(basicTemplatePath, postspageTemplatePath)),
    }

    return &Theme{templates, themePath}
}

func (t *Theme) Templates(name string) *template.Template {
    documentTemplate, ok := t.templates[name]
    if !ok {
        log.Fatalln(name, " does not exists.")
    }
    return documentTemplate
}

// Copy copy the resources from the source to the output folder
func (t *Theme) Copy(outputFolder string) {
    names := []string{
        "css",
        "js",
        "images"}
    for _, name := range names {
        folder := filepath.Join(t.folder, name)
        joinOutputFolder := filepath.Join(outputFolder, name)
        fmt.Println(folder, joinOutputFolder)
        copyFolder(folder, joinOutputFolder)
    }
}

func copyFolder(src, dst string) error {
    files, err := ioutil.ReadDir(src)
    if err != nil {
        log.Println(err)
        return err
    }
    if _, err := os.Stat(dst); os.IsNotExist(err) {
        os.Mkdir(dst, os.ModePerm)
    }
    for _, file := range files {
        fmt.Println("copy file", file.Name())
        if file.IsDir() {
            folderSrc := filepath.Join(src, file.Name())
            folderDst := filepath.Join(dst, file.Name())
            copyFolder(folderSrc, folderDst);
            continue
        }
        fileSrc := filepath.Join(src, file.Name())
        fileDst := filepath.Join(dst, file.Name())
        fmt.Println("copy:", fileSrc, fileDst)
        copyFile(fileSrc, fileDst)
    }
    return nil
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
