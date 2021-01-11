package builder

import(
    "testing"
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
)


func TestImage(t *testing.T){
    imageTemplate := []byte(`
# Image

![image test](testdata/image)
    `)
    t.Log("Hello!")
    parser := parser.NewWithExtensions(parser.CommonExtensions)
    html := markdown.ToHTML(imageTemplate, parser, nil)
    t.Log(string(html))
}
