package footballsquaregameservices

import (
	"github.com/longvu727/FootballSquaresLibs/services"
	"github.com/longvu727/FootballSquaresLibs/util"
)

type CreateFootballSquareGameService interface {
	Request(config *util.Config) (CreateFootballSquareGameResponse, error)
}

type CreateFootballSquareGameResponse struct {
	FootballSquaresGameIDs []int64 `json:"football_square_game_ids"`
	ErrorMessage           string  `json:"error_message"`
}

type CreateFootballSquareGame struct {
	GameID     int `json:"game_id"`
	SquareID   int `json:"square_id"`
	SquareSize int `json:"square_size"`
}

func NewCreateFootballSquareGameService() CreateFootballSquareGameService {
	return &CreateFootballSquareGame{}
}

func (service *CreateFootballSquareGame) Request(config *util.Config) (CreateFootballSquareGameResponse, error) {
	createFootballSquareGameResponse := CreateFootballSquareGameResponse{}
	createFootballSquareGameURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/CreateFootballSquareGame"

	client := services.ServiceClient{
		Request:  service,
		Endpoint: createFootballSquareGameURL,
	}

	err := client.Post(&createFootballSquareGameResponse)

	return createFootballSquareGameResponse, err
}
