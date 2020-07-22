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
	// think it would make readability worse for new commers.

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
