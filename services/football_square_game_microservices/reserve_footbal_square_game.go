package footballsquaregameservices

import (
	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util"
)

type ReserveFootballSquareService interface {
	Request(config *util.Config) (ReserveFootballSquareResponse, error)
}

type ReserveFootballSquareResponse struct {
	Reserved     bool   `json:"reserved"`
	ErrorMessage string `json:"error_message"`
}

type ReserveFootballSquare struct {
	GameID      int `json:"game_id"`
	UserID      int `json:"user_id"`
	RowIndex    int `json:"row_index"`
	ColumnIndex int `json:"column_index"`
}

func NewReserveFootballSquareService() ReserveFootballSquareService {
	return &ReserveFootballSquare{}
}

func (service *ReserveFootballSquare) Request(config *util.Config) (ReserveFootballSquareResponse, error) {
	reserveFootballSquareGameResponse := ReserveFootballSquareResponse{}
	reserveFootballSquareGameURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/ReserveFootballSquare"

	client := services.ServiceClient{
		Request:  service,
		Endpoint: reserveFootballSquareGameURL,
	}

	err := client.Post(&reserveFootballSquareGameResponse)

	return reserveFootballSquareGameResponse, err
}
