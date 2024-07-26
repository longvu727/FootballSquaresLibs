package footballsquaregameservices

import (
	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util"
)

type GetFootballSquareGameByGameIDService interface {
	Request(config *util.Config) (GetFootballSquareGameByGameIDResponse, error)
}

type GetFootballSquareGameByGameIDResponse struct {
	FootballSquares []FootballSquareGameElement `json:"football_squares"`
	ErrorMessage    string                      `json:"error_message"`
}

type GetFootballSquareGameByGameID struct {
	GameID int `json:"game_id"`
}

func NewGetFootballSquareGameByGameIDService() GetFootballSquareGameByGameIDService {
	return &GetFootballSquareGameByGameID{}
}

func (service *GetFootballSquareGameByGameID) Request(config *util.Config) (GetFootballSquareGameByGameIDResponse, error) {
	getFootballSquareGameByGameIDURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/GetFootballSquareGameByGameID"
	getFootballSquareGameByGameIDResponse := GetFootballSquareGameByGameIDResponse{}

	client := services.ServiceClient{
		Request:  service,
		Endpoint: getFootballSquareGameByGameIDURL,
	}

	err := client.Post(&getFootballSquareGameByGameIDResponse)

	return getFootballSquareGameByGameIDResponse, err
}
