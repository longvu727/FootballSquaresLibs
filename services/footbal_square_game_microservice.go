package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type CreateFootballSquareGameResponse struct {
	FootballSquaresGameIDs []int64 `json:"football_square_game_ids"`
	ErrorMessage           string  `json:"error_message"`
}

type CreateFootballSquareGame struct {
	GameID     int32 `json:"game_id"`
	SquareID   int32 `json:"square_id"`
	SquareSize int32 `json:"square_size"`

	Response CreateFootballSquareGameResponse
}

func (service CreateFootballSquareGame) Request() (CreateFootballSquareGameResponse, error) {
	createFootballSquareGameResponse := CreateFootballSquareGameResponse{}
	microServicesConfig, _ := util.LoadConfig(".", "microservices", "json")

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	createFootballSquareGameURL := microServicesConfig.MICROSERVICES["footballsquaregamemicroservices"].BaseUrl + "/CreateFootballSquareGame"

	request, err := http.NewRequest("POST", createFootballSquareGameURL, bytes.NewBuffer(serviceJson))
	if err != nil {
		return createFootballSquareGameResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return createFootballSquareGameResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to create FootballSquareGame`
		createFootballSquareGameResponse.ErrorMessage = errStr

		return createFootballSquareGameResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&createFootballSquareGameResponse)

	return createFootballSquareGameResponse, nil
}
