package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/serkan-kara/leakhound/finding"
)

func ScanPath(path string, excludes []string) []finding.Finding {
	var all []finding.Finding

	info, err := os.Stat(path)
	if err != nil {
		return all
	}

	if !info.IsDir() {
		// single file exclude control
		if shouldExclude(path, excludes) {
			return all
		}
		return ScanFile(path)
	}

	_ = filepath.WalkDir(path, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}

		// if path contains fragment skip
		if shouldExclude(p, excludes) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// skip some folders like node_modules etc.
		if d.IsDir() {
			name := d.Name()
			if isIgnoredDir(name) {
				return filepath.SkipDir
			}
			return nil
		}

		// skip some files like jpg, png etc.
		if isLikelyBinaryByExt(p) {
			return nil
		}

		findings := ScanFile(p)
		all = append(all, findings...)

		return nil
	})

	return all
}

// if any path inside exlude list returns true
// for example it returns true for excludes=["testFiles"] and path=".../testFiles/a.txt"
func shouldExclude(path string, excludes []string) bool {
	if len(excludes) == 0 {
		return false
	}

	p := filepath.ToSlash(path)
	for _, exclude := range excludes {
		exclude = strings.TrimSpace(exclude)
		if exclude == "" {
			continue
		}
		if strings.Contains(p, exclude) {
			return true
		}
	}

	return false
}

func isIgnoredDir(name string) bool {
	ignored := []string{
		".git", "node_modules", "vendor", ".idea", ".vscode", "dist", "build",
	}
	for _, x := range ignored {
		if name == x {
			return true
		}
	}
	return false
}

func isLikelyBinaryByExt(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	binExt := []string{
		".png", ".jpg", ".jpeg", ".gif", ".webp",
		".pdf", ".zip", ".gz", ".tar", ".7z",
		".exe", ".dll", ".so", ".dylib",
	}
	for _, b := range binExt {
		if ext == b {
			return true
		}
	}
	return false
}
