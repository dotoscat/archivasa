package builder

import (
	"fmt"
	"testing"
)

func TestWebpage(t *testing.T) {
	site := NewWebsite("test", 1, 1)
	document := Document{Webpage{site, "test.html"}, "test", "hello", "testprefix"}
	fmt.Println(document)
}
