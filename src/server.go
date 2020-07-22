package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func serve(host, port string) {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/servers/raw", HandleRaw).Methods("GET")
	r.HandleFunc("/servers/raw/{server}", HandleRaw).Methods("GET")

	r.HandleFunc("/servers", HandleAll).Methods("GET")
	r.HandleFunc("/servers/{server}", HandleServer).Methods("GET")

	// route created to satisfy the prerequisites of the challenge as it was
	// unclear on the README description.
	r.HandleFunc("/spec", HandleSpecific)

	s := http.Server{
		Addr: host + ":" + port,
		//Handler: handlers.CORS()(r),
		Handler: r,
	}

	erro := make(chan error)
	quit := make(chan os.Signal)

	go func() {
		erro <- s.ListenAndServe()
	}()
	logger.Info("Server initializing on %s:%s", host, port)

	signal.Notify(quit, os.Kill, syscall.SIGKILL)

	for {
		select {
		case <-erro:
			logger.Error("Server error. Reason %v", <-erro)
			os.Exit(1)
		case <-quit:
			err := s.Close()

			if err != nil {
				msg := "Server close failed and is exiting either way." +
					"Reason %v"
				logger.Error(msg, err)

				os.Exit(1)
			}
			os.Exit(0)
		}
	}
}
