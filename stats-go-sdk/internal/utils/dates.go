package utils

import (
	"stats-go-sdk/internal/config"
	"time"
)

func ToUnix(aDate string) int {
	t, err := time.Parse(config.DateFormat, aDate)
	if err != nil {
		panic(err)
	} else {
		return int(t.Unix())
	}
}