package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

type fileStats struct {
	byteCount      int64
	wordCount      int64
	lineBreakCount int64
	characterCount int64
}

func GetFileStats(file *os.File) fileStats {
	var wordCount int64
	var bytes int64
	var lineBreakCount int64
	var charsCount int64

	reader := bufio.NewReader(file)
	inWord := false

	for {
		c, sz, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				if inWord {
					wordCount++
				}
				break
			} else {
				log.Fatal(err)
			}
		}

		bytes += int64(sz)
		charsCount++

		if unicode.IsSpace(c) {
			if inWord {
				wordCount++
			}
			if c == '\n' {
				lineBreakCount++
			}
			inWord = false
		} else {
			inWord = true
		}
	}
	return fileStats{byteCount: bytes, lineBreakCount: lineBreakCount, wordCount: wordCount, characterCount: charsCount}
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

	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	var file *os.File
	if (stat.Mode() & os.ModeCharDevice) == 0 { // file through pipe operator
		file = os.Stdin
	} else { // file through cmd line
		file, err = os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	stats := GetFileStats(file)

	output := ""

	if wordCount {
		output += strconv.FormatInt(stats.wordCount, 10) + " "
	}
	if lineCount {
		output += strconv.FormatInt(stats.lineBreakCount, 10) + " "
	}
	if byteCount {
		output += strconv.FormatInt(stats.byteCount, 10) + " "
	}
	if characterCount {
		output += strconv.FormatInt(stats.characterCount, 10) + " "
	}
	output += filename
	fmt.Println(output)
}
