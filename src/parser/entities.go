package parser

type (
	// Server is the abstraction of the server object inside a JSON.
	Server struct {
		Hostname    string      `json:"hostname"`
		CPU         ServerAttribute `json:"cpu_load"`
		MemorySize  ServerAttribute `json:"memory_size"`
		MemoryUsage ServerAttribute `json:"memory_usage"`
		DiskSize    ServerAttribute `json:"disk_size"`
		DiskUsage   ServerAttribute `json:"disk_usage"`
	}

	// ServerAttribute is a server attribute which is encoded in JSON
	ServerAttribute struct {
		Unit  string
		Value float64
	}

	// Data ...
	Data struct {
		CPU []float64
		MemorySize float64
		MemoryUsage []float64
		DiskSize float64
		DiskUsage []float64
	}

	// ServerCollection maps the hostname of a server to its telemetry data
	ServerCollection map[string]*Data
)
