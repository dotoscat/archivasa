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
