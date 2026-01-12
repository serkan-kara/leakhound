package main

import (
	"bufio"
	"os"
	"regexp"
)

func ScanFile(path string) []string {
	var findings []string

	file, err := os.Open(path)

	if err != nil {
		return findings
	}
	defer file.Close()

	re := regexp.MustCompile(`AKIA[0-9A-Z]{16}`) // aws key

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			findings = append(findings, line)
		}
	}

	return findings
}
