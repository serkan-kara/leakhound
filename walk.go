package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/serkan-kara/leakhound/finding"
)

func ScanPath(path string, excludes []string) (finding.ScanResult, error) {
	start := time.Now()
	res := finding.ScanResult{} // findings nil, filesscanned 0
	defer func() { res.Duration = time.Since(start) }()

	info, err := os.Stat(path)
	if err != nil {
		return res, ErrRuntime
	}

	if !info.IsDir() {
		// single file exclude control
		if shouldExclude(path, excludes) {
			return res, nil
		}

		res.FilesScanned++
		findings := ScanFile(path)
		res.Findings = append(res.Findings, findings...)

		return res, nil
	}

	walkErr := filepath.WalkDir(path, func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return ErrRuntime
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

		res.FilesScanned++
		findings := ScanFile(p)
		res.Findings = append(res.Findings, findings...)

		return nil
	})

	if walkErr != nil {
		return res, ErrRuntime
	}

	return res, nil
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
