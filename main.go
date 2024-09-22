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
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	inp := args[1]
	fmt.Println("starting crawl: ", inp)
	fmt.Println("REPORT for", inp)
	/*
		head, err := getHTML(inp)
		if err != nil {
			fmt.Println(err)
		}
	*/
	// init page map
	var pages = make(map[string]int)

	crawlPage(inp, inp, pages)

	printMap(pages)

	return

}
func printMap(m map[string]int) {
	fmt.Println("Map contents:")
	if len(m) == 0 {
		fmt.Println("  (empty map)")
		return
	}
	for key, value := range m {
		fmt.Printf("  %s: %d\n", key, value)
	}
}
