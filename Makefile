#
# Build make file for Linux and go1.14 where go must be available system wide
# (on $path).
#
# Rafael Campos Nunes <rafaelnunes@engineer.com>
#

CC=`which go`
SRC=src/main.go src/server.go src/api.go src/process.go src/message.go

OUT=./bin/stats.out

all:
	$(CC) build -o "$(OUT)" $(SRC)