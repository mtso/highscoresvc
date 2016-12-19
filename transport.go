// transport.go
// mtso 2016

package highscoresvc

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

func DecodePostScoreRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request postScoreRequest
	if error := json.NewDecoder(r.Body).Decode(&request); error != nil {
		return nil, error
	}
	return request, nil
}

func DecodeGetScoreRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getScoreRequest
	if error := json.NewDecoder(r.Body).Decode(&request); error != nil {
		return nil, error
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
