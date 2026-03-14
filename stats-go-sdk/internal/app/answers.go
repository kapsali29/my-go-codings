package app

import (
	"cmp"
	"os"
	"slices"
	"stats-go-sdk/internal/domain"

	"github.com/markkurossi/tabulate"
)


func GetAnswerUniqueOwners(resp domain.Response) *domain.UserAnswerStats {
	var users []domain.User
	var usersMap = make(map[int]int)
	answers := resp.Answers
	for _, answer := range answers {
		user := answer.UserAnswered
		_,ok := usersMap[user.Id]
		if !ok {
			usersMap[user.Id] = 1
			users = append(users, user)
		} else {
			usersMap[user.Id] ++
		}
	}
	return &domain.UserAnswerStats{
		UserCounts: usersMap,
		UniqueIds: users,
	}
}

func containsQuestionById(stats []domain.QuestionAnswerStats, questionId int) int {
	var resFlag int = -1
	for idx, statRec := range stats {
		if statRec.QuestionId == questionId {
			resFlag = idx
		}
	}
	return resFlag
}

func createQuestionAnswers(
	QuestionId, NumberOfAnswers int,
	AcceptedAnswers []domain.Answer,
	) *domain.QuestionAnswerStats {
	return &domain.QuestionAnswerStats{
		QuestionId: QuestionId,
		NumberOfAnswers: NumberOfAnswers,
		AcceptedAnswers: AcceptedAnswers,
	}
}

func GetAcceptedAnswers(resp domain.Response) []domain.QuestionAnswerStats {
	var questionStats []domain.QuestionAnswerStats
	answers := resp.Answers
	for idx := range answers {
		isInStats := containsQuestionById(
			questionStats,
			answers[idx].QuestionId,
		)
		if isInStats == -1 {
			var AcceptedAnswers []domain.Answer
			if answers[idx].IsAccepted {
				AcceptedAnswers = append(AcceptedAnswers, answers[idx])
			}
			newQuestionStat := *createQuestionAnswers(
				answers[idx].QuestionId,
				1,
				AcceptedAnswers,
			)
			questionStats = append(questionStats, newQuestionStat)
		} else {
			statToUpdate := &questionStats[isInStats]
			statToUpdate.NumberOfAnswers ++
			var AcceptedAnswers []domain.Answer
			if answers[idx].IsAccepted {
				AcceptedAnswers = append(AcceptedAnswers, answers[idx])
			}
		}
	}
	slices.SortFunc(questionStats, func(a, b domain.QuestionAnswerStats) int {
		return cmp.Compare(b.NumberOfAnswers, a.NumberOfAnswers)
	})
	return questionStats
}

func TabulateAcceptedAnswers(stats []domain.QuestionAnswerStats) {
	tab := tabulate.New(tabulate.ASCII)
	tab.Header("N.A").SetAlign(tabulate.ML)
	tab.Header("Question Stats")
	err := tabulate.Reflect(tab, 0, nil, stats)
	if err != nil {
		panic(err)
	}
	tab.Print(os.Stdout)
}