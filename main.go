package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough args")
	}

	inputPath := os.Args[1]

	content, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	output := markdown.ToHTML(content, nil, nil)

	outputPath := filepath.Join(".", "out")
	err = os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filepath.Join(outputPath, "output.html"), output, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
