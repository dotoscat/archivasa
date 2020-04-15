package content

type Postpage struct {
	Posts []*Document
}

func NewPostpage(nPosts int) *Postpage {
	return &Postpage{make([]*Document, nPosts)}
}

func (postpage *Postpage) AddPost(post *Document) {
	postpage.Posts = append(postpage.Posts, post)
}
