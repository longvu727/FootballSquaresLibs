package footballsquaregameservices

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type ReserveFootballSquareResponse struct {
	Reserved     bool   `json:"reserved"`
	ErrorMessage string `json:"error_message"`
}

type ReserveFootballSquareService struct {
	GameGUID    int `json:"game_guid"`
	RowIndex    int `json:"row_index"`
	ColumnIndex int `json:"column_index"`

	Response ReserveFootballSquareResponse
}

func (service ReserveFootballSquareService) Request(config *util.Config) (ReserveFootballSquareResponse, error) {
	reserveFootballSquareGameResponse := ReserveFootballSquareResponse{}

	client := &http.Client{}
	serviceJson, _ := json.Marshal(service)

	reserveFootballSquareGameURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/ReserveFootballSquare"

	request, err := http.NewRequest("POST", reserveFootballSquareGameURL, bytes.NewBuffer(serviceJson))
	if err != nil {
		return reserveFootballSquareGameResponse, err
	}

	response, err := client.Do(request)
	if err != nil {
		return reserveFootballSquareGameResponse, err
	}

	if response.StatusCode != http.StatusOK {
		errStr := `unable to reserve FootballSquareGame`
		reserveFootballSquareGameResponse.ErrorMessage = errStr

		return reserveFootballSquareGameResponse, errors.New(errStr)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&reserveFootballSquareGameResponse)

	return reserveFootballSquareGameResponse, nil
}
