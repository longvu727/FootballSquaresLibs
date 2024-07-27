package services

import "github.com/longvu727/FootballSquaresLibs/util"

type CreateSquareResponse struct {
	SquareID     int    `json:"square_id"`
	ErrorMessage string `json:"error_message"`
}

type CreateSquareRequest struct {
	SquareSize int    `json:"square_size"`
	Sport      string `json:"sport"`
}

type Square struct {
	SquareID     int    `json:"square_id"`
	SquareGUID   string `json:"square_guid"`
	SquareSize   int    `json:"square_size"`
	RowPoints    string `json:"row_points"`
	ColumnPoints string `json:"column_points"`
}

type GetSquareResponse struct {
	Square
	ErrorMessage string `json:"error_message"`
}

type GetSquareRequest struct {
	SquareID int `json:"square_id"`
}

func (service ServiceClient) CreateSquare(config *util.Config, request CreateSquareRequest) (CreateSquareResponse, error) {
	createSquareUrl := config.MICROSERVICESBASEURL["squaremicroservices"] + "/CreateSquare"
	response := CreateSquareResponse{}

	err := service.Post(createSquareUrl, request, &response)

	return response, err

}

func (service ServiceClient) GetSquare(config *util.Config, request GetSquareRequest) (GetSquareResponse, error) {
	createSquareUrl := config.MICROSERVICESBASEURL["squaremicroservices"] + "/GetSquare"
	response := GetSquareResponse{}

	err := service.Post(createSquareUrl, request, &response)

	return response, err
}
