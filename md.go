package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Document represents a markdown document
type Document struct {
	Title string
}

var dir = flag.String("dir", "", "Markdown files directory")

func main() {
	flag.Parse()

	globPath := filepath.Join(*dir, "*.md")

	files, _ := filepath.Glob(globPath)
	for _, fpath := range files {
		log.Println(fpath)

		buffer, err := ioutil.ReadFile(fpath)
		if err != nil {
			log.Fatal(err)
		}

		md := &Markdown{Buffer: string(buffer)}
		md.Init()

		if err := md.Parse(); err != nil {
			log.Fatal(err)
		}
		log.Println(md.Title)
	}
}
