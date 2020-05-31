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

// Website contains data about the site
type Website struct {
	Title      string
	Pages      []*Document
	Postspages []*Postspage
}

func NewWebsite(title string, nPages, nPostspages int) *Website {
	return &Website{title, make([]*Document, nPages), make([]*Postspage, nPostspages)}
}

/*
// Export generates content
func Export(title string, cwd string) {
	site := new(Website)
	outputDirectory := filepath.Join(cwd, "output")
	pagesDirectory := filepath.Join(outputDirectory, "pages")
	postsDirectory := path.Join(outputDirectory, "posts")
	MakeOutputdirIfNotExists(pagesDirectory)
	MakeOutputdirIfNotExists(postsDirectory)

	if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
		os.Mkdir(outputDirectory, os.ModePerm)
	}
	contentFolder := filepath.Join(cwd, "content")
	contentPagesDirectory := filepath.Join(contentFolder, "pages")
	contentPostsDirectory := filepath.Join(contentFolder, "posts")
	fmt.Println("content folder", contentFolder)
	site.Pages = GetDocumentsFromDir(contentPagesDirectory, outputDirectory, "/pages", site)
	posts := GetDocumentsFromDir(contentPostsDirectory, outputDirectory, "/posts", site)
	sort.Sort(posts)
	fmt.Println("posts:")
	fmt.Println(posts)
	site.Postspages = CreatePostspages(posts, 5, outputDirectory, "/", site)
	fmt.Println("postspages", len(site.Postspages))
	site.RenderPostspages("postspage")
	site.RenderPosts("document")
	site.RenderDocuments(site.Pages, "document")
	// theme.Copy(outputDirectory)
}

func (site *Website) String() string {
	return site.Title
}

func (site *Website) RenderDocuments(documents []*Document, templateName string) {
	for _, document := range documents {
		document.BuildContent()
		// site.Theme.Render(templateName, document)
	}
}

func (site *Website) RenderPostspages(templateName string) {
	for _, page := range site.Postspages {
		fmt.Println("Prev | next", page, page.Prev, page.Next)
		// site.Theme.Render(templateName, page)
	}
}

func (site *Website) RenderPosts(templateName string) {
	for _, page := range site.Postspages {
		fmt.Println("Posts of page: ", len(page.Posts))
		for _, webpage := range page.Posts {
			fmt.Println("Render webpage", webpage)
			webpage.BuildContent()
			// site.Theme.Render(templateName, webpage)
		}
	}
}
*/
