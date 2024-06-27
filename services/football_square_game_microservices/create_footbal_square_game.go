package footballsquaregameservices

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

type CreateFootballSquareGameService struct {
	GameID     int `json:"game_id"`
	SquareID   int `json:"square_id"`
	SquareSize int `json:"square_size"`

	Response CreateFootballSquareGameResponse
}

func (service CreateFootballSquareGameService) Request(config *util.Config) (CreateFootballSquareGameResponse, error) {
	createFootballSquareGameResponse := CreateFootballSquareGameResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	createFootballSquareGameURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/CreateFootballSquareGame"

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
