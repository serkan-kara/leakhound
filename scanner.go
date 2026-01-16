package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/serkan-kara/leakhound/detector"
	"github.com/serkan-kara/leakhound/finding"
)

const (
	MaskDefault detector.MaskStrategy = iota // first 4 - last 4
	MaskJWT                                  // first 3 - last 10
	MaskRedact                               // redacted
)

var detectors = []detector.Detector{
	{
		Type: "AWS_ACCESS_KEY_ID",
		Re:   regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
		Mask: MaskDefault,
	},
	{
		Type: "JWT",
		// Basit JWT pattern: header.payload.signature (base64url)
		Re:   regexp.MustCompile(`eyJ[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+`),
		Mask: MaskJWT,
	},
	{
		Type: "PRIVATE_KEY",
		// Private key başlangıcı (satır içi yakalayacağız)
		Re:   regexp.MustCompile(`-----BEGIN (RSA |EC |OPENSSH )?PRIVATE KEY-----`),
		Mask: MaskRedact,
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
				finding, err := finding.New(path, lineNo, detector.Type, match, detector.Mask)

				if err != nil {
					fmt.Println(err)
				}

				findings = append(findings, finding)
			}
		}
	}

	return findings
}
