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

func CreatePostspage(website *Website, nPosts int, prev, next *Postspage) *Postspage {
	return &Postspage{Webpage{website, "", ""}, make([]*Document, nPosts), 0, prev, next}
}

func (postspage *Postspage) Init(website *Website, nDocuments int, prev, next *Postspage) {
	*postspage = Postspage{Webpage{website, "", ""}, make([]*Document, nDocuments), 0, prev, next}
}

func (postspage *Postspage) AddPost(post *Document) bool {
	fmt.Println("AddPost -> ", postspage.iCurrentPost, " : ", len(postspage.Posts))
	if postspage.iCurrentPost >= len(postspage.Posts) {
		return false
	}
	postspage.Posts[postspage.iCurrentPost] = post
	postspage.iCurrentPost++
	fmt.Println("posts: ", postspage.Posts)
	return true
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
		if i == 0 {
			pages[i] = CreatePostspage(website, documentsPerPage, nil, pages[1])
			page := pages[i]
			// page.Init(website, documentsPerPage, nil, &pages[1])
			page.DistributeDocumets(iDocuments, iDocuments+documentsPerPage, documents)
			iDocuments += documentsPerPage
		} else if i == len(pages)-1 {
			lefts := documentsLeft
			if documentsLeft == 0 {
				lefts = documentsPerPage
			}
			pages[i] = CreatePostspage(website, lefts, pages[i-1], nil)
			//page.Init(website, lefts, &pages[i-1], nil)
			page := pages[i]
			page.DistributeDocumets(iDocuments, iDocuments+lefts, documents)
			iDocuments += lefts
		} else {
			pages[i] = CreatePostspage(website, documentsPerPage, pages[i-1], pages[i+1])
			// page.Init(website, documentsPerPage, &pages[i-1], &pages[i+1])
			page := pages[i]
			page.DistributeDocumets(iDocuments, iDocuments+documentsPerPage, documents)
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
	return pages
}
