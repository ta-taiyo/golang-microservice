package services

import (
	"github.com/ta-taiyo/golang-microservice/mvc/domain"
	"github.com/ta-taiyo/golang-microservice/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)

}
