package main

import (
	"flag"
	"fmt"
)

func main() {
	pending := flag.Int("pending", 0, "number of items awaiting review")
	flag.Parse()

	if *pending == 0 {
		fmt.Println("No attachments are waiting for review.")
		return
	}

	fmt.Println("Your file has been shared with the team.")
	fmt.Println("Open the document to continue.")
	fmt.Println("Attach a document to notify the team.")
}
