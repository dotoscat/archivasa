package builder

import (
	"fmt"
	"testing"
)

func TestWebpage(t *testing.T) {
	site := NewWebsite("test", "")
	document := NewDocument(site, "", "")
	fmt.Println(document)
}
