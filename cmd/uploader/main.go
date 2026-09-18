package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	count := flag.Int("count", 0, "number of uploaded items")
	purge := flag.Bool("purge", false, "delete the project after reporting")
	flag.Parse()

	if *count == 0 {
		fmt.Println("No attachments uploaded yet. Drop an attachment here to get started.")
		os.Exit(0)
	}

	fmt.Printf("%d file ready\n", *count)
	fmt.Println("The file was uploaded successfully.")
	fmt.Println("Rename this document before sharing it.")

	if *purge {
		fmt.Println("Deleting a project removes all its content. This cannot be undone.")
	}
}
