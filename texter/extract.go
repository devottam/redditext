package texter

import (
	"../network"
	"bytes"
	"code.google.com/p/go.net/html"
	"strings"
)

// Document is a text holder.
// It holds the text content of the URL.
// TODO: It should contain different types of readable text
// for example html text, cli text, etc.
type Document struct {
	URL  string
	Body string
	Size int
}

// Extracts the text after fetching the URL.
// TODO: This method should be more elaborate
// to extract text for html, cli, etc.
// ** NO CHECK FOR ERRS: WHY? **
func (d *Document) ExtractText() {
	if d.URL != "" {
		b, _ := network.ContentFromURL(&d.URL)
		txt, _ := textFromHTML(&b)
		d.Body = *txt
		d.Size = len(d.Body)
	}
}

func NewTexter(url string) *Document {
	d := new(Document)
	d.URL = url
	d.ExtractText()
	return d
}

func textFromHTML(b *[]byte) (*string, error) {
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
			// FIXME: For now it is the only thing
			// that determines if the text is for reading.
			if len(str) > 10 {
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
