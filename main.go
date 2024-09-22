package main

import (
	"crawler/spinner"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss/list"
)

var htmlmap = make(map[string]string)

func main() {

	args := os.Args
	if len(args) < 2 && len(args) != 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	inp := args[1]
	fmt.Println("starting crawl: ", inp)

	s := spinner.New()
	s.Run()
	/*
		head, err := getHTML(inp)
		if err != nil {
			fmt.Println(err)
		}
	*/
	// init page map
	var pages = make(map[string]int)

	crawlPage(inp, inp, pages, htmlmap)

	printMap(pages)

	fmt.Println("REPORT for", inp)
	return

}
func printMap(m map[string]int) {
	s := spinner.New()
	s.Run()

	fmt.Println("Map contents:")
	if len(m) == 0 {
		fmt.Println("  (empty map)")
		return
	}
	l := list.New()

	for key, value := range m {
		l.Items(key, value)
	}
	fmt.Println(l)
	fmt.Printf("map size: %v\n", len(m))

	/*
		for key, value := range m {
			fmt.Printf("  %s: %d\n", key, value)
		}
	*/
}
