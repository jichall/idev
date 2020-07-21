package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Parse ...
func Parse(filename string) (*ServerCollection, error) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var collection *ServerCollection

	var objects interface{}
	err = json.Unmarshal(data, &objects)

	if err != nil {
		return nil, err
	}

	// ensure it's an array
	array := objects.([]interface{})

	for _, object := range array {
		//
		raw := object.(Server)

		if err != nil {
			return nil, err
		} else {
			log.Printf("%v", raw)
			/*var server Server

			err = json.Unmarshal(raw, &server)

			if err != nil {
				return nil, err
			}

			collection = append(collection, server)*/
		}
	}

	return collection, nil
}
