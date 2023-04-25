package test

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/utils"
	"testing"
)

func TestKMPMatch(t *testing.T){
	testCases := []struct {
		text     string
		pattern  string
	}{
		{"", ""},
		{"abcde", "cd"},
		{"abcde", "cf"},
		{"abcde", "ab"},
		{"abcde", "ed"},
		{"abcde", "abcde"},
		{"abcde", "abcdef"},
		// Edge cases
		{"a", "a"},          // single character
		{"a", "b"},         // single character mismatch
		{"ab", "bc"},       // two characters, no match
		{"ab", "ba"},       // two characters, reverse order
		{"aa", "a"},         // repeating characters, match
		{"aa", "b"},        // repeating characters, mismatch
		{"abcdefg", "abc"},  // pattern at beginning
		{"abcdefg", "efg"},  // pattern at end
		{"abcdefg", "def"},  // pattern in middle
		{"abcdefg", "hij"}, // pattern not in text
		// Real Text
		{"hello world", "world"},
		{"the quick brown fox jumps over the lazy dog", "dog"},
		{"programming is fun", "ing"},
		{"to be or not to be", "not"},
		{"the cake is a lie", "pie"},
		{"don't stop believing", "leaving"},
		// Large text
		{largeText, "example"},
	}

	for _, tc := range testCases{
		if res := utils.KnuthMorrisPrattMatch(tc.text, tc.pattern) ; res == -1 {
			t.Errorf(" The string does not match.")
		} else {
			t.Errorf(" The string is match at position %d", res)
		}
	}
}

