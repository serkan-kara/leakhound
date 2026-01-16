package main

import "strings"

// mask secret values on outputs
func MaskValue(val string) string {

	// mask everything if its shorter than 6
	if len(val) <= 6 {
		return strings.Repeat("*", len(val))
	}

	prefix := 4
	suffix := 4

	if len(val) < prefix+suffix {
		return strings.Repeat("*", len(val))
	}

	middleLen := len(val) - (prefix + suffix)

	return val[:prefix] + strings.Repeat("*", middleLen) + val[len(val)-suffix:]
}
