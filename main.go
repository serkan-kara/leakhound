package main

import (
	"fmt"
	"os"
	"runtime"
)

type Options struct {
	Path     string
	Excludes []string
}

func parseArgs(argv []string) (Options, error) {
	var options Options

	i := 1
	for i < len(argv) {
		argument := argv[i]

		switch argument {
		case "--exclude":
			if i+1 >= len(argv) {
				return options, fmt.Errorf("missing value for --exclude")
			}
			options.Excludes = append(options.Excludes, argv[i+1])
			i += 2
			continue
		default:
			// consider it as a path if not flag
			if len(argument) > 0 && argument[0] == '-' {
				return options, fmt.Errorf("Unknown flag: %s", argument)
			}
			if options.Path == "" {
				options.Path = argument
				i++
				continue
			}
			// if second positional provided throw error
			return options, fmt.Errorf("Unexpected argument: %s", argument)
		}
	}

	if options.Path == "" {
		return options, fmt.Errorf("Missing path")
	}
	return options, nil
}

func main() {

	helpWanted := false
	versionWanted := false

	// skip the leakhound name with slice
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h", "--help":
			helpWanted = true
		case "-v", "--version":
			versionWanted = true
		}
	}

	if helpWanted {
		message := `LeakHound - DevSecOps Secret Scanner

Usage:
  leakhound <path> [options]

Options:
  --exclude    Exclude files or directories from scanning (repeatable)
  --version    Show version and build information
  --help       Show this help message
`
		fmt.Print(message)
		os.Exit(0)
	} else if versionWanted {
		ops := runtime.GOOS
		arch := runtime.GOARCH

		message := `LeakHound %s

Commit     : %s
Platform   : %s-%s
Build date : %s
`
		fmt.Printf(message, Version, Commit, ops, arch, BuildDate)
		os.Exit(0)
	}

	options, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println("Usage: leakhound <file or path> [--exclude <name> ...]")
		fmt.Println("Error:", err)
		os.Exit(2)
	}

	results := ScanPath(options.Path, options.Excludes)

	for _, f := range results {
		fmt.Printf("%s:%d [%s] %s\n", f.File, f.Line, f.FindingType, MaskWithStrategy(f.Match, f.Mask))
	}

	if len(results) > 0 {
		os.Exit(1)
	}
}
