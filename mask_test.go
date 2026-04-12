package main

import (
	"testing"

	"github.com/serkan-kara/leakhound/detector"
)

func TestMaskWithStrategy_Default(t *testing.T) {

	tests := []struct {
		name     string
		in       string
		strategy detector.MaskStrategy
		want     string
	}{
		{
			name:     "default mask keeps first and last 4",
			in:       "AKIA1234567890CDEF",
			strategy: MaskDefault,
			want:     "AKIA**********CDEF",
		},
		{
			name:     "jwt mask keeps first 3 and last 10",
			in:       "eyJabcdefghij1234567890",
			strategy: MaskJWT,
			want:     "eyJ**********1234567890",
		},
		{
			name:     "redact always returns constant",
			in:       "anything",
			strategy: MaskRedact,
			want:     "REDACTED",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := MaskWithStrategy(testCase.in, testCase.strategy)
			if got != testCase.want {
				t.Fatalf("MaskWithStrategy(%q, %v) = %q, want %q", testCase.in, testCase.strategy, got, testCase.want)
			}
		})
	}
}
