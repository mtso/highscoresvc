package main

import (
	"log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/mtso/highscoresvc"
	"golang.org/x/net/context"
)

const (
	local_db   = "user=kingcandy password=cupcakes dbname=highscoresvc sslmode=disable"
	local_port = "3000"
)

func main() {
	ctx := context.Background()

	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		db_url = local_db
	}

	// Business domain
	svc := highscoresvc.NewPostgreService(db_url)

	// Handles
	// Endpoint domain
	postScoreHandler := httptransport.NewServer(
		ctx,
		highscoresvc.MakePostScoreEndpoint(svc),
		highscoresvc.DecodePostScoreRequest,
		highscoresvc.EncodeResponse,
	)

	getScoreHandler := httptransport.NewServer(
		ctx,
		highscoresvc.MakeGetScoreEndpoint(svc),
		highscoresvc.DecodeGetScoreRequest,
		highscoresvc.EncodeResponse,
	)

	// Determine port
	port := os.Getenv("PORT")
	if port == "" {
		port = local_port // If port was not set as an environment variable
	}

	// Routes
	http.Handle("/post", postScoreHandler)
	http.Handle("/get", getScoreHandler)

	//	errc := make(chan error)
	//	go func() {
	//		errc <- http.ListenAndServe(":"+local_port, nil)
	//	}()

	//	// HTTP transport
	//	go func() {
	//		logger:=log.NewContext(logger).With("transport", "HTTP")
	//		h := highscoresvc.makeHTTPHandler(
	//			ctx,
	//			endpoints,
	//			tracer,
	//			logger
	//		)
	//		logger.Log("addr", *httpAddr)
	//		error <- http.ListenAndServe(*httpAddr, h)
	//	}

	// Run
	log.Println("Listening on", port)
	log.Println(http.ListenAndServe(":"+local_port, nil))
	//	logger.Log("exit", <-error)
}
