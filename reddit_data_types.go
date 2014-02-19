package main

import (
  "fmt"
)

const (
  minTextLength int = 10
)

type (
  Response struct {
    Data struct {
      Token    string `json:"modash"`
      Children []struct {
        Kind string
        Data Item
      }
      After  string
      Before string
    }
  }

  // FIXME: not sure how nested structs in go works
  Replies struct {
    Data struct {
      Children []struct {
        Kind string
        Data Item
      }
    }
  }

  Item struct {
    RedditId   string `json:"id"`
    Title      string
    SelfText   string
    Domain     string
    URL        string
    RedditLink string `json:"permalink"`
    Author     string
    Likes      int
    UPVotes    int `json:"downs"`
    DownVotes  int `json:"ups"`
    Score      int
    Comments   int     `json:"num_comments"`
    CreatedAt  float32 `json:"created"`
  }
)

func (i *Item) String() string {
  var comments string
  switch i.Comments {
  case 0:
    comments = " (No Comments)"
  case 1:
    comments = " (1 Comment)"
  default:
    comments = fmt.Sprintf(" (%d Comments)", i.Comments)
  }
  return fmt.Sprintf("%s%s\n%s\n%s\n", i.Title, comments, i.URL, i.RedditLink)
}
