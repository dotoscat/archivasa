package content

// Postpages are generated from Webpages (or documents)
type Postpage struct {
	Posts []*Document
	Url   string
}

func NewPostpage(nPosts int) *Postpage {
	return &Postpage{make([]*Document, nPosts), ""}
}

func (postpage *Postpage) AddPost(post *Document) {
	postpage.Posts = append(postpage.Posts, post)
}
