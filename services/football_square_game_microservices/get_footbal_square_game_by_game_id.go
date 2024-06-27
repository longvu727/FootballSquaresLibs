package footballsquaregameservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type GetFootballSquareGameByGameIDResponse struct {
	FootballSquares []FootballSquareGameElement `json:"football_squares"`
	ErrorMessage    string                      `json:"error_message"`
}

type GetFootballSquareGameByGameIDService struct {
	GameID   int `json:"game_id"`
	Response GetFootballSquareGameByGameIDResponse
}

func (service GetFootballSquareGameByGameIDService) Request(config *util.Config) (GetFootballSquareGameByGameIDResponse, error) {
	getFootballSquareGameByGameIDResponse := GetFootballSquareGameByGameIDResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	getFootballSquareGameByGameIDURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/GetFootballSquareGameByGameID"

	request, err := http.NewRequest("POST", getFootballSquareGameByGameIDURL, bytes.NewBuffer(serviceJson))
	if err != nil {
		return getFootballSquareGameByGameIDResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return getFootballSquareGameByGameIDResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to create FootballSquareGame`
		getFootballSquareGameByGameIDResponse.ErrorMessage = errStr

		return getFootballSquareGameByGameIDResponse, errors.New(errStr)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&getFootballSquareGameByGameIDResponse)
	if err != nil {
		log.Print(err.Error())
	}

	return getFootballSquareGameByGameIDResponse, nil
}
