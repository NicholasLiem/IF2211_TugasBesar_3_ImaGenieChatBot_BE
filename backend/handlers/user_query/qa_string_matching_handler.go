package user_query

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/utils"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/query_utils"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/models"
)

func QAStringMatchingHandler(query string) (string, error) {
	qas, err := query_utils.GetAllQuestionAnswers()
	if err != nil {
		return "", err
	}

	for _, qa := range qas {
		if utils.BoyerMooreMatch(query, qa.Question) {
			return qa.Answer, nil
		}
		// if (utils.KnuthMorrisPrattMatch(query, qa.Question) != -1){
		// 	return qa.Answer, nil
		// }
	}

	// Sorting similarity score
	similarities := make([]utils.SimilarityScore, len(qas))

	// Handle empty db
	if len(similarities) == 0 {
		return "Belum ada pertanyaan mas :(", nil
	}

	for i, qa := range qas {
		score := utils.Similarity(query, qa.Question)
		similarities[i] = utils.SimilarityScore{Question: qa.Question, Score: score}
	}

	utils.SortSimilarityScores(similarities)

	// TO DO : get the top one if the similiarity is >90%
	if similarities[0].Score > 90 {
		return getAnswerFromQuestion(similarities[0].Question, qas), nil
	} else {
		// Otherwise, Get top 3 most similar questions
		if len(similarities) > 3 {
			similarities = similarities[:3]
		}
	}

	similarQuestions := ""
	for _, s := range similarities {
		similarQuestions += s.Question + "\n"
	}

	return "Sorry, I couldn't find the answer to your question. Here are some similar questions: \n" + similarQuestions, nil
}

func getAnswerFromQuestion(q string, ans []models.QuestionAnswer) string {
	for _, qa := range ans {
		if q == qa.Question {
			return qa.Answer
		}
	}
	return ""
}
