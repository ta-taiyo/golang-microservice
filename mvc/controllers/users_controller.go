package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ta-taiyo/golang-microservice/mvc/services"
	"github.com/ta-taiyo/golang-microservice/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64) // 10進数, bitSize 64
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		jsonValue, _ := json.Marshal(apiErr)

		// resp.WriteHeader(http.StatusNotFound)
		// id=123A など数値以外が入ると BadRequest (400)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}

//
