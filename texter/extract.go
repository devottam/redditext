package texter

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"strings"
)

const MinLengthOfText = 20

// TODO: This is just a dummy function now.
// This should be replaced/converted to more generic type
// Use `interface` to handle text extraction.
// By rule of thumb, with given url, byte, string, etc. it should extract the url
func TextFromHTML(b *[]byte) (*string, error) {
	buffer := new(bytes.Buffer)
	defer buffer.Reset()

	t := buffer.String()
	doc, err := html.Parse(bytes.NewReader(*b))
	if err != nil {
		return &t, err
	}

	var extract func(*html.Node)
	extract = func(node *html.Node) {
		if node.Type == html.TextNode {
			str := strings.TrimSpace(node.Data)
			if len(str) > MinLengthOfText {
				buffer.WriteString(str)
				buffer.WriteString("\n")
			}
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}

	extract(doc)
	t = buffer.String()
	return &t, nil
}
