package userservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type GetUserByGUIDResponse struct {
	User
	ErrorMessage string `json:"error_message"`
}

type GetUserByGUIDService struct {
	UserGUID string `json:"user_guid"`
}

func (service GetUserByGUIDService) Request(config *util.Config) (GetUserByGUIDResponse, error) {
	getUserByGUIDResponse := GetUserByGUIDResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	getUserByGUIDUrl := config.MICROSERVICESBASEURL["usermicroservices"] + "/GetUserByGUID"

	request, err := http.NewRequest("POST", getUserByGUIDUrl, bytes.NewBuffer(serviceJson))
	if err != nil {
		return getUserByGUIDResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return getUserByGUIDResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to get User by GUID`
		getUserByGUIDResponse.ErrorMessage = errStr

		return getUserByGUIDResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&getUserByGUIDResponse)

	return getUserByGUIDResponse, nil
}
