package parser

/*{
	"hostname": "server7",
	"cpu_load": {
	 "Value": 0.5627961727877718,
	 "Unit": "%"
	},
	"memory_size": {
	 "Value": 16,
	 "Unit": "GB"
	},
	"memory_usage": {
	 "Value": 5.881018176090025,
	 "Unit": "GB"
	},
	"disk_size": {
	 "Value": 100,
	 "Unit": "GB"
	},
	"disk_usage": {
	 "Value": 33.89572263431528,
	 "Unit": "GB"
	}
   },*/

type (
	// Server ...
	Server struct {
		Hostname    string      `json:"hostname"`
		CPU         ServerParam `json:"cpu_load"`
		MemorySize  ServerParam `json:"memory_size"`
		MemoryUsage ServerParam `json:"memory_usage"`
		DiskSize    ServerParam `json:"disk_size"`
		DiskUsage   ServerParam `json:"disk_usage"`
	}

	// ServerParam ...
	ServerParam struct {
		Value string
		Unit  string
	}

	// ServerCollection ...
	ServerCollection []Server
)
