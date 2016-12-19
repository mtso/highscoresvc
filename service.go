// mtso copyright 2016
// Defines a highscore service interface

package highscoresvc

import (
	"golang.org/x/net/context"
)

type Service interface {
	PostScore(ctx context.Context, h Highscore) (*Highscore, error)
	GetScore(ctx context.Context, username string) (*Highscore, error)
}

type Highscore struct {
	Username string
	Value    int
}
