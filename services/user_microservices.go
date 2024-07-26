package services

import "github.com/longvu727/FootballSquaresLibs/util"

type CreateUserRequest struct {
	IP         string `json:"ip"`
	DeviceName string `json:"device_name"`
	UserName   string `json:"user_name"`
	Alias      string `json:"alias"`
}

type CreateUserResponse struct {
	UserID       int    `json:"user_id"`
	ErrorMessage string `json:"error_message"`
}

type User struct {
	UserID     int    `json:"user_id"`
	UserGUID   string `json:"user_guid"`
	IP         string `json:"ip"`
	DeviceName string `json:"device_name"`
	UserName   string `json:"user_name"`
	Alias      string `json:"alias"`
}

type GetUserByGUIDResponse struct {
	User
	ErrorMessage string `json:"error_message"`
}

type GetUserByGUIDRequest struct {
	UserGUID string `json:"user_guid"`
}

type GetUserResponse struct {
	User
	ErrorMessage string `json:"error_message"`
}

type GetUserRequest struct {
	UserID int `json:"user_id"`
}

func (service ServiceClient) CreateUser(config *util.Config, request CreateUserRequest) (CreateUserResponse, error) {
	createUserUrl := config.MICROSERVICESBASEURL["usermicroservices"] + "/CreateUser"
	response := CreateUserResponse{}

	err := service.Post(createUserUrl, request, response)

	return response, err
}

func (service ServiceClient) GetUserByGUID(config *util.Config, request GetUserByGUIDRequest) (GetUserByGUIDResponse, error) {
	getUserByGUIDUrl := config.MICROSERVICESBASEURL["usermicroservices"] + "/GetUserByGUID"
	response := GetUserByGUIDResponse{}

	err := service.Post(getUserByGUIDUrl, request, response)

	return response, err
}

func (service ServiceClient) GetUser(config *util.Config, request GetUserRequest) (GetUserResponse, error) {
	getUserUrl := config.MICROSERVICESBASEURL["usermicroservices"] + "/GetUser"
	response := GetUserResponse{}

	err := service.Post(getUserUrl, request, response)

	return response, err
}
