package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type GetSquareResponse struct {
	SquareID     int    `json:"square_id"`
	ErrorMessage string `json:"error_message"`
}

type GetSquare struct {
	SquareSize int    `json:"square_size"`
	Sport      string `json:"sport"`
	Response   GetSquareResponse
}

func (service GetSquare) Request(config *util.Config) (GetSquareResponse, error) {
	getSquareResponse := GetSquareResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	getSquareUrl := config.MICROSERVICESBASEURL["squaremicroservices"] + "/GetSquare"

	request, err := http.NewRequest("POST", getSquareUrl, bytes.NewBuffer(serviceJson))
	if err != nil {
		return getSquareResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return getSquareResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to get square`
		getSquareResponse.ErrorMessage = errStr

		return getSquareResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&getSquareResponse)

	return getSquareResponse, nil
}
