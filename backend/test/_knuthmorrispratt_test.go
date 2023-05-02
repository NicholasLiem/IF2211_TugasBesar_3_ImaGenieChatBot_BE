package test

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/utils"
	"testing"
)

func TestKMPMatch(t *testing.T){
	testCases := []struct {
		text     string
		pattern  string
		pos int
	}{
		{"", "", 0},
		{"abcde", "cd", 2},
		{"abcde", "cf",-1},
		{"abcde", "ab", 0},
		{"abcde", "ed", -1},
		{"abcde", "abcde", 0},
		{"abcde", "abcdef", -1},
		// Edge cases
		{"a", "a", 0},          // single character
		{"a", "b", -1},         // single character mismatch
		{"ab", "bc", -1},       // two characters, no match
		{"ab", "ba", -1},       // two characters, reverse order
		{"aa", "a", 0},         // repeating characters, match
		{"aa", "b", -1},        // repeating characters, mismatch
		{"abcdefg", "abc", 0},  // pattern at beginning
		{"abcdefg", "efg", 4},  // pattern at end
		{"abcdefg", "def", 3},  // pattern in middle
		{"abcdefg", "hij", -1}, // pattern not in text
		// Real Text
		{"hello world", "world",6},
		{"the quick brown fox jumps over the lazy dog", "dog", 40},
		{"programming is fun", "ing", 8},
		{"to be or not to be", "not", 9},
		{"the cake is a lie", "pie",-1},
		{"don't stop believing", "leaving", -1},
		// Large text
		{largeText, "example",-1},
	}

	for _, tc := range testCases{
		if res := utils.KnuthMorrisPrattMatch(tc.text, tc.pattern) ; res != tc.pos {
			t.Errorf(" The string does not match.")
		}
	}
}

