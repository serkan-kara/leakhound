package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: leakhound <file or path>")
		os.Exit(2) // usage error
	}

	path := os.Args[1]
	results := ScanPath(path)

	for _, f := range results {
		fmt.Printf("%s:%d [%s] %s\n", f.File, f.Line, f.FindingType, MaskWithStrategy(f.Match, f.Mask))
	}

	if len(results) > 0 {
		os.Exit(1)
	}
}
