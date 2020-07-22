package main

import (
	"github.com/jichall/idev/src/parser"
	"github.com/montanaflynn/stats"
)

func process(servers parser.ServerCollection) {
	for hostname, server := range servers {
		calculate(hostname, server)
	}
}

func calculate(hostname string, server *parser.ServerData) {
	// I could arrays to represent data which would make the code smaller but I
	// think it would make readability worse for new comers.
	
	cpuMean, err := stats.Mean(server.CPU)
	memoryMean, err := stats.Mean(server.MemoryUsage)
	diskMean, err := stats.Mean(server.DiskUsage)

	if err != nil {
		msg := "Couldn't calculate mean stats of server %s. Reason %v"
		logger.Error(msg, hostname, err)
	} else {
		server.CPUStats.Mean = cpuMean
		server.MemoryStats.Mean = memoryMean
		server.DiskStats.Mean = diskMean

		// The server usage status is a metric of how overloaded is the server
		// on a scale of 0-100. To evaluate that I use all the metrics of a
		// server in a weighted average.
		//
		// The CPU, the primal resource of a server has a weight of 10, the
		// following resources for the memory and disk usage, respectively, are
		// 7 and 3.
		//
		// The values are first normalized to use a percentage scale before
		// the weighted average formula.
		server.Usage = (cpuMean*10 + (memoryMean/16)*7 + (diskMean/100)*3) / 20
	}

	cpuMode, err := stats.Mode(server.CPU)
	memoryMode, err := stats.Mode(server.MemoryUsage)
	diskMode, err := stats.Mode(server.DiskUsage)

	if err != nil {
		msg := "Couldn't calculate mode stats of server %s. Reason %v"
		logger.Error(msg, hostname, err)
	} else {
		server.CPUStats.Mode = cpuMode
		server.MemoryStats.Mode = memoryMode
		server.DiskStats.Mode = diskMode
	}
}
