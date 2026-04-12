package main

import (
	"os"
	"path/filepath"
	"testing"
)

// create a temp directory and file with AWS key in it.
func TestScanFile_DetectsAWSKey(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")

	content := "hello\naws=AKIA1234567890ABCDEF\nbye\n"

	// 0o644 is for file permission
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	findings := ScanFile(path)

	if len(findings) != 1 {
		t.Fatalf("len(findings) = %d, want 1", len(findings))
	}

	f := findings[0]

	if f.FindingType != "AWS_ACCESS_KEY_ID" {
		t.Fatalf("FindingType = %q, want %q", f.FindingType, "AWS_ACCESS_KEY_ID")
	}

	// check if the finding line is correct
	// that
	if f.Line != 2 {
		t.Fatalf("Line = %d, want 2", f.Line)
	}

	// verify if the scanner captured exact secret text
	if f.Match != "AKIA1234567890ABCDEF" {
		t.Fatalf("Match = %q, want %q", f.Match, "AKIA1234567890ABCDEF")
	}

}

func TestScanFile_DoesNotDetectInvalidAWSKey(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "invalid.txt")

	content := "hello\naws=AKIA1234\nbye\n"

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}
}
