package services

import "github.com/longvu727/FootballSquaresLibs/util"

type CreateGameResponse struct {
	GameGUID     string `json:"game_guid"`
	GameID       int64  `json:"game_id"`
	ErrorMessage string `json:"error_message"`
}

type CreateGameRequest struct {
	Sport      string `json:"sport"`
	SquareSize int32  `json:"square_size"`
	TeamA      string `json:"team_a"`
	TeamB      string `json:"team_b"`
}

type Game struct {
	GameGUID string `json:"game_guid"`
	GameID   int64  `json:"game_id"`
	Sport    string `json:"sport"`
	TeamA    string `json:"team_a"`
	TeamB    string `json:"team_b"`
}

type GetGameByGUIDResponse struct {
	Game
	ErrorMessage string `json:"error_message"`
}

type GetGameByGUIDRequest struct {
	GameGUID string `json:"game_guid"`
}

type GetGameResponse struct {
	Game
	ErrorMessage string `json:"error_message"`
}

type GetGameRequest struct {
	GameID   int64 `json:"game_id"`
	Response GetGameResponse
}

func (service ServiceClient) CreateGame(config *util.Config, request CreateGameRequest) (CreateGameResponse, error) {
	createGameUrl := config.MICROSERVICESBASEURL["gamemicroservices"] + "/CreateGame"
	response := CreateGameResponse{}

	err := service.Post(createGameUrl, request, &response)

	return response, err
}

func (service ServiceClient) GetGameByGUID(config *util.Config, request GetGameByGUIDRequest) (GetGameByGUIDResponse, error) {
	getGameByGUIDUrl := config.MICROSERVICESBASEURL["gamemicroservices"] + "/GetGameByGUID"
	response := GetGameByGUIDResponse{}

	err := service.Post(getGameByGUIDUrl, request, &response)

	return response, err
}

func (service ServiceClient) GetGame(config *util.Config, request GetGameRequest) (GetGameResponse, error) {
	getGameUrl := config.MICROSERVICESBASEURL["gamemicroservices"] + "/GetGame"
	response := GetGameResponse{}

	err := service.Post(getGameUrl, request, &response)

	return response, err

}
