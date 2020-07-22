package parser

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

// Parse ...
func Parse(folder string) (ServerCollection, error) {

	files, err := ioutil.ReadDir(folder)

	if err != nil {
		return nil, err
	}

	collection := make(ServerCollection)

	for _, file := range files {
		var s Server

		b, err := ioutil.ReadFile(path.Join(folder, file.Name()))

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(b, &s)

		if err != nil {
			return nil, err
		}

		if collection[s.Hostname] != nil {
			collection[s.Hostname].CPU = append(collection[s.Hostname].CPU, s.CPU.Value)
			collection[s.Hostname].MemoryUsage = append(collection[s.Hostname].MemoryUsage, s.MemoryUsage.Value)
			collection[s.Hostname].DiskUsage = append(collection[s.Hostname].DiskUsage, s.DiskUsage.Value)
		} else {
			collection[s.Hostname] = &ServerData{
				CPU:         []float64{s.CPU.Value},
				MemorySize:  s.MemorySize.Value,
				MemoryUsage: []float64{s.MemoryUsage.Value},
				DiskSize:    s.DiskSize.Value,
				DiskUsage:   []float64{s.DiskUsage.Value},
			}
		}
	}

	return collection, nil
}
