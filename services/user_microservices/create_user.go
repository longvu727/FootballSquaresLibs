package userservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type CreateUserResponse struct {
	UserID       int    `json:"user_id"`
	ErrorMessage string `json:"error_message"`
}

type CreateUserService struct {
	Response CreateUserResponse
}

func (service CreateUserService) Request(config *util.Config) (CreateUserResponse, error) {
	createUserResponse := CreateUserResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	createUserUrl := config.MICROSERVICESBASEURL["usermicroservices"] + "/CreateUser"

	request, err := http.NewRequest("POST", createUserUrl, bytes.NewBuffer(serviceJson))
	if err != nil {
		return createUserResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return createUserResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to create user`
		createUserResponse.ErrorMessage = errStr

		return createUserResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&createUserResponse)

	return createUserResponse, nil
}
