package main

import (
	"flag"
	"log"

	"github.com/atridadl/daisygen/generator"
)

const version = "0.1.5"

func main() {
	extensions := flag.String("extensions", "", "Comma-separated list of file extensions")
	directory := flag.String("directory", ".", "Directory to search for files")
	outputDir := flag.String("output-dir", "../../public/css", "Output directory for generated CSS")
	versionFlag := flag.Bool("version", false, "Print version")
	flag.Parse()

	// Print version and exit if version flag is set
	if *versionFlag {
		log.Printf("daisygen version %s", version)
		return
	}

	if *extensions == "" {
		log.Fatal("No extensions provided")
	}

	err := generator.Generate(*extensions, *directory, *outputDir)
	if err != nil {
		log.Fatal(err)
	}
}
