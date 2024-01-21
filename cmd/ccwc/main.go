package main

import (
	"flag"
	"fmt"
)

func main() {
	var byteCount bool
	flag.BoolVar(&byteCount, "c", false, "count bytes in the file")
	flag.Parse()
	fmt.Println(byteCount)
}
