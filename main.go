package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"
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
				return options, fmt.Errorf("%w: missing value for --exclude", ErrUsage)
			}
			options.Excludes = append(options.Excludes, argv[i+1])
			i += 2
			continue
		default:
			// consider it as a path if not flag
			if len(argument) > 0 && argument[0] == '-' {
				return options, fmt.Errorf("%w: Unknown flag: %s", ErrUsage, argument)
			}
			if options.Path == "" {
				options.Path = argument
				i++
				continue
			}
			// if second positional provided throw error
			return options, fmt.Errorf("%w: Unexpected argument: %s", ErrUsage, argument)
		}
	}

	if options.Path == "" {
		return options, fmt.Errorf("%w: Missing path", ErrUsage)
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
	if errors.Is(err, ErrUsage) {
		fmt.Println("Usage: leakhound <file or path> [--exclude <name> ...]")
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(3)
	}

	results, scanErr := ScanPath(options.Path, options.Excludes)

	if scanErr != nil {
		fmt.Fprintln(os.Stderr, "Error:", scanErr)
	}

	for _, f := range results.Findings {
		fmt.Printf("%s:%d [%s] %s\n", f.File, f.Line, f.FindingType, MaskWithStrategy(f.Match, f.Mask))
	}

	fmt.Printf(
		"\nSummary: files=%d findings=%d duration=%s\n",
		results.FilesScanned,
		len(results.Findings),
		results.Duration.Truncate(time.Millisecond),
	)

	if errors.Is(scanErr, ErrRuntime) {
		os.Exit(3)
	}

	if len(results.Findings) > 0 {
		os.Exit(1)
	}

	os.Exit(0)
}
