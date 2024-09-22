package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	inp := args[1]
	fmt.Println("starting crawl: ", inp)

	head, err := getHTML(inp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Good Head", head)
	b := getURLSfromHTML(head, inp)
	for _, s := range b {
		fmt.Printf("\n----\nLink: %v\n----\n", s)

	}

	return

}
