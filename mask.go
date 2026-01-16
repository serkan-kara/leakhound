package main

import (
	"strings"

	"github.com/serkan-kara/leakhound/detector"
)

func MaskWithStrategy(val string, strategy detector.MaskStrategy) string {
	switch strategy {
	case MaskRedact:
		return "REDACTED"
	case MaskJWT:
		return maskCustom(val, 3, 10)
	case MaskDefault:
		fallthrough
	default:
		return maskCustom(val, 4, 4)
	}
}

// mask secret values on outputs

func maskCustom(val string, prefix, suffix int) string {
	// mask everything if its shorter than 6
	if len(val) <= 6 {
		return strings.Repeat("*", len(val))
	}

	if len(val) <= prefix+suffix || len(val) <= 6 {
		return strings.Repeat("*", len(val))
	}

	middleLen := len(val) - (prefix + suffix)

	return val[:prefix] + strings.Repeat("*", middleLen) + val[len(val)-suffix:]
}
