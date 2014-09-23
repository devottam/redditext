package texter

import (
	"../network"
	"bytes"
	"code.google.com/p/go.net/html"
	"regexp"
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
		txt := textFromHTML(&b)
		d.Body = *txt
		d.Size = len(d.Body)
	}
}

// Allocate and return new Document
// that holds the text from the URL
func NewTexter(url string) *Document {
	d := new(Document)
	d.URL = url
	d.ExtractText()
	return d
}


func textFromHTML(b *[]byte) *string {
	var text string
	var z *html.Tokenizer

	z = html.NewTokenizer(bytes.NewReader(*b))
	text = textUsingTokenizer(z)
	return &text
}

// Extract text using tokenizer from lib `go.net/html`.
// More extraction logic inside.
// TODO: have more description here.
func textUsingTokenizer(z *html.Tokenizer) string {
	var txt, clTag string
	var tags []string
	nl := false
	buffer := new(bytes.Buffer)
	defer buffer.Reset()

	for {
		tt := z.Next()
		tag, _ := z.TagName()
		txt = string(z.Text())

		switch tt {
		case html.SelfClosingTagToken, html.CommentToken, html.DoctypeToken:
			if isNewLineTag(&tag) {
				nl = true
			}
			continue

		case html.StartTagToken:
			if string(tag) == "script" {
				tt = z.Next()
				continue
			}
			tags = append(tags, string(tag))

		case html.EndTagToken:
			li := len(tags) - 1
			if li > 0 {
				clTag = tags[li]
				tags = tags[:li]
			}
		}

		if tt == html.EndTagToken && isNewLineTag(&tag) && string(tag) == clTag {
			nl = true
		}

		if tt == html.TextToken {
			writeStringToBuffer(buffer, &txt, &nl)
		} else if tt == html.ErrorToken {
			buffer.WriteString("\n\n-oo-\n\n")
			return buffer.String()
		}
	}
	return buffer.String()
}

// Writes string from the current node/token to the buffer passed.
// `nl` testifies if newline needs to be written too.
// TODO: Probably would be better if I can use channels for this
func writeStringToBuffer(b *bytes.Buffer, t *string, nl *bool) {
	*t = strings.TrimSpace(*t)
	matched, err := regexp.Match("[a-zA-Z0-9]+", []byte(*t))
	if matched && err == nil {
		if *nl {
			b.WriteString("\n\n")
			*nl = false
		}
		_, err := b.WriteString(*t)
		if err != nil {
			panic("text extraction failed.")
		}
	}
}

// Determines if the `tag` needs a newline.
func isNewLineTag(t *[]byte) (nl bool) {
	nl, err := regexp.Match("^(br|li|h1|div|p)$", *t)
	if err != nil {
		nl = false
	}
	return
}
