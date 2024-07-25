package services

import "github.com/longvu727/FootballSquaresLibs/util"

type Services interface {
	Request(config *util.Config) (Response, error)
}

type Response interface{}
