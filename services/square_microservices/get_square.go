package squareservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type GetSquare struct {
	SquareID     int    `json:"square_id"`
	SquareGUID   string `json:"square_guid"`
	SquareSize   int    `json:"square_size"`
	RowPoints    string `json:"row_points"`
	ColumnPoints string `json:"column_points"`
}

type GetSquareResponse struct {
	GetSquare
	ErrorMessage string `json:"error_message"`
}

type GetSquareService struct {
	SquareID int `json:"square_id"`
	Response GetSquareResponse
}

func (service GetSquareService) Request(config *util.Config) (GetSquareResponse, error) {
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
