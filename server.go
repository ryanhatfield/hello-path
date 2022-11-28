package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexflint/go-arg"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var args struct {
	Port     int    `arg:"-p,--port,env:HELLO_PATH_PORT" default:"8081"`
	PathRoot string `arg:"positional,required"`
}

var logger *zap.Logger
var sugar *zap.SugaredLogger

func setupArgs() {
	args.Port = 8081
	arg.MustParse(&args)
}

func setupLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar = logger.Sugar()
}

func setupRouter() {
	rtr := mux.NewRouter()

	path := func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		thePath := params["path"]
		message := fmt.Sprintf("Path: %s", thePath)
		w.Write([]byte(message))
		sugar.With("path", thePath)
		sugar.Info(message)
	}

	rtr.HandleFunc(fmt.Sprintf("%s{path:.*}", args.PathRoot), path).Methods("GET")

	http.Handle("/", rtr)

	sugar.Infof("Args: %+v", args)
}

func serve() {
	err := http.ListenAndServe(fmt.Sprintf(":%d", args.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	setupArgs()
	setupLogger()
	setupRouter()

	serve()
}
