package main

import (
	"flag"
	"fmt"
)

func main() {
	count := flag.Int("count", 0, "number of exported items")
	flag.Parse()

	if *count == 0 {
		fmt.Println("No documents found in this folder.")
		return
	}

	fmt.Println("Your file is ready to download.")
	fmt.Println("The document was archived successfully.")
}
