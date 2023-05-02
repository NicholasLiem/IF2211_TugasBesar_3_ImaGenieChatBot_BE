package test

import (
	"fmt"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/utils"
	"testing"
)

func TestLevenshteinDistanceAndSimilarity(t *testing.T) {
	testCases := []struct {
		textOne    string
		textTwo    string
		distance   int
		similarity float64
	}{
		{"", "", 0, 100},
		{"abcde", "abcde", 0, 100},
		{"abcde", "abcd", 1, 80},
		{"abcd", "abcde", 1, 80},
		{"abc", "def", 3, 0},
		{"kitten", "sitting", 3, 57.14285714285714},
		{"", "xyz", 3, 0},
		{"abc", "", 3, 0},
		{"Apa kabar", "b", 0, 0.3},
		{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer vestibulum vel augue eget malesuada. Sed ac libero euismod, fermentum risus vel, ultrices ipsum. Nullam vel aliquet risus. Sed malesuada, nulla non tristique faucibus, nibh nibh convallis nisl, a suscipit risus nulla in diam. Praesent eget quam at nulla consectetur finibus. Suspendisse eget ultricies nulla, vel maximus elit. Nulla facilisi. Donec volutpat tortor ac velit gravida aliquet. Duis ut ipsum feugiat, interdum sapien vel, maximus ipsum. Duis vel lorem ut lectus commodo faucibus. Vivamus pretium aliquet bibendum. Donec lobortis est eget ex pharetra, vel commodo lorem blandit. In hac habitasse platea dictumst. Proin vel ante eleifend, lobortis nulla in, faucibus elit.",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer vestibulum vel augue eget malesuada. Sed ac libero euismod, fermentum risus vel, ultrices ipsum. Nullam vel aliquet risus. Sed malesuada, nulla non tristique faucibus, nibh nibh convallis nisl, a suscipit risus nulla in diam. Praesent eget quam at nulla consectetur finibus. Suspendisse eget ultricies nulla, vel maximus elit. Nulla facilisi. Donec volutpat tortor ac velit gravida aliquet. Duis ut ipsum feugiat, interdum sapien vel, maximus ipsum. Duis vel lorem ut lectus commodo faucibus. Vivamus pretium aliquet bibendum. Donec lobortis est eget ex pharetra, vel commodo lorem blandit. In hac habitasse platea dictumst. Proin vel ante eleifend, lobortis nulla in, faucibus elit.",
			0,
			100,
		},
	}

	for _, tc := range testCases {
		distance := utils.Levenshtein_distance(tc.textOne, tc.textTwo)
		if distance != tc.distance {
			t.Errorf("Levenshtein_distance(%q, %q) = %d, expected %d", tc.textOne, tc.textTwo, distance, tc.distance)
		}

		similarity := utils.Similarity(tc.textOne, tc.textTwo)
		if similarity != tc.similarity {
			t.Errorf("Similarity(%q, %q) = %f, expected %f", tc.textOne, tc.textTwo, similarity, tc.similarity)
		}
	}
}

func TestFindMostSimilarQuestion(t *testing.T) {
	// Daftar pertanyaan yang tersimpan dalam sistem
	questions := []string{
		"Apa itu Go?",
		"Cara mengganti tampilan di Go?",
		"Apakah Go adalah bahasa pemrograman baru?",
		"Bagaimana cara mengatasi error di Go?",
		"Apa kelebihan dari Go dibandingkan bahasa pemrograman lain?",
		"Kapan saya harus menggunakan Go?",
		"Bagaimana cara menginstall Go?",
		"Apakah Go sulit dipelajari?",
	}

	// Query yang diberikan oleh pengguna
	query := "Cara mengganti format tanggal di Go?"

	// Mencari pertanyaan yang paling mirip dengan query pengguna
	minDistance := -1
	var mostSimilarQuestion string
	for _, question := range questions {
		similarity := utils.Similarity(query, question)
		fmt.Printf("Similarity between query '%s' and question '%s' is %f\n", query, question, similarity)
		distance := utils.Levenshtein_distance(query, question)
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
			mostSimilarQuestion = question
		}
	}

	// Menguji apakah pertanyaan yang ditemukan sudah benar
	expectedQuestion := "Cara mengganti tampilan di Go?"
	if mostSimilarQuestion != expectedQuestion {
		t.Errorf("FindMostSimilarQuestion(%q) = %q, expected %q", query, mostSimilarQuestion, expectedQuestion)
	}
}
