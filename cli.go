package main

import (
	"./network"
	"./sources/reddit"
	"./texter"
	"fmt"
	"os"
)

func Start() {
	srs, err := reddit.SRItems(sr)
	if err != nil {
		fmt.Printf("System Error occurred. Panicking!!!")
		panic(fmt.Sprintf("Err: %v", err))
	}

	for i := range srs.Items {
		var input string
		fmt.Println(srs.Items[i])
		fmt.Println("Read article (y|n|a): ")
		fmt.Scanf("%s", &input)
		switch input {
		case "y":
			b, err := network.ContentFromURL(&srs.Items[i].URL)
			if err != nil {
				panic(err)
			}
			s, err := texter.TextFromHTML(&b)
			if err != nil {
				panic(err)
			}
			fmt.Println(*s)

		case "n":
			continue

		default:
			fmt.Println("Aborting...")
			os.Exit(1)
		}
	}
}
