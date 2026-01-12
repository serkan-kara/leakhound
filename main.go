package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: leakhound <file>")
		return
	}

	file := os.Args[1]
	results := ScanFile(file)

	for _, r := range results {
		fmt.Println(r)
	}
}
