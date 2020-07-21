# The iDev interview

This is a golang challenge where I read a set of data from a .json file and
serve its content over an HTTP server.

# How to Build

## Prerequisites

- [go 1.14](https://golang.org/dl/)
- [jq](https://stedolan.github.io/jq/download/)

You have to first generate the files to be readd by the program with the
`filter.sh` script, to do that you can simply invoke the program `filter.sh
data.json data` and that will read the **data.json** file and output files on
**data**.

## Linux

After that you can simply use the Makefile with the already built target for
you, to use the Makefile just run `make` on the project root folder and it will
generate a binary called `stats.out` under **bin/**

## Windows

To build it on windows you can simply invoke `go build -o bin/stats.exe
src/main.go src/api.go src/server.go` and it will generate you a binary under
**bin/**.

# Documentation