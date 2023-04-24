package levenshtein

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/utils"
	"testing"
)

func TestDistance(t *testing.T) {
	testCases := []struct {
		s1       string
		s2       string
		expected int
	}{
		{"kitten", "sitting", 3},
		{"rosettacode", "raisethysword", 8},
		{"hello", "hello", 0},
		{"hello", "", 5},
		{"", "", 0},
	}

	for _, tc := range testCases {
		if got := utils.Levenshtein_distance(tc.s1, tc.s2); got != tc.expected {
			t.Errorf("Distance(%q, %q) = %d, want %d", tc.s1, tc.s2, got, tc.expected)
		}
	}
}
