package domain

import (
	"fmt"
)

type User struct {
	Id          int      `json:"user_id"`
	Name        string   `json:"display_name"`
	Reputation  int      `json:"reputation"`
    Type        string   `json:"user_type"`
}

type Answer struct {
	Id            int    `json:"answer_id"`
	IsAccepted    bool   `json:"is_accepted"`
	Score         int    `json:"score"`
	Created       int    `json:"creation_date"`
	QuestionId    int    `json:"question_id"`
	License       string `json:"content_license"`
	UserAnswered  User   `json:"owner"`
}

type Response struct {
	Answers        []Answer `json:"items"`      
	HasMore        bool     `json:"has_more"`
	QuotaMax       int      `json:"quota_max"`
	QuotaRemaining int      `json:"quota_remaining"`
}

type UserAnswerStats struct {
	UserCounts   map[int]int   `json:"UserCounts"`
	UniqueIds    []User        `json:"UniqueIds"`
}

type QuestionAnswerStats struct {
	QuestionId int            `json:"QuestionId"`
	NumberOfAnswers int       `json:"NumberOfAnswers"`
	AcceptedAnswers []Answer  `json:"AcceptedAnswers"`
}

type Question struct {
	IsAnswered bool      `json:"is_answered"`
	QuestionId  int      `json:"question_id"`
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Score       int      `json:"score"`
	AnswerCount int      `json:"answer_count"`
	ViewCount   int      `json:"view_count"`
	Tags        []string `json:"tags"`
	Owner       User     `json:"owner"`
}

type QuestionApiResponse struct {
	Questions      []Question       `json:"items"`
	HasMore        bool             `json:"has_more"`
	QuotaMax       int              `json:"quota_max"`
	QuotaRemaining int              `json:"quota_remaining"`
}

type AcceptedAnswersInQuestion struct {
	Score       int    `json:"score"`
	AnswerId    int    `json:"answer_id"`
	AnswerOwner string `json:"answer_owner"`
}

func (q *Question) PrettyPrintQuestion() {
	fmt.Printf(
		"Id:%d\nTitle:%s\nTags:%v\nScore:%v\nIsAnswered:%v\nAnswerCount:%v\nViewCount:%v\nOwnerId:%v\nLink:%v\n",
		q.QuestionId,
		q.Title,
		q.Tags,
		q.Score,
		q.IsAnswered,
		q.AnswerCount,
		q.ViewCount,
		q.Owner.Id,
		q.Link,
	)
}