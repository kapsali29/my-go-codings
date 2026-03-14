package main

import (
	//"fmt"
	"flag"
	"fmt"
	"os"
	"stats-go-sdk/internal/api"
	"stats-go-sdk/internal/app"
	"stats-go-sdk/internal/utils"
)

func AnswersCmd(from, to string, printUsers bool) {
	//fromDateUnix := utils.ToUnix("2026-03-05 10:00:00")
	//toDateUnix := utils.ToUnix("2026-03-07 23:00:00")
	toDateUnix := utils.ToUnix(to)
	fromDateUnix := utils.ToUnix(from)
	answersResponse := api.GetAnswers(fromDateUnix, toDateUnix)
	answerStats := app.GetAnswerUniqueOwners(answersResponse)
	if printUsers {
	 	app.PrintUsers(answerStats.UniqueIds)
	}
	questionStats := app.GetAcceptedAnswers(answersResponse)
	app.TabulateAcceptedAnswers(questionStats)
}

func QuestionsCmd(questionId int, list, disQid, questionAnswers, getAccepted bool) {
	questionsApiResponse := api.ListQuestions()
	if list {
		uniqueQIds := app.GetUniqueQuestionIds(questionsApiResponse)
		fmt.Println(uniqueQIds)
	}
	if questionId >=0 {
		question, exists := app.DisplayQuestionData(
			questionId,
			questionsApiResponse.Questions,
		)
		if !exists {
			fmt.Printf("Question with Id:%v does not exists", questionId)
		} else {
			if disQid {
				question.PrettyPrintQuestion()
			}
			if questionAnswers {
				questionAnswer := api.QuestionAnswers(question.QuestionId)
				if getAccepted {
					has, _ := app.QuestionHasAcceptedAnswers(
						question.QuestionId,
						questionAnswer.Answers,
					)
					fmt.Println(has)
				} else {
					fmt.Println(questionAnswer)
				}
			}
		}
	}
}

func main() {
	aCmd := flag.NewFlagSet("answers", flag.ExitOnError)
	from := aCmd.String("from", "", "Start Date")
	to := aCmd.String("to", "", "End time")
	printUsers := aCmd.Bool("print-users", false, "Print users")

	qCmd := flag.NewFlagSet("questions", flag.ExitOnError)
	list := qCmd.Bool("list", false, "List Questions")
	qId := qCmd.Int("id", -1, "Question Id")
	disQid := qCmd.Bool("display", false, "Display Question Data")
	questionAnswers :=qCmd.Bool("answers", false, "Retrieve Question Answers")
	getAccepted := qCmd.Bool("accepted", false, "Retrieve Accepted Answers")

	switch os.Args[1] {
	case "answers":
		aCmd.Parse(os.Args[2:])
		AnswersCmd(*from, *to, *printUsers)
	case "questions":
		qCmd.Parse(os.Args[2:])
		QuestionsCmd(*qId, *list, *disQid, *questionAnswers, *getAccepted)
	default:
		fmt.Println("Not valid command")
	}

}