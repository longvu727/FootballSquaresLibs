package footballsquaregameservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type GetFootballSquareGameResponse struct {
	FootballSquaresGameIDs []int64 `json:"football_square_game_ids"`
	ErrorMessage           string  `json:"error_message"`
}

type GetFootballSquareGameByGameID struct {
	GameID   int `json:"game_id"`
	Response GetFootballSquareGameResponse
}

func (service GetFootballSquareGameByGameID) Request(config *util.Config) (GetFootballSquareGameResponse, error) {
	getFootballSquareGameResponse := GetFootballSquareGameResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	getFootballSquareGameURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/GetFootballSquareGameByGameID"

	request, err := http.NewRequest("POST", getFootballSquareGameURL, bytes.NewBuffer(serviceJson))
	if err != nil {
		return getFootballSquareGameResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return getFootballSquareGameResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to create FootballSquareGame`
		getFootballSquareGameResponse.ErrorMessage = errStr

		return getFootballSquareGameResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&getFootballSquareGameResponse)

	return getFootballSquareGameResponse, nil
}
