package gameservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type GetGameResponse struct {
	Game
	ErrorMessage string `json:"error_message"`
}

type GetGameService struct {
	GameID   int64 `json:"game_id"`
	Response GetGameResponse
}

func (service GetGameService) Request(config *util.Config) (GetGameResponse, error) {
	getGameResponse := GetGameResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	getGameUrl := config.MICROSERVICESBASEURL["gamemicroservices"] + "/GetGame"

	request, err := http.NewRequest("POST", getGameUrl, bytes.NewBuffer(serviceJson))
	if err != nil {
		return getGameResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return getGameResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to get Game`
		getGameResponse.ErrorMessage = errStr

		return getGameResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&getGameResponse)

	return getGameResponse, nil
}
