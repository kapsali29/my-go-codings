package api

import (
	"net/http"
	"fmt"
	"stats-go-sdk/internal/config"
	"encoding/json"
	"stats-go-sdk/internal/domain"
)

func SetBaseUrl(endpoint string) string{
	baseUrl := fmt.Sprintf(
		"%v/%v/%v",
		config.ApiHost,
		config.ApiVersion,
		endpoint,
	)
	return baseUrl
}

func GetAnswers(fromDate, toDate int) domain.Response{
	var apiResp domain.Response
	baseUrl := SetBaseUrl(config.AnswersEndpoint)
	fullUrl := fmt.Sprintf(
		"%v?fromdate=%v&order=desc&max=%v&sort=activity&site=stackoverflow",
		baseUrl,
		fromDate,
		toDate,
	)
	resp, err := http.Get(fullUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	jsonErr := json.NewDecoder(resp.Body).Decode(&apiResp)
	if jsonErr != nil {
		panic(jsonErr)
	}
	return apiResp
}
