package main

import (
	"flag"
	"log"
	"os"

	"github.com/rafaelcn/jeolgyu"

	"github.com/jichall/idev/src/parser"
)

var (
	logger *jeolgyu.Jeolgyu

	collection parser.ServerCollection

	folder = flag.String("folder", "./data", "The folder to where look for .json files")
	host = flag.String("host", "localhost", "Where to serve the application")
	port = flag.String("port", "8080", "What port to serve the application")
	production = flag.Bool("production", false, "Is the app running in production environment?")
)

func main() {
	flag.Parse()

	var st = jeolgyu.SinkBoth

	if *production {
		st = jeolgyu.SinkFile
	}

	var err error
	logger, err = jeolgyu.New(jeolgyu.Settings{
		Filepath: "log",
		Filename: "",
		SinkType: st,
	})

	if err != nil {
		log.Fatalf("[!] Couldn't initialize custom logger. Reason %v", err)
	}

	// initialize data input processing
	collection, err = parser.Parse(*folder)

	if err != nil {
		logger.Error("Failed to process input file. Reason %v", err)
		os.Exit(1)
	}

	// process calculates each server usage metric
	process(collection)

	// initialize HTTP server
	serve(*host, *port)
}
