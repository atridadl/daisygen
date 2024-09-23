package main

import (
	"flag"
	"log"

	"github.com/atridadl/daisygen/internal/generator"
)

func main() {
	extensions := flag.String("extensions", "", "Comma-separated list of file extensions")
	directory := flag.String("directory", ".", "Directory to search for files")
	outputDir := flag.String("output-dir", "../../public/css", "Output directory for generated CSS")
	flag.Parse()

	if *extensions == "" {
		log.Fatal("No extensions provided")
	}

	err := generator.Generate(*extensions, *directory, *outputDir)
	if err != nil {
		log.Fatal(err)
	}
}
