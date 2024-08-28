package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"math/rand"

	"github.com/google/uuid"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/stretchr/testify/suite"
)

type GameTestSuite struct {
	suite.Suite
	config util.Config
}

func TestGameTestSuite(t *testing.T) {
	suite.Run(t, new(GameTestSuite))
}

func (suite *GameTestSuite) getConfig(server *httptest.Server) util.Config {
	return util.Config{
		MICROSERVICESBASEURL: map[string]string{
			"gamemicroservices": server.URL,
		},
	}
}

func (suite *GameTestSuite) SetupServer(route string, mockResponse string) *httptest.Server {
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

func (suite *GameTestSuite) TestCreateGame() {
	response := randomCreateGameResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/CreateGame", string(mockResponse))
	defer server.Close()

	suite.config = suite.getConfig(server)

	client := NewServices()
	actualResponse, err := client.CreateGame(&suite.config, CreateGameRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func (suite *GameTestSuite) TestGetGameByGUID() {
	response := randomGetGameByGUIDResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/GetGameByGUID", string(mockResponse))
	defer server.Close()

	suite.config = suite.getConfig(server)

	client := NewServices()
	actualResponse, err := client.GetGameByGUID(&suite.config, GetGameByGUIDRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func (suite *GameTestSuite) TestGetGame() {
	response := randomGetGameResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/GetGame", string(mockResponse))
	defer server.Close()

	suite.config = suite.getConfig(server)

	client := NewServices()
	actualResponse, err := client.GetGame(&suite.config, GetGameRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func randomCreateGameResponse() CreateGameResponse {
	return CreateGameResponse{
		GameGUID: uuid.NewString(),
		GameID:   rand.Int63n(1000),
	}
}

func randomGetGameByGUIDResponse() GetGameByGUIDResponse {
	return GetGameByGUIDResponse{
		Game: randomGame(),
	}
}

func randomGetGameResponse() GetGameResponse{
	return GetGameResponse{
		Game: randomGame(),
	}
}

func randomGame() Game {
	return Game{
		GameGUID: uuid.NewString(),
		GameID:   rand.Int63n(1000),
		Sport:    "football",
		TeamA:    "Blue",
		TeamB:    "Red",
	}
}
