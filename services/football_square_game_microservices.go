package services

import "github.com/longvu727/FootballSquaresLibs/util"

type CreateFootballSquareGameResponse struct {
	FootballSquaresGameIDs []int64 `json:"football_square_game_ids"`
	ErrorMessage           string  `json:"error_message"`
}

type CreateFootballSquareGameRequest struct {
	GameID     int `json:"game_id"`
	SquareID   int `json:"square_id"`
	SquareSize int `json:"square_size"`
}

type FootballSquareGameElement struct {
	FootballSquareGameID int  `json:"football_square_game_id"`
	ColumnIndex          int  `json:"column_index"`
	RowIndex             int  `json:"row_index"`
	WinnerQuaterNumber   int  `json:"winner_quater_number"`
	Winner               bool `json:"winner"`
	UserID               int  `json:"user_id"`
	SquareID             int  `json:"square_id"`
	GameID               int  `json:"game_id"`

	UserName  string `json:"user_name"`
	UserAlias string `json:"user_alias"`
}

type GetFootballSquareGameByGameIDResponse struct {
	FootballSquares []FootballSquareGameElement `json:"football_squares"`
	ErrorMessage    string                      `json:"error_message"`
}

type GetFootballSquareGameByGameIDRequest struct {
	GameID int `json:"game_id"`
}

type ReserveFootballSquareResponse struct {
	Reserved     bool   `json:"reserved"`
	ErrorMessage string `json:"error_message"`
}

type ReserveFootballSquareRequest struct {
	GameID      int `json:"game_id"`
	UserID      int `json:"user_id"`
	RowIndex    int `json:"row_index"`
	ColumnIndex int `json:"column_index"`
}

func (service *ServiceClient) CreateFootballSquareGame(config *util.Config, request CreateFootballSquareGameRequest) (CreateFootballSquareGameResponse, error) {
	createFootballSquareGameURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/CreateFootballSquareGame"
	response := CreateFootballSquareGameResponse{}

	err := service.Post(createFootballSquareGameURL, request, &response)

	return response, err
}

func (service *ServiceClient) GetFootballSquareGameByGameID(config *util.Config, request GetFootballSquareGameByGameIDRequest) (GetFootballSquareGameByGameIDResponse, error) {
	getFootballSquareGameByGameIDURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/GetFootballSquareGameByGameID"
	response := GetFootballSquareGameByGameIDResponse{}

	err := service.Post(getFootballSquareGameByGameIDURL, request, &response)

	return response, err
}

func (service *ServiceClient) ReserveFootballSquare(config *util.Config, request ReserveFootballSquareRequest) (ReserveFootballSquareResponse, error) {
	reserveFootballSquareGameURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/ReserveFootballSquare"
	reserveFootballSquareGameResponse := ReserveFootballSquareResponse{}

	err := service.Post(reserveFootballSquareGameURL, request, &reserveFootballSquareGameResponse)

	return reserveFootballSquareGameResponse, err
}
