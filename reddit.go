package main

import (
  "bytes"
  "code.google.com/p/go.net/html"
  "encoding/json"
  "fmt"
  "strings"
)

// TODO: subreddit should be passed as options like { limit: 10, name: 'golang' }
func FetchSubReddit(subRedditSting string) ([]Item, error) {
  return fetchReddit(fmt.Sprintf("http://reddit.com/r/%s.json", subRedditSting))
}

func FetchRedditComments(urlString string) ([]Item, error) {
  return fetchReddit(urlString)
}

func PostedLinkContent(link string) (string, error) {
  body, err := Fetch(link)
  if err != nil {
    return "", err
  }

  buffer := new(bytes.Buffer)
  defer buffer.Reset()

  doc, err := html.Parse(bytes.NewReader(body))
  if err != nil {
    return "", err
  }

  var extract func(*html.Node)
  extract = func(node *html.Node) {
    if node.Type == html.TextNode {
      str := strings.TrimSpace(node.Data)
      if len(str) > minTextLength {
        buffer.WriteString(str)
        buffer.WriteString("\n")
      }
    }
    for c := node.FirstChild; c != nil; c = c.NextSibling {
      extract(c)
    }
  }

  extract(doc)
  return buffer.String(), nil
}

/*
  ------------------
  Private to package
  ------------------
*/
func fetchReddit(urlString string) ([]Item, error) {
  body, err := Fetch(urlString)
  if err != nil {
    return nil, err
  }
  content := new(Response)
  err = json.NewDecoder(bytes.NewReader(body)).Decode(content)
  if err != nil {
    return nil, err
  }
  items := make([]Item, len(content.Data.Children))
  for i, child := range content.Data.Children {
    items[i] = child.Data
    items[i].RedditLink = fmt.Sprintf("http://reddit.com%s", child.Data.RedditLink)
  }
  return items, nil
}
