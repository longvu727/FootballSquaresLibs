package services

import "github.com/longvu727/FootballSquaresLibs/util"

type CreateFootballSquareGameResponse struct {
	FootballSquaresGameIDs []int64 `json:"football_square_game_ids"`
	ErrorMessage           string  `json:"error_message"`
}

type CreateFootballSquareGame struct {
	GameID     int `json:"game_id"`
	SquareID   int `json:"square_id"`
	SquareSize int `json:"square_size"`
}

type FootballSquareGameElement struct {
	FootballSquaresGameID int  `json:"football_square_game_id"`
	ColumnIndex           int  `json:"column_index"`
	RowIndex              int  `json:"row_index"`
	WinnerQuaterNumber    int  `json:"winner_quater_number"`
	Winner                bool `json:"winner"`
	UserID                int  `json:"user_id"`
	SquareID              int  `json:"square_id"`
	GameID                int  `json:"game_id"`

	UserName  string `json:"user_name"`
	UserAlias string `json:"user_alias"`
}

type GetFootballSquareGameByGameIDResponse struct {
	FootballSquares []FootballSquareGameElement `json:"football_squares"`
	ErrorMessage    string                      `json:"error_message"`
}

type GetFootballSquareGameByGameID struct {
	GameID int `json:"game_id"`
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

func (service *ServiceClient) RequestCreateFootballSquareGame(config *util.Config) (CreateFootballSquareGameResponse, error) {
	createFootballSquareGameResponse := CreateFootballSquareGameResponse{}
	createFootballSquareGameURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/CreateFootballSquareGame"

	client := ServiceClient{
		Request:  service,
		Endpoint: createFootballSquareGameURL,
	}

	err := client.Post(&createFootballSquareGameResponse)

	return createFootballSquareGameResponse, err
}

func (service *ServiceClient) RequestGetFootballSquareGameByGameID(config *util.Config) (GetFootballSquareGameByGameIDResponse, error) {
	getFootballSquareGameByGameIDURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/GetFootballSquareGameByGameID"
	getFootballSquareGameByGameIDResponse := GetFootballSquareGameByGameIDResponse{}

	client := ServiceClient{
		Request:  service,
		Endpoint: getFootballSquareGameByGameIDURL,
	}

	err := client.Post(&getFootballSquareGameByGameIDResponse)

	return getFootballSquareGameByGameIDResponse, err
}

func (service *ServiceClient) RequestReserveFootballSquare(config *util.Config) (ReserveFootballSquareResponse, error) {
	reserveFootballSquareGameResponse := ReserveFootballSquareResponse{}
	reserveFootballSquareGameURL := config.MICROSERVICESBASEURL["footballsquaregamemicroservices"] + "/ReserveFootballSquare"

	client := ServiceClient{
		Request:  service,
		Endpoint: reserveFootballSquareGameURL,
	}

	err := client.Post(&reserveFootballSquareGameResponse)

	return reserveFootballSquareGameResponse, err
}
