package utils;

import (
	"fmt"
)

func KnuthMorrisPrattMatch(text string, pattern string) int {

	var n int = len(text);
	var m int = len(pattern);

	if ( m <= n){
		b := computeBorder(pattern);

		fmt.Println(text, "===", pattern)
		fmt.Println(b)
	
		var i int = 0;
		var j int = 0;
	
		for ( i < n){
			if (pattern[j] == text[i]){
				if (j ==  m -1){
					return i - m + 1; // match
				}
				i++;
				j++;
			} else if (j > 0) {
				j = b[j-1]
			} else {
				i++;
			}
		}

		return -1; 
	}
	return -1; // no match
}

func computeBorder(pattern string) []int {

	if (len(pattern) > 1){
		var m int = len(pattern)
		var b = make([]int, m-1); // array of b with the length of (pattern -1)
	
		b[0] = 0;
		for k := 1; k < m-1 ; k++{
			s1 := pattern[0:k+1] // the one to search its prefix
			s2 := pattern[1:k+1] // the one to search its suffix
			b[k] = getMaximumPrefixSuffixSimilar(s1,s2)
		}
	
		return b;
	} else {
		return []int{0};
	}

}

func getMaximumPrefixSuffixSimilar(s1 string, s2 string) int {
	for i := 0; i < len(s2) ; i++ {
		temp1 := s1[:i+1]
		temp2 := s2[len(s2)-i-1:]
		if (isStringMatch(temp1, temp2)){
			return len(temp1)
		}
	}
	return 0;
}

func isStringMatch(s1 string, s2 string) bool {

	if (len(s1) != len(s2)){
		return false;
	} else {
		for i:= 0; i < len(s1); i++ {
			if (s1[i] != s2[i]){
				return false;
			}
		}
		return true;
	}
}

	// var j int = 0
	// var i int = 0

	// for ( i <  m){
	// 	if (pattern[j] == pattern[i]){
	// 		b[i] = j + 1;
	// 		i++;
	// 		j++;
	// 	} else if ( j > 0){
	// 		j = b[j-1];
	// 	} else{
	// 		b[i] = 0;
	// 		i++;
	// 	}
	// }