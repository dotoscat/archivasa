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
)

// Postpages are generated from Webpages (or documents)
type Postspage struct {
	Webpage
	Posts        []*Document
	iCurrentPost int
	Prev         *Postspage
	Next         *Postspage
}

func FillPostspages(website *Website, posts []*Document, postsPerPage int) {
	postsLeft := len(posts) % postsPerPage
	if len(website.Postspages) == 1 {
		website.Postspages[0] = CreatePostspage(website, posts, "/index.html")
		return
	}
	iPost := 0
	fmt.Println("fill, posts per page", postsPerPage)
	fmt.Println("posts left", postsLeft)
	for i := 0; i < len(website.Postspages) && len(website.Postspages) > 1; i++ {
		var URL string
		if i == 0 {
			URL = "/index.html"
		} else {
			URL = fmt.Sprintf("/page%v.html", i)
		}
		var postsChunk []*Document
		fmt.Println("i: ", i, len(website.Postspages)-1)
		if i == len(website.Postspages)-1 && postsLeft != 0 {
			postsChunk = posts[iPost : iPost+postsLeft]
		} else {
			postsChunk = posts[iPost : iPost+postsPerPage]
			iPost += postsPerPage
		}
		website.Postspages[i] = CreatePostspage(website, postsChunk, URL)
	}
	// Link postspages between them
	fmt.Println("Link pages")
	pages := website.Postspages
	for i, page := range pages {
		if i == 0 {
			page.LinkPages(nil, pages[1])
		} else if i == len(pages)-1 {
			page.LinkPages(pages[i-1], nil)
		} else {
			page.LinkPages(pages[i-1], pages[i+1])
		}
	}
	fmt.Println("End link pages")
}

func CreatePostspage(website *Website, posts []*Document, URL string) *Postspage {
	nPosts := len(posts)
	postspage := Postspage{Webpage: Webpage{website, URL}, Posts: make([]*Document, nPosts)}
	for _, post := range posts {
		postspage.AddPost(post)
	}
	return &postspage
}

func (postspage *Postspage) AddPost(post *Document) bool {
	if postspage.iCurrentPost >= len(postspage.Posts) {
		return false
	}
	postspage.Posts[postspage.iCurrentPost] = post
	postspage.iCurrentPost++
	return true
}

func (postspage *Postspage) LinkPages(prev, next *Postspage) {
	postspage.Prev = prev
	postspage.Next = next
}

func (postspage *Postspage) String() string {
	return fmt.Sprintf("page: %v\nNext: %v\nPrev: %v\n", postspage.URL, postspage.Next, postspage.Prev)
}

/*

func (postspage *Postspage) DistributeDocuments(start, end int, documents []*Document) {
	chunk := documents[start:end]
	for _, document := range chunk {
		postspage.AddPost(document)
	}
	fmt.Println("Anyadidos? ", postspage.Posts, ", ", len(postspage.Posts))
}

func CreatePostspages(documents []*Document, documentsPerPage int, outputDir, prefix string, website *Website) []*Postspage {
	numberOfPages := len(documents) / documentsPerPage
	documentsLeft := len(documents) % documentsPerPage
	fmt.Println("documents left", documentsLeft, "number of pages", numberOfPages)
	if documentsLeft != 0 {
		numberOfPages++
	}
	fmt.Println("(2) documents left", documentsLeft, "number of pages", numberOfPages)
	pages := make([]*Postspage, numberOfPages)
	iDocuments := 0
	for i := 0; i < len(pages) && len(pages) > 1; i++ {
		if i == len(pages)-1 {
			lefts := documentsLeft
			if documentsLeft == 0 {
				lefts = documentsPerPage
			}
			pages[i] = CreatePostspage(website, lefts)
			pages[i].DistributeDocumets(iDocuments, iDocuments+lefts, documents)
			iDocuments += lefts
		} else {
			pages[i] = CreatePostspage(website, documentsPerPage)
			pages[i].DistributeDocumets(iDocuments, iDocuments+documentsPerPage, documents)
			iDocuments += documentsPerPage
		}
		if i == 0 {
			pages[i].BuildURL("", "/index.html")
			pages[i].BuildOutputPath(outputDir, "/index.html")
		} else {
			URLName := fmt.Sprint("page", i, ".html")
			pages[i].BuildURL(prefix, URLName)
			pages[i].BuildOutputPath(outputDir, URLName)
		}
		fmt.Println("Anyadidos? (again)", pages[i].Posts, ", ", len(pages[i].Posts))
	}
	if len(pages) == 1 {
		pages[0] = CreatePostspage(website, documentsPerPage)
		pages[0].DistributeDocumets(0, documentsPerPage, documents)
		pages[0].BuildURL("", "/index.html")
		pages[0].BuildOutputPath(outputDir, "/index.html")
		return pages
	}
	// Link postspages between them
	for i, page := range pages {
		if i == 0 {
			page.LinkPages(nil, pages[1])
		} else if i == len(pages)-1 {
			page.LinkPages(pages[i-1], nil)
		} else {
			page.LinkPages(pages[i-1], pages[i+1])
		}
	}
	return pages
}
*/
