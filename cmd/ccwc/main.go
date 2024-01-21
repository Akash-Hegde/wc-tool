package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var byteCount bool
	flag.BoolVar(&byteCount, "c", false, "count bytes in the file")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("file input missing")
		os.Exit(1)
	}
	file := os.Args[2]
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(data), file)
}
