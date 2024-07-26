package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/longvu727/FootballSquaresLibs/util"
	"github.com/longvu727/FootballSquaresLibs/util/resources"
)

type Services interface {
	Post(serviceType string, resource *resources.Resources) error
	RequestCreateFootballSquareGame(config *util.Config) (CreateFootballSquareGameResponse, error)
	RequestGetFootballSquareGameByGameID(config *util.Config) (GetFootballSquareGameByGameIDResponse, error)
	RequestReserveFootballSquare(config *util.Config) (ReserveFootballSquareResponse, error)
}

type ServiceClient struct {
	Request  interface{}
	Endpoint string
}

func (service *ServiceClient) Post(response interface{}) error {

	serviceJson, _ := json.Marshal(service.Request)

	httpRequest, err := http.NewRequest("POST", service.Endpoint, bytes.NewBuffer(serviceJson))
	if err != nil {
		return err
	}

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return err
	}

	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("error while calling %s, StatusCode: %d", service.Endpoint, httpResponse.StatusCode)
	}

	defer httpResponse.Body.Close()

	if err := json.NewDecoder(httpResponse.Body).Decode(response); err != nil {
		return err
	}

	return nil
}
