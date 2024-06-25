package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type CreateSquareResponse struct {
	GameGUID     string `json:"game_guid"`
	ErrorMessage string `json:"error_message"`
}

type CreateSquare struct {
	SquareSize int    `json:"square_size"`
	Sport      string `json:"sport"`
	Response   CreateSquareResponse
}

func (service CreateSquare) Request() (CreateSquareResponse, error) {
	createSquareResponse := CreateSquareResponse{}
	microServicesConfig, _ := util.LoadConfig(".", "microservices", "json")

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	createSquareUrl := microServicesConfig.MICROSERVICES["squaremicroservices"].BaseUrl + "/CreateSquare"

	request, err := http.NewRequest("POST", createSquareUrl, bytes.NewBuffer(serviceJson))
	if err != nil {
		return createSquareResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return createSquareResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to create square`
		createSquareResponse.ErrorMessage = errStr

		return createSquareResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&createSquareResponse)

	return createSquareResponse, nil
}
