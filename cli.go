package main

import (
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
			t := texter.NewTexter(srs.Items[i].URL)
			fmt.Println(t)

		case "n":
			continue

		default:
			fmt.Println("Aborting...")
			os.Exit(1)
		}
	}
}
