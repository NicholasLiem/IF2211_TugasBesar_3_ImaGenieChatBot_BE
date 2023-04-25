package main

import (
	"fmt"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/utils"
)

func main (){
	testCases := []struct {
		text     string
		pattern  string
	}{
		{"aaaaabaaba", "abaaba"},
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
			fmt.Println(" The string does not match.")
		} else {
			fmt.Printf(" The string is match at position %d \n", res)
		}
	}

}
const largeText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus vel risus eget ipsum eleifend scelerisque. Nunc eleifend, ante a hendrerit pellentesque, erat urna placerat erat, ac semper neque odio vitae nunc. Quisque ut purus eget augue bibendum blandit a vel enim. Suspendisse tincidunt augue eget ex suscipit faucibus. Donec euismod odio quam, a egestas risus egestas in. Integer ultrices augue vel ex mollis interdum. Donec sit amet semper nulla. Ut id nisl nec felis finibus elementum. Praesent bibendum non erat a lacinia. Pellentesque sed mi suscipit, pharetra libero at, ullamcorper purus. Duis ultricies mi ipsum, non auctor augue bibendum vitae. Proin vulputate consequat quam, vel aliquam nulla elementum nec. Duis euismod eget"
