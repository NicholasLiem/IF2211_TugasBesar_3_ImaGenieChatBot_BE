package utils

import "math"

// referensi https://www.youtube.com/watch?v=We3YDTzNXEk
func Levenshtein_distance(stringOne, stringTwo string) int {
	lengthStringOne := len(stringOne)
	lengthStringTwo := len(stringTwo)

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
	distance := Levenshtein_distance(textOne, textTwo)
	maxLength := math.Max(float64(len(textOne)), float64(len(textTwo)))
	return (1 - float64(distance)/maxLength) * 100
}
