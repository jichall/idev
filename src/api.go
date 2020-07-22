package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jichall/idev/src/parser"
)

// HandleSpecific was created because on the challenge description it is not
// clear what method and what payload (if any) is sent with the request and I
// did this as a filler for what I found it was related to the description
// provided.
//
// The challenge didn't provide a clear description as how or what is the
// payload sent to the server.
//
func HandleSpecific(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var request Request
	err = json.Unmarshal(body, &request)

	if err != nil {
		logger.Error("Failed to unmarshal request data into response. Reason %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s := Response{}
	ok := s.prepare(request.Hostname)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, err := json.Marshal(&s)

	if err != nil {
		logger.Error("Failed to marshal response data. Reason %v", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Write(response)
	}
}

// HandleRaw is needed for testing the application's parser capabilities.
func HandleRaw(w http.ResponseWriter, r *http.Request) {
	hostname := mux.Vars(r)["hostname"]

	if len(hostname) > 0 {
		server := collection[hostname]

		b, err := json.Marshal(&server)

		if err != nil {
			logger.Error("Failed to marshal response data Reason %v", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Write(b)
		}
	} else {
		var servers []*parser.ServerData

		for _, v := range collection {
			servers = append(servers, v)
		}

		type RawResponse struct {
			S []*parser.ServerData `json:"server"`
		}

		load := RawResponse{
			S: servers,
		}

		b, err := json.Marshal(&load)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Write(b)
		}
	}
}

// HandleAll returns the information about all the processed servers sorted in
// usage rate
func HandleAll(w http.ResponseWriter, r *http.Request) {
	// the collection of servers that will be sent to the client as JSON
	var coll ResponseCollection

	for hostname := range collection {
		s := Response{}
		s.prepare(hostname)

		coll = append(coll, s)
	}

	response, err := json.Marshal(&coll)

	if err != nil {
		logger.Error("Failed to marshal response data. Reason %v", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Write(response)
	}
}

// HandleServer returns information about one specific server
func HandleServer(w http.ResponseWriter, r *http.Request) {
	hostname := mux.Vars(r)["hostname"]

	if len(hostname) > 0 {
		s := Response{}
		ok := s.prepare(hostname)

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		response, err := json.Marshal(&s)

		if err != nil {
			logger.Error("Failed to marshal response data. Reason %v", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.Write(response)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
