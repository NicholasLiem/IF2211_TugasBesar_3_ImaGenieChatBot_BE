package user_query

import (
	"github.com/NicholasLiem/Tubes3_ImagineKelar/algorithms/utils"
	"github.com/NicholasLiem/Tubes3_ImagineKelar/handlers/query_utils"
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
	}

	// Find top 3 most similar questions
	similarities := make([]utils.SimilarityScore, len(qas))
	for i, qa := range qas {
		score := utils.Similarity(query, qa.Question)
		similarities[i] = utils.SimilarityScore{Question: qa.Question, Score: score}
	}

	utils.SortSimilarityScores(similarities)

	if len(similarities) > 3 {
		similarities = similarities[:3]
	}

	similarQuestions := ""
	for _, s := range similarities {
		similarQuestions += s.Question + "\n"
	}

	return "Sorry, I couldn't find the answer to your question. Here are some similar questions: \n" + similarQuestions, nil
}
