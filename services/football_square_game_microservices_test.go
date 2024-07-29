package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"math/rand"

	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/stretchr/testify/suite"
)

type FootballSquareGameTestSuite struct {
	suite.Suite
	config util.Config
}

func TestFootballSquareGameTestSuite(t *testing.T) {
	suite.Run(t, new(FootballSquareGameTestSuite))
}

func (suite *FootballSquareGameTestSuite) SetupServer(route string, mockResponse string) *httptest.Server {
	router := http.NewServeMux()
	router.HandleFunc(
		"POST "+route,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, string(mockResponse))
		}),
	)

	server := httptest.NewServer(router)

	return server
}

func (suite *FootballSquareGameTestSuite) TestCreateFootballSquareGame() {
	response := randomCreateFootballSquareGameResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/CreateFootballSquareGame", string(mockResponse))
	defer server.Close()

	suite.config = util.Config{
		MICROSERVICESBASEURL: map[string]string{
			"footballsquaregamemicroservices": server.URL,
		},
	}

	client := NewServices()

	actualResponse, err := client.CreateFootballSquareGame(&suite.config, CreateFootballSquareGameRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func (suite *FootballSquareGameTestSuite) TestGetFootballSquareGameByGameID() {
	response := randomGetFootballSquareGameByGameIDResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/GetFootballSquareGameByGameID", string(mockResponse))
	defer server.Close()

	suite.config = util.Config{
		MICROSERVICESBASEURL: map[string]string{
			"footballsquaregamemicroservices": server.URL,
		},
	}

	client := NewServices()

	actualResponse, err := client.GetFootballSquareGameByGameID(&suite.config, GetFootballSquareGameByGameIDRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func (suite *FootballSquareGameTestSuite) TestReserveFootballSquare() {
	response := randomReserveFootballSquareResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/ReserveFootballSquare", string(mockResponse))
	defer server.Close()

	suite.config = util.Config{
		MICROSERVICESBASEURL: map[string]string{
			"footballsquaregamemicroservices": server.URL,
		},
	}

	client := NewServices()

	actualResponse, err := client.ReserveFootballSquare(&suite.config, ReserveFootballSquareRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func randomReserveFootballSquareResponse() ReserveFootballSquareResponse{
	return ReserveFootballSquareResponse{
		Reserved:     true,
		ErrorMessage: "",
	}
}

func randomCreateFootballSquareGameResponse() CreateFootballSquareGameResponse {
	return CreateFootballSquareGameResponse{
		FootballSquaresGameIDs: []int64{rand.Int63n(1000), rand.Int63n(1000), rand.Int63n(1000)},
	}
}

func randomGetFootballSquareGameByGameIDResponse() GetFootballSquareGameByGameIDResponse {
	return GetFootballSquareGameByGameIDResponse{
		FootballSquares: []FootballSquareGameElement{
			{
				FootballSquaresGameID: rand.Intn(1000),
				ColumnIndex:           rand.Intn(10),
				RowIndex:              rand.Intn(10),
				WinnerQuaterNumber:    rand.Intn(4),
				Winner:                false,
				UserID:                rand.Intn(1000),
				SquareID:              rand.Intn(1000),
				GameID:                rand.Intn(1000),
				UserName:              "test name" + strconv.Itoa(rand.Intn(1000)),
				UserAlias:             "test alias" + strconv.Itoa(rand.Intn(1000)),
			},
		},
	}
}
