package userservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type GetUserResponse struct {
	User
	ErrorMessage string `json:"error_message"`
}

type GetUserService struct {
	UserID   int `json:"user_id"`
	Response GetUserResponse
}

func (service GetUserService) Request(config *util.Config) (GetUserResponse, error) {
	getUserResponse := GetUserResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	getUserUrl := config.MICROSERVICESBASEURL["Usermicroservices"] + "/GetUser"

	request, err := http.NewRequest("POST", getUserUrl, bytes.NewBuffer(serviceJson))
	if err != nil {
		return getUserResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return getUserResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to get user`
		getUserResponse.ErrorMessage = errStr

		return getUserResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&getUserResponse)

	return getUserResponse, nil
}
