package utils

func BoyerMooreMatch(text, pattern string) bool {
	last := buildLast(pattern)
	n := len(text)
	m := len(pattern)

	switch {
	case m > n:
		return false
	case m == n:
		return text == pattern
	}

	i := m - 1
	j := m - 1
	for i <= n-1 {
		if pattern[j] == text[i] {
			if j == 0 {
				return true
			}
			i--
			j--
		} else {
			lastOccurrence := last[text[i]]
			i = i + m - min(j, 1+lastOccurrence)
			j = m - 1
		}
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func buildLast(pattern string) []int {

	last := make([]int, 256)
	for i := 0; i < 256; i++ {
		last[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		last[pattern[i]] = i
	}
	return last
}
