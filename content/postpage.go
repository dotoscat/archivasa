package content

import "fmt"

// Postpages are generated from Webpages (or documents)
type Postspage struct {
	Webpage
	Posts        []*Document
	iCurrentPost int
	Prev         *Postspage
	Next         *Postspage
}

func CreatePostspage(website *Website, nPosts int) *Postspage {
	return &Postspage{Webpage: Webpage{website, "", ""}, Posts: make([]*Document, nPosts)}
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

func (postspage *Postspage) DistributeDocumets(start, end int, documents []*Document) {
	chunk := documents[start:end]
	for _, document := range chunk {
		postspage.AddPost(document)
	}
	fmt.Println("Anyadidos? ", postspage.Posts, ", ", len(postspage.Posts))
}

func CreatePostspages(documents []*Document, documentsPerPage int, outputDir, prefix string, website *Website) []*Postspage {
	numberOfPages := len(documents) / documentsPerPage
	documentsLeft := len(documents) % documentsPerPage
	if documentsLeft != 0 {
		numberOfPages++
	}
	pages := make([]*Postspage, numberOfPages)
	iDocuments := 0
	for i := 0; i < len(pages); i++ {
		pages[i] = CreatePostspage(website, documentsPerPage)
		if i == len(pages)-1 {
			lefts := documentsLeft
			if documentsLeft == 0 {
				lefts = documentsPerPage
			}
			pages[i].DistributeDocumets(iDocuments, iDocuments+lefts, documents)
			iDocuments += lefts
		} else {
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
