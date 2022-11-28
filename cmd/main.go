package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RossoDiablo/html_link_parser/link"
)

var htmlFile = flag.String("filename", "ex1.html", "HTML file for parsing")

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *htmlFile == "" {
		exit("Error parsing filename")
	}
	f, err := os.Open(*htmlFile)
	if err != nil {
		exit("Error opening file")
	}
	defer f.Close()

	links, err := link.Parse(f)
	if err != nil {
		exit("Error parsing html doc!")
	}
	link.ShowLinks(links)
}
