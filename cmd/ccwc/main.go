package main

import (
	"flag"
	"log"
	"os"
)

type fileStats struct {
	byteCount      int
	wordCount      int
	linCount       int
	characterCount int
}

func getFileStats(filename string) *fileStats {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return &fileStats{
		byteCount: len(data),
	}
}

func main() {
	var byteCount, wordCount, lineCount, characterCount bool
	flag.BoolVar(&byteCount, "c", false, "count bytes in the file")
	flag.BoolVar(&wordCount, "w", false, "count words in the file")
	flag.BoolVar(&lineCount, "l", false, "count lines in the file")
	flag.BoolVar(&characterCount, "m", false, "count characters in the file")
	flag.Parse()

	filename := flag.CommandLine.Arg(0)

	if !byteCount && !wordCount && !lineCount && !characterCount {
		byteCount = true
		wordCount = true
		lineCount = true
	}

	fileStats := getFileStats(filename)
	log.Println(fileStats.byteCount)
}
