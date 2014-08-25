package reddit

import (
	"../../network"
	"encoding/json"
	"fmt"
)

// The `reddit` URL
const SubredditJSONURL = "https://reddit.com/r/%s.json"
const RedditPermaLink = "https://reddit.com%s"

// Main content of the data from the response from
// fetching the `reddit URL`
type Item struct {
	RedditId   string `json:"id"`
	Title      string
	SelfText   string
	Domain     string
	URL        string
	RedditLink string `json:"permalink"`
	Author     string
	Likes      int
	UpVotes    int `json:"ups"`
	DownVotes  int `json:"downs"`
	Score      int
	Comments   int     `json:"num_comments"`
	CreatedAt  float32 `json:"created"`
	NSFW       bool    `json:"over_18"`
}

// String implementation of the Item
func (i *Item) String() string {
	var comments string
	switch i.Comments {
	case 0:
		comments = "No Comments"
	case 1:
		comments = "1 Comment"
	default:
		comments = fmt.Sprintf("%d Comments", i.Comments)
	}
	return fmt.Sprintf("\x1b[1;34m%s \x1b[1;33m(%s)\x1b[0m\n\x1b[1;31mScore: %d\x1b[1;37m | \x1b[1;32m%s\x1b[0m\n", i.Title, i.Domain, i.Score, comments)
}

// Holds the subreddit data
type SubredditData struct {
	Subreddit string
	Items     []*Item
}

// While unmarshaling the JSON data
// Populate the subreddit data.
func (srd *SubredditData) UnmarshalJSON(b []byte) error {
	res := new(response)
	err := json.Unmarshal(b, res)
	if err != nil {
		return err
	}

	srd.Items = make([]*Item, len(res.Data.Children))
	for i := range res.Data.Children {
		srd.Items[i] = &res.Data.Children[i].Data
		srd.Items[i].RedditLink = fmt.Sprintf(RedditPermaLink, srd.Items[i].RedditLink)
	}
	return nil
}

// The `json response` from the URL
type response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

// Main function to return the latest items of given subreddit
func SRItems(sr *string) (*SubredditData, error) {
	var url = fmt.Sprintf(SubredditJSONURL, *sr)
	resData, err := network.ContentFromURL(&url)
	if err != nil {
		return nil, fmt.Errorf("Error fetching subreddit '%s'. Please try again later", *sr)
	}
	sd := new(SubredditData)
	sd.Subreddit = *sr

	err = json.Unmarshal(resData, sd)
	if err != nil {
		return nil, fmt.Errorf("Response couldn't be read. Error: %s", err)
	}

	return sd, nil
}
