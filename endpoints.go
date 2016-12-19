package highscoresvc

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

func MakePostScoreEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postScoreRequest)
		convert := Highscore{
			Username: req.Username,
			Value:    req.Score,
		}
		highscore, err := svc.PostScore(ctx, convert)
		if err != nil {
			return nil, err
		}
		return postScoreResponse{
			Username:  highscore.Username,
			Highscore: highscore.Value,
		}, nil
	}
}

func MakeGetScoreEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getScoreRequest)
		highscore, err := svc.GetScore(ctx, req.Username)
		if err != nil {
			return nil, err
		}
		return getScoreResponse{
			Username:  highscore.Username,
			Highscore: highscore.Value,
		}, nil
	}
}

type postScoreRequest struct {
	Username string
	Score    int
}

type postScoreResponse struct {
	Username  string `json:'username'`
	Highscore int    `json:'highscore'`
}

type getScoreRequest struct {
	Username string
}

type getScoreResponse struct {
	Username  string `json:'username'`
	Highscore int    `json:'highscore'`
}

/*
POST {"username":"name","score":40}
  -> {"username":"name","highscore":40} // Returns the highscore so far

GET  {"username":"name"}
  -> {"username":"name","highscore":40} // Returns the name too
  -> {"username":"name","highscore":0}  // If the user does not exist
*/
