package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
)

type Services interface {
	Post(enpoint string, request interface{}, response interface{}) error
	CreateFootballSquareGame(config *util.Config, request CreateFootballSquareGameRequest) (CreateFootballSquareGameResponse, error)
	GetFootballSquareGameByGameID(config *util.Config, request GetFootballSquareGameByGameIDRequest) (GetFootballSquareGameByGameIDResponse, error)
	ReserveFootballSquare(config *util.Config, request ReserveFootballSquareRequest) (ReserveFootballSquareResponse, error)

	CreateUser(config *util.Config, request CreateUserRequest) (CreateUserResponse, error)
	GetUserByGUID(config *util.Config, request GetUserByGUIDRequest) (GetUserByGUIDResponse, error)
	GetUser(config *util.Config, request GetUserRequest) (GetUserResponse, error)

	CreateSquare(config *util.Config, request CreateSquareRequest) (CreateSquareResponse, error)
	GetSquare(config *util.Config, request GetSquareRequest) (GetSquareResponse, error)

	CreateGame(config *util.Config, request CreateGameRequest) (CreateGameResponse, error)
	GetGameByGUID(config *util.Config, request GetGameByGUIDRequest) (GetGameByGUIDResponse, error)
	GetGame(config *util.Config, request GetGameRequest) (GetGameResponse, error)
}

type ServiceClient struct {
	Client *http.Client
}

func NewServices() Services {
	return &ServiceClient{
		Client: &http.Client{},
	}
}

func (service *ServiceClient) Post(enpoint string, request interface{}, response interface{}) error {

	serviceJson, _ := json.Marshal(request)

	httpRequest, err := http.NewRequest("POST", enpoint, bytes.NewBuffer(serviceJson))
	if err != nil {
		return err
	}

	httpResponse, err := service.Client.Do(httpRequest)
	if err != nil {
		return err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("error while calling %s, StatusCode: %d", enpoint, httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()

	if err := json.NewDecoder(httpResponse.Body).Decode(response); err != nil {
		return err
	}

	return nil
}
