package config

type SortBy string

const (
	DateFormat = "2006-01-02 15:04:05"
	ApiHost = "https://api.stackexchange.com"
	ApiVersion = "2.3"
	AnswersEndpoint = "answers"
	SortByActivity SortBy = "activity"
	SortByVersion SortBy = "version"
	SortByVotes SortBy = "votes"
	QuestionEndpoint = "questions"
)