package finding

import "errors"

type Finding struct {
	File        string
	Line        int
	FindingType string
	Match       string
}

func New(file string, line int, findingType string, match string) (Finding, error) {
	if file == "" || findingType == "" || match == "" {
		return Finding{}, errors.New("File, finding type, match can not be empty")
	}

	return Finding{
		File:        file,
		Line:        line,
		FindingType: findingType,
		Match:       match,
	}, nil
}
