package finding

import (
	"errors"

	"github.com/serkan-kara/leakhound/detector"
)

type Finding struct {
	File        string
	Line        int
	FindingType string
	Match       string
	Mask        detector.MaskStrategy
}

func New(file string, line int, findingType string, match string, maskStrategy detector.MaskStrategy) (Finding, error) {
	if file == "" || findingType == "" || match == "" {
		return Finding{}, errors.New("File, finding type, match can not be empty")
	}

	return Finding{
		File:        file,
		Line:        line,
		FindingType: findingType,
		Match:       match,
		Mask:        maskStrategy,
	}, nil
}
