package app

import (
	"fmt"
	"stats-go-sdk/internal/domain"
)

func GetUniqueQuestionIds(resp domain.QuestionApiResponse) []int {
	var questionIds []int
	var seen map[int]struct{}
	for _, rec := range resp.Questions {
		_, exists := seen[rec.QuestionId]
		if !exists {
			questionIds = append(questionIds, rec.QuestionId)
		}
	}
	return questionIds
}

func DisplayQuestionData(
	questionId int,
	listOfQuestions []domain.Question,
	) (*domain.Question, bool) {
		for _, question := range listOfQuestions {
			if question.QuestionId == questionId {
				return &question, true
			}
		}
		return &domain.Question{},false
}

func QuestionHasAcceptedAnswers(questionId int, answers []domain.Answer) (bool, []domain.AcceptedAnswersInQuestion) {
	var acceptedAnswers []domain.AcceptedAnswersInQuestion
	for _, val := range answers {
		if val.IsAccepted {
			acceptedAnswer := domain.AcceptedAnswersInQuestion{
				Score: val.Score,
				AnswerId: val.Id,
				AnswerOwner: fmt.Sprintf(
					"Id: %d -- Name: %s",
					val.UserAnswered.Id,
					val.UserAnswered.Name,
				),
			}
			acceptedAnswers = append(acceptedAnswers, acceptedAnswer)
		}
	}
	if len(acceptedAnswers) >0 {
		return true, acceptedAnswers
	} else {
		return false, acceptedAnswers
	}
}