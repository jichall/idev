package main

import (
	"github.com/rafaelcn/jeolgyu"

	"github.com/jichall/idev/src/parser"
)

var (
	logger *jeolgyu.Jeolgyu

	inputFile = flag.String("input", "assets/data.json", "")

	host = flag.String("host", "localhost", "Where to serve the application")
	port = flag.Int("port", "8080", "What port to serve the application")

	isProduction = flag.Bool("production", false, "Is the app running in production environment?")
)

func main() {
	var st = jeolgyu.SinkBoth

	if *production {
		st = jeolgyu.SinkFile
	}

	logger, err := jeolgyu.New(jeolgyu.Settings{
		Filepath: "log",
		Filename: "",
		SinkType: st,
	})

	if err != nil {
		log.Fatalf("[!] Couldn't initialize custom logger. Reason %v", err)
	}

	// initialize data input processing
	collection, err := parser.Parse(*inputFile)

	if err != nil {
		logger.Error("Failed to process input file. Reason %v", err)
		os.Exit(1)
	}

	// initialize HTTP server
	serve(*host, *port)
}