package utils

import (
	"math"
	"sort"
)

type SimilarityScore struct {
	Question string
	Score    float64
}

func SortSimilarityScores(scores []SimilarityScore) []SimilarityScore {
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})
	return scores
}

// referensi https://www.youtube.com/watch?v=We3YDTzNXEk
func Levenshtein_distance(stringOne, stringTwo string) int {
	lengthStringOne := len(stringOne)
	lengthStringTwo := len(stringTwo)

	if lengthStringOne == 0 {
		return lengthStringTwo
	}

	if lengthStringTwo == 0 {
		return lengthStringOne
	}

	distanceMatrix := make([][]int, lengthStringOne+1)
	for i := 0; i <= lengthStringOne; i++ {
		distanceMatrix[i] = make([]int, lengthStringTwo+1)
	}

	for i := 0; i <= lengthStringOne; i++ {
		distanceMatrix[i][0] = i
	}

	for j := 0; j <= lengthStringTwo; j++ {
		distanceMatrix[0][j] = j
	}

	// computing distance
	for j := 1; j <= lengthStringTwo; j++ {
		for i := 1; i <= lengthStringOne; i++ {
			if stringOne[i-1] == stringTwo[j-1] {
				distanceMatrix[i][j] = distanceMatrix[i-1][j-1]
			} else {
				distanceMatrix[i][j] = 1 + int(math.Min(float64(distanceMatrix[i-1][j]), math.Min(float64(distanceMatrix[i][j-1]), float64(distanceMatrix[i-1][j-1]))))
			}
		}
	}
	return distanceMatrix[lengthStringOne][lengthStringTwo]
}

func Similarity(textOne, textTwo string) float64 {
	if textOne == "" && textTwo == "" {
		return 100
	}

	distance := Levenshtein_distance(textOne, textTwo)

	if len(textOne) == 0 && len(textTwo) == 0 {
		return 100
	} else if len(textOne) == 0 || len(textTwo) == 0 {
		return 0
	}

	maxLength := max(len(textOne), len(textTwo))
	lenResult := 100 - 100*math.Abs(float64(len(textOne)-len(textTwo))/float64(maxLength))
	levenshteinResult := (1 - float64(distance)/float64(maxLength)) * 100
	result := 0.2*lenResult + 0.8*levenshteinResult
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
