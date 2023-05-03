package test

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/utils"
	"testing"
)

func TestBoyerMooreMatch(t *testing.T) {
	testCases := []struct {
		text     string
		pattern  string
		expected bool
	}{
		{"", "", true},
		{"abcde", "cd", true},
		{"abcde", "cf", false},
		{"abcde", "ab", true},
		{"abcde", "ed", false},
		{"abcde", "abcde", true},
		{"abcde", "abcdef", false},
		// Edge cases
		{"a", "a", true},          // single character
		{"a", "b", false},         // single character mismatch
		{"ab", "bc", false},       // two characters, no match
		{"ab", "ba", false},       // two characters, reverse order
		{"aa", "a", true},         // repeating characters, match
		{"aa", "b", false},        // repeating characters, mismatch
		{"abcdefg", "abc", true},  // pattern at beginning
		{"abcdefg", "efg", true},  // pattern at end
		{"abcdefg", "def", true},  // pattern in middle
		{"abcdefg", "hij", false}, // pattern not in text
		// Real Text
		{"hello world", "world", true},
		{"the quick brown fox jumps over the lazy dog", "dog", true},
		{"programming is fun", "ing", true},
		{"to be or not to be", "not", true},
		{"the cake is a lie", "pie", false},
		{"don't stop believing", "leaving", false},
		// Large text
		{largeText, "example", false},
	}

	for _, tc := range testCases {
		if got := utils.BoyerMooreMatch(tc.text, tc.pattern); got != tc.expected {
			t.Errorf("BoyerMooreMatch(%q, %q) = %v, expected %v", tc.text, tc.pattern, got, tc.expected)
		}
	}
}

const largeText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus vel risus eget ipsum eleifend scelerisque. Nunc eleifend, ante a hendrerit pellentesque, erat urna placerat erat, ac semper neque odio vitae nunc. Quisque ut purus eget augue bibendum blandit a vel enim. Suspendisse tincidunt augue eget ex suscipit faucibus. Donec euismod odio quam, a egestas risus egestas in. Integer ultrices augue vel ex mollis interdum. Donec sit amet semper nulla. Ut id nisl nec felis finibus elementum. Praesent bibendum non erat a lacinia. Pellentesque sed mi suscipit, pharetra libero at, ullamcorper purus. Duis ultricies mi ipsum, non auctor augue bibendum vitae. Proin vulputate consequat quam, vel aliquam nulla elementum nec. Duis euismod eget"
