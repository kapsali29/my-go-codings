package app

import (
	"fmt"
	"stats-go-sdk/internal/domain"
)

func PrintUsers(users []domain.User) {
	for _, val := range users {
		fmt.Println("----")
		fmt.Printf(
			"userId: %d\nuserName: %s\nuserType: %s \n",
			val.Id,
			val.Name,
			val.Type,
		)
		fmt.Println("----")
	}
}