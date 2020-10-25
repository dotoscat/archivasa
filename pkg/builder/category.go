package builder

type Documents []*Document
type Category map[string]Documents

func newCategory() Category {
	categories := make(Category)
	return categories
}

func (c Category) AddDocument(categoryName string, document *Document) {
	currentDocuments, ok := c[categoryName]
	if ok != true {
		var documents Documents
		currentDocuments = documents
	}
	newDocuments := append(currentDocuments, document)
	c[categoryName] = newDocuments
}

func (c Category) Contains(document *Document) bool {
	for _, documents := range c {
		for i := range documents {
			if documents[i] == document {
				return true
			}
		}
	}
	return false
}
