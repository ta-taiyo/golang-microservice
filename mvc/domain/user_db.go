package domain

import (
	"fmt"
	"net/http"

	"github.com/ta-taiyo/golang-microservice/mvc/utils"
)

// DB の変わり
var users = map[int64]*User{
	123: {Id: 123, FirstName: "Fede", LastName: "Leon", Email: "test@example.com"},
}

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	// fmt.Println("test")
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
