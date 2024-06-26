package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type CreateGameResponse struct {
	GameGUID     string `json:"game_guid"`
	GameID       int64  `json:"game_id"`
	ErrorMessage string `json:"error_message"`
}

type CreateGame struct {
	Sport      string `json:"sport"`
	SquareSize int32  `json:"square_size"`
	TeamA      string `json:"team_a"`
	TeamB      string `json:"team_b"`

	Response CreateGameResponse
}

func (service CreateGame) Request() (CreateGameResponse, error) {
	createGameResponse := CreateGameResponse{}
	microServicesConfig, _ := util.LoadConfig("./", "microservices", "json")

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	createGameUrl := microServicesConfig.MICROSERVICES["gamemicroservices"].BaseUrl + "/CreateGame"

	request, err := http.NewRequest("POST", createGameUrl, bytes.NewBuffer(serviceJson))
	if err != nil {
		return createGameResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return createGameResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to create Game`
		createGameResponse.ErrorMessage = errStr

		return createGameResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&createGameResponse)

	return createGameResponse, nil
}
