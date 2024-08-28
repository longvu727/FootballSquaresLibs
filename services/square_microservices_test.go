package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"math/rand"

	"github.com/google/uuid"
	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/stretchr/testify/suite"
)

type SquareTestSuite struct {
	suite.Suite
	config util.Config
}

func TestSquareTestSuite(t *testing.T) {
	suite.Run(t, new(SquareTestSuite))
}

func (suite *SquareTestSuite) getConfig(server *httptest.Server) util.Config {
	return util.Config{
		MICROSERVICESBASEURL: map[string]string{
			"squaremicroservices": server.URL,
		},
	}
}

func (suite *SquareTestSuite) SetupServer(route string, mockResponse string) *httptest.Server {
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

func (suite *SquareTestSuite) TestCreateSquare() {
	response := randomCreateSquareResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/CreateSquare", string(mockResponse))
	defer server.Close()

	suite.config = suite.getConfig(server)

	client := NewServices()
	actualResponse, err := client.CreateSquare(&suite.config, CreateSquareRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func (suite *SquareTestSuite) TestGetSquare() {
	response := randomGetSquareResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/GetSquare", string(mockResponse))
	defer server.Close()

	suite.config = suite.getConfig(server)

	client := NewServices()
	actualResponse, err := client.GetSquare(&suite.config, GetSquareRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func randomCreateSquareResponse() CreateSquareResponse {
	return CreateSquareResponse{
		SquareID: rand.Intn(1000),
	}
}

func randomGetSquareResponse() GetSquareResponse {
	return GetSquareResponse{
		Square: randomSquare(),
	}
}

func randomSquare() Square {
	rowpoints := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	colpoints := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	rand.Shuffle(len(rowpoints), func(i, j int) {
		rowpoints[i], rowpoints[j] = rowpoints[j], rowpoints[i]
	})

	rand.Shuffle(len(colpoints), func(i, j int) {
		colpoints[i], colpoints[j] = colpoints[j], colpoints[i]
	})

	return Square{
		SquareGUID:   uuid.NewString(),
		SquareID:     rand.Intn(1000),
		SquareSize:   rand.Intn(10),
		RowPoints:    strings.Join(rowpoints, ","),
		ColumnPoints: strings.Join(rowpoints, ","),
	}
}
