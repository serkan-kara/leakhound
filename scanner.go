package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/serkan-kara/leakhound/finding"
)

var detectors = []struct {
	Type string
	Re   *regexp.Regexp
}{
	{
		Type: "AWS_ACCESS_KEY_ID",
		Re:   regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
	},
}

func ScanFile(path string) []finding.Finding {
	var findings []finding.Finding

	file, err := os.Open(path)

	if err != nil {
		return findings
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNo := 0

	for scanner.Scan() {
		lineNo++
		line := scanner.Text()

		for _, detector := range detectors {
			matches := detector.Re.FindAllString(line, -1)
			for _, match := range matches {
				finding, err := finding.New(path, lineNo, detector.Type, match)

				if err != nil {
					fmt.Println(err)
				}

				findings = append(findings, finding)
			}
		}
	}

	return findings
}
