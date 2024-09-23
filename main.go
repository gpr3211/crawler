package main

import (
	"fmt"
	"os"

	"sync"

	//	"github.com/gpr3211/crawler/internal/cool"
	"github.com/gpr3211/crawler/spinner"
	cool "github.com/gpr3211/log-cool"
	//"github.com/charmbracelet/lipgloss/list"
)

type Config struct {
	pages       map[string]int
	bodies      map[string]string
	baseURL     *string
	mu          *sync.Mutex
	concControl chan struct{}
	wg          *sync.WaitGroup
}

// htmlmap contains the resp.bodies of the requests, unique
var htmlmap = make(map[string]string)

func main() {
	//conf := &cool.Coolfig{LogFileName: "app.log", LogDir: "/tmp/"}
	conf := &cool.Coolfig{
		LogFileName: "crawler.log",
		LogDir:      os.TempDir(),
	}
	err := cool.Initialize(conf)
	if err != nil {
		// handle error
	}

	args := os.Args
	if len(args) < 2 && len(args) != 1 {
		fmt.Println("no website provided")
		cool.Println("FATAL: not enough args")
		os.Exit(1)
	} else if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	inp := args[1]

	cool.Info("Starting crawl... target >> " + inp)
	cool.Warn(" test Warning blah")
	cool.Fatal(" Fatal error cool")

	// tui element
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
	mu := sync.Mutex{}
	var cc chan struct{}

	_ = &Config{
		pages:       pages,
		bodies:      htmlmap,
		baseURL:     &inp,
		mu:          &mu,
		concControl: cc,
	}

	// start recursive crawl of the entire base link
	//	cfg.crawlPage(inp)

	printMap(pages)
	cool.Info("Printing report for " + inp)

	return

}

func GetMapBody(m map[string]string, s string) string {
	if m[s] == "" {
		return ""
	}
	return m[s]

}
func printMap(m map[string]int) {
	s := spinner.New()
	s.Run()

	fmt.Println("Map contents:")
	if len(m) == 0 {
		fmt.Println("  (empty map)")
		cool.Warn("Empty map")
		return
	}
	/*	l := list.New()

		for key, value := range m {
			l.Items(key, value)
		}
		fmt.Println(l)
		fmt.Printf("map size: %v\n", len(m))
	*/
	cool.Info("Printing map")
	for key, value := range m {
		fmt.Printf("Found %d internal links to %s\n", value, key)
	}

}
