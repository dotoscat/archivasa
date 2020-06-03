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
	"fmt"
	"os"
	"path/filepath"

	"github.com/dotoscat/archivasa/config"
	"github.com/dotoscat/archivasa/content"
	"github.com/dotoscat/archivasa/theme"
)

func Run(config config.Config, content *content.Content, theme *theme.Theme) {
	outputPath := filepath.Join(config.Cwd, "output")
	pagesPath := filepath.Join(config.Cwd, "output", "pages")
	postsPath := filepath.Join(config.Cwd, "output", "posts")
	MakeOutputdirIfNotExists(outputPath)
	MakeOutputdirIfNotExists(pagesPath)
	MakeOutputdirIfNotExists(postsPath)
	fmt.Println("nPostspages", len(content.Posts), config.PostsPerPage)
	nPostspages := len(content.Posts) / config.PostsPerPage
	if len(content.Posts)%config.PostsPerPage != 0 {
		nPostspages++
	}
	fmt.Println("nPostspages(2)", nPostspages)
	website := NewWebsite(config.Title, len(content.Pages), nPostspages)
	// Build pages
	for i := 0; i < len(website.Pages); i++ {
		website.Pages[i] = NewDocument(website, content.Pages[i], "pages")
	}
	fmt.Println(website.Pages)
	posts := make([]*Document, len(content.Posts))
	for i := 0; i < len(posts); i++ {
		posts[i] = NewDocument(website, content.Posts[i], "posts")
	}
	fmt.Println(posts)
	FillPostspages(website, posts, config.PostsPerPage)
	documentTemplate := theme.Templates("document")
	postspageTemplate := theme.Templates("postspage")
	for _, page := range website.Pages {
		page.Render(documentTemplate, outputPath)
	}
	for _, postsPage := range website.Postspages {
		postsPage.Render(postspageTemplate, outputPath)
	}
	for _, post := range posts {
		post.Render(documentTemplate, outputPath)
	}
	theme.Copy(outputPath)
}

func MakeOutputdirIfNotExists(outputDirectory string) {
	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.MkdirAll(outputDirectory, os.ModePerm)
	}
}
