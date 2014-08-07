package main

import (
	"flag"
)

var sr *string
var lim *int

const user = "dan"

func init() {
	sr = flag.String("subreddit", "ALL", "Name of the subreddit to browse. Defaults to ALL.")
	lim = flag.Int("limit", 10, "Total articles to fetch. Default is 10. 5 is the minimum number to fetch.")

	flag.Parse()
}

func main() {
	// If less than 5, set limit to 5.
	if *lim < 5 {
		lim = 5
	}

	// Start the process.
	fmt.Println("Starting process ...")
	Start(sr, *lim)
}
