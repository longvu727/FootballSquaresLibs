package gameservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type GetGameByGUIDResponse struct {
	Game
	ErrorMessage  string `json:"error_message"`
}

type GetGameByGUIDService struct {
	GameGUID string `json:"game_guid"`
}

func (service GetGameByGUIDService) Request(config *util.Config) (GetGameByGUIDResponse, error) {
	getGameByGUIDResponse := GetGameByGUIDResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	getGameByGUIDUrl := config.MICROSERVICESBASEURL["gamemicroservices"] + "/GetGameByGUID"

	request, err := http.NewRequest("POST", getGameByGUIDUrl, bytes.NewBuffer(serviceJson))
	if err != nil {
		return getGameByGUIDResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return getGameByGUIDResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to get Game by GUID`
		getGameByGUIDResponse.ErrorMessage = errStr

		return getGameByGUIDResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&getGameByGUIDResponse)

	return getGameByGUIDResponse, nil
}
