package api

import (
	"net/http"
	"fmt"
	"stats-go-sdk/internal/config"
	"encoding/json"
	"stats-go-sdk/internal/domain"
)

func ListQuestions() domain.QuestionApiResponse {
	var qApiResp domain.QuestionApiResponse
	fullUrl := fmt.Sprintf(
		"%v?order=desc&sort=activity&site=stackoverflow",
		SetBaseUrl(config.QuestionEndpoint),
	)
	resp, err := http.Get(fullUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&qApiResp)
	return qApiResp
}

func QuestionAnswers(questionId int) domain.Response {
	var jsonRes domain.Response
	fullUrl := fmt.Sprintf(
		"%v/%d/answers?order=desc&sort=activity&site=stackoverflow",
		SetBaseUrl(config.QuestionEndpoint),
		questionId,
	)
	fmt.Println(fullUrl)
	resp, err := http.Get(fullUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&jsonRes)
	return jsonRes
}