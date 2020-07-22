package main

// Request ...
type Request struct {
	Hostname string `json:"hostname"`
}

// Response is the struct that abstracts the data sent to the client as JSON
type Response struct {
	Hostname string `json:"hostname"`

	CPU struct {
		Mean float64   `json:"mean"`
		Mode []float64 `json:"mode"`
	}

	Memory struct {
		Mean float64   `json:"mean"`
		Mode []float64 `json:"mode"`
	}

	Disk struct {
		Mean float64   `json:"mean"`
		Mode []float64 `json:"mode"`
	}

	Usage float64
}

// ResponseCollection is a collection of Response
type ResponseCollection []Response

// prepare fills a response with the required data from a server hostname and
// returns a bool indicating whether or not a server was found with the given
// hostname
func (m *Response) prepare(hostname string) bool {
	server := collection[hostname]

	if server != nil {
		m.Hostname = hostname
		m.CPU.Mean = server.CPUStats.Mean
		m.CPU.Mode = server.CPUStats.Mode
		m.Memory.Mean = server.MemoryStats.Mean
		m.Memory.Mode = server.MemoryStats.Mode
		m.Disk.Mean = server.DiskStats.Mean
		m.Disk.Mode = server.DiskStats.Mode

		m.Usage = server.Usage

		return true
	}

	return false
}
