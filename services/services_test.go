package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServicesTestSuite struct {
	suite.Suite
}

func (suite *ServicesTestSuite) TestPostInvalidEndPoint() {
	router := http.NewServeMux()
	router.HandleFunc("POST /test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"api_response_var1": 1, "api_response_var2": 2}`)
	}))

	server := httptest.NewServer(router)
	defer server.Close()

	client := ServiceClient{
		Request:  struct{}{},
		Endpoint: string([]byte{0x7f}),
	}

	err := client.Post(&struct{}{})
	fmt.Println(err.Error(), server.URL)

	suite.Error(err)
}

func (suite *ServicesTestSuite) TestPostHTTPError() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "test error", http.StatusBadRequest)
	}))
	defer server.Close()

	client := ServiceClient{
		Request:  struct{}{},
		Endpoint: server.URL,
	}

	err := client.Post(&struct{}{})
	suite.Error(err)
}

func (suite *ServicesTestSuite) TestPostDoError() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"api_response_var1": 1, "api_response_var2": 2}`)
	}))
	defer server.Close()

	client := ServiceClient{
		Request:  struct{}{},
		Endpoint: "dakfjakljal",
	}

	err := client.Post(&struct{}{})
	suite.Error(err)
}

func (suite *ServicesTestSuite) TestPostJsonDecodeError() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"api_response_var1": 1, "api_response_var2": 2`)
	}))
	defer server.Close()

	client := ServiceClient{
		Request:  struct{}{},
		Endpoint: server.URL,
	}

	err := client.Post(&struct{}{})
	suite.Error(err)
}

func (suite *ServicesTestSuite) TestPost() {
	type testRequest struct {
		RequestParameter1 int `json:"request_param_var1"`
		RequestParameter2 int `json:"request_param_var2"`
	}

	type testResponse struct {
		APIResponseVar1 int `json:"api_response_var1"`
		APIResponseVar2 int `json:"api_response_var2"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"api_response_var1": 1, "api_response_var2": 2}`)
	}))
	defer server.Close()

	client := ServiceClient{
		Request: testRequest{
			RequestParameter1: 1,
			RequestParameter2: 2,
		},
		Endpoint: server.URL,
	}

	myResponse := &testResponse{}

	err := client.Post(myResponse)

	suite.NoError(err)
	suite.Equal(1, myResponse.APIResponseVar1)
	suite.Equal(2, myResponse.APIResponseVar2)
}

func TestServicesTestSuite(t *testing.T) {
	suite.Run(t, new(ServicesTestSuite))
}
