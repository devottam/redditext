package redditext

import (
  "os"
  "fmt"
  "log"
)

var greet string = `
Welcome to Redditext
--------------------

`

var takeInput string = "Read text from links submitted to 'subreddit': "

func Start() {
  var subreddit string

  fmt.Println(greet)

  if len(os.Args) > 1 {
    subreddit = os.Args[1]
  } else {
    fmt.Print(takeInput)
    fmt.Scanf("%s", &subreddit)
  }

  items, err := FetchSubReddit(subreddit)
  if err != nil {
    log.Fatal(err)
  }

  externalLinks := make([]string, len(items))
  for i, item := range items {
    externalLinks[i] = item.URL
  }

  var read func(*string)
  read = func (link *string) {
    content, err := PostedLinkContent(*link)
    if err != nil {
      log.Fatal(err)
    }
    fmt.Println(content)
  }

  loop:
  for i := 0; i < len(externalLinks); i++ {
    fmt.Printf("Getting content for %s\n", externalLinks[i])
    fmt.Print("Read article (y|n|a): ")
    var input string
    fmt.Scanf("%s", &input)
    switch input {
    case "y":
      read(&externalLinks[i])

    case "n":
      continue

    case "a":
      break loop

    default:
      fmt.Println("No input recorded. Aborting...")
      os.Exit(1)
    }
  }
}
