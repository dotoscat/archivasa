package content

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
	if postspage.iCurrentPost >= len(postspage.Posts) {
		return false
	}
	postspage.Posts[postspage.iCurrentPost] = post
	postspage.iCurrentPost++
	return true
}

func CreatePostspages(documents []*Document, documentsPerPage int, website *Website) []Postspage {
	numberOfPages := len(documents) / documentsPerPage
	documentsLeft := len(documents) % documentsPerPage
	if documentsLeft != 0 {
		numberOfPages++
	}
	pages := make([]Postspage, numberOfPages)
	for i, page := range pages {
		if i == 0 {
			page.Init(website, documentsPerPage, nil, &pages[1])
		} else if i == len(pages)-1 {
			lefts := documentsLeft
			if documentsLeft == 0 {
				lefts = documentsPerPage
			}
			page.Init(website, lefts, &pages[i-1], nil)
		} else {
			page.Init(website, documentsPerPage, &pages[i-1], &pages[i+1])
		}
	}
	return pages
}

// func distributeDocumets
