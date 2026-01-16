package detector

import "regexp"

type MaskStrategy int

type Detector struct {
	Type string
	Re   *regexp.Regexp
	Mask MaskStrategy
}
