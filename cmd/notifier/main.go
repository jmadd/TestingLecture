package main

import (
	"flag"
	"fmt"
)

func main() {
	pending := flag.Int("pending", 0, "number of items awaiting review")
	flag.Parse()

	if !hasPending(*pending) {
		fmt.Println("No attachments are waiting for review.")
		return
	}

	fmt.Println("Your file has been shared with the team.")
	fmt.Println("Open the document to continue.")
	fmt.Println("Attach a document to notify the team.")
}

func hasPending(n int) bool {
	return n > 0
}
