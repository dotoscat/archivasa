package context

// Context allows any struct that represents
// a document to give its url in any way
type Context interface {
	URL() string
	OutputPath() string
}
