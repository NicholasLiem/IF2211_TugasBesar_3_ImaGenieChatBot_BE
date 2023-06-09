package user_query

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/utils"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/query_utils"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
	"strconv"
	"strings"
)

func QAStringMatchingHandler(query string, patternType string) (string, error) {
	qas, err := query_utils.GetAllQuestionAnswers()
	if err != nil {
		return "", err
	}

	for _, qa := range qas {
		if patternType == "BM" {
			if utils.BoyerMooreMatch(qa.Question, query) {
				return qa.Answer, nil
			}
		} else {
			if utils.KnuthMorrisPrattMatch(qa.Question, query) != -1 {
				return qa.Answer, nil
			}
		}
	}

	// Sorting similarity score
	similarities := make([]utils.SimilarityScore, len(qas))

	// Handle empty db
	if len(similarities) == 0 {
		return "You have not added any questions to the database, please add them first.", nil
	}

	for i, qa := range qas {
		score := utils.Similarity(query, qa.Question)
		similarities[i] = utils.SimilarityScore{Question: qa.Question, Score: score}
	}

	utils.SortSimilarityScores(similarities)

	// TO DO : get the top one if the similarity is >90%
	if similarities[0].Score > 90 {
		return getAnswerFromQuestion(similarities[0].Question, qas), nil
	} else {
		// Otherwise, Get top 3 most similar questions
		if len(similarities) > 3 {
			similarities = similarities[:3]
		}
	}

	var similarQuestions string
	for i, s := range similarities {
		if i == 0 {
			similarQuestions += strconv.Itoa(i+1) + ". " + strings.Title(s.Question)
		} else {
			similarQuestions += "\n" + strconv.Itoa(i+1) + ". " + strings.Title(s.Question)
		}
	}

	if len(similarities) > 0 {
		return "Sorry, I couldn't find the answer to your question. " + "\nHere are some similar questions:\n" + similarQuestions, nil
	} else {
		return "Sorry, I couldn't find the answer to your question.", nil
	}
}

func getAnswerFromQuestion(q string, ans []models.QuestionAnswer) string {
	for _, qa := range ans {
		if q == qa.Question {
			return qa.Answer
		}
	}
	return ""
}
