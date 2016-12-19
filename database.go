package highscoresvc

import (
	"database/sql"

	"golang.org/x/net/context"
)

const (
	getQuery    = "SELECT * FROM Highscore WHERE username='$1'"
	updateQuery = "UPDATE Highscore SET $2 = 40 WHERE username='$1' AND value < $2"

	createTable = `CREATE TABLE IF NOT EXISTS Highscore (
		username varchar(64) PRIMARY KEY,
		value INT NOT NULL
	);`
)

type dbService struct {
	db *sql.DB
}

func NewPostgreService(location string) Service {
	if db, err := InitializeDatabase("postgres", location); err != nil {
		return nil
	} else {
		return &dbService{db}
	}
}

func InitializeDatabase(databaseType, databaseUrl string) (*sql.DB, error) {
	// Open database type with url
	db, err := sql.Open(databaseType, databaseUrl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// Create `Highscore` table if it doesn't exist
	if _, error := db.Exec(createTable); error != nil {
		return nil, error
	}

	return db, nil
}

func (svc *dbService) PostScore(ctx context.Context, h Highscore) (*Highscore, error) {

	if _, error := svc.db.Exec(updateQuery, h.Username, h.Value); error != nil {
		return nil, error
	}

	row := svc.db.QueryRow(getQuery, h.Username)
	result := new(Highscore)
	if error := row.Scan(&result.Username, &result.Value); error != nil {
		return nil, error
	}

	return &Highscore{h.Username, h.Value}, nil
}

func (svc *dbService) GetScore(ctx context.Context, username string) (*Highscore, error) {

	row := svc.db.QueryRow(getQuery, username)
	result := new(Highscore)
	if error := row.Scan(&result.Username, &result.Value); error != nil {
		return nil, error
	}

	return result, nil
}
