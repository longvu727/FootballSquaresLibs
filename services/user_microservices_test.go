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

type UserTestSuite struct {
	suite.Suite
	config util.Config
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (suite *UserTestSuite) getConfig(server *httptest.Server) util.Config {
	return util.Config{
		MICROSERVICESBASEURL: map[string]string{
			"usermicroservices": server.URL,
		},
	}
}

func (suite *UserTestSuite) SetupServer(route string, mockResponse string) *httptest.Server {
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

func (suite *UserTestSuite) TestCreateUser() {
	response := randomCreateUserResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/CreateUser", string(mockResponse))
	defer server.Close()

	suite.config = suite.getConfig(server)

	client := NewServices()
	actualResponse, err := client.CreateUser(&suite.config, CreateUserRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func (suite *UserTestSuite) TestGetUser() {
	response := randomGetUserResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/GetUser", string(mockResponse))
	defer server.Close()

	suite.config = suite.getConfig(server)

	client := NewServices()
	actualResponse, err := client.GetUser(&suite.config, GetUserRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func (suite *UserTestSuite) TestGetUserByGUID() {
	response := randomGetUserByGUIDResponse()
	mockResponse, _ := json.Marshal(response)

	server := suite.SetupServer("/GetUserByGUID", string(mockResponse))
	defer server.Close()

	suite.config = suite.getConfig(server)

	client := NewServices()
	actualResponse, err := client.GetUserByGUID(&suite.config, GetUserByGUIDRequest{})

	suite.NoError(err)
	suite.Equal(response, actualResponse)
}

func randomCreateUserResponse() CreateUserResponse {
	return CreateUserResponse{
		UserID: rand.Intn(1000),
	}
}

func randomGetUserResponse() GetUserResponse {
	return GetUserResponse{
		User: randomUser(),
	}
}

func randomGetUserByGUIDResponse() GetUserByGUIDResponse {
	return GetUserByGUIDResponse{
		User: randomUser(),
	}
}

func randomUser() User {
	return User{
		UserGUID:   uuid.NewString(),
		UserID:     rand.Intn(1000),
		IP:         "127.0.0.1",
		DeviceName: "Chrome",
		UserName:   "Long Vu",
		Alias:      "lvu",
	}
}
