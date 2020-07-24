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

Although you could create a binary on windows the software will only work after
pre processing the data. Therefore you'll need `jq`, that doesn't comes with
windows natively.

# Documentation

To use the software you have to build the application it and prepare the data.
Only then you can execute it and use its HTTP RESTful api. To do that read the
**Build** section above.

The binary itself has four flags you can set, they are described in the table
below.

| Flag | Default | Description |
|------|---------|-------------|
| host | `localhost`| IP address to where initialize the HTTP server |
| port | `8080` | The port used by the HTTP server |
| folder | `data/` | The folder used by the application to read telemetry from JSON files |
| production | `false` | Makes the environment run on production increasing performance by a bit |

To initialize the application invoke the binary with the command `./stats.out`
in your terminal. If you turned the production mode on you won't see any output
sent to the stdout, it will instead be sent to a log file under a log folder
created by the binary with the name given by the time of execution.

### Routes

The available routes are defined below.

| Route | Method | Variable/Body | Description |
|-------|--------|----------|-------------|
| /servers/ | GET | none | Returns statistical information data of all servers |
| /servers/{hostname} | GET | `hostname` | Returns statistical information about a specific server hostname |
| /severs/raw/ | GET | none | Returns all data and statistical information about a specific server hostname |
| /severs/raw/{hostname} | GET | `hostname` | Returns all data and statistical information about a specific server hostname |
| /spec | POST | `{hostname: "serverX"}` where `X` is the server number | Returns statistical information about a specific server hostname |

Each endpoint returns a JSON containing the needed information to assess the
status of the server(s). The response of a call to
`curl -X GET localhost:8080/servers/server` would return:

```
{
    "hostname":"server0",
    "CPU":{"mean":0.5030763415151056,"mode":[]},
    "Memory":{"mean":4.960234299570407,"mode":[]},
    "Disk":{"mean":25.77459401368746,"mode":[]},
    "Usage":0.3987051870811867
}
```

Have in mind that the `/spec` route returns data identical as the
`/servers/{hostname}` route. It was implemented because the description of the
challenge wasn't that clear and what appears to be a body request data is what I
implemented in this route. If you would like to use the route `/spec` you have
to send the required data as follows:

`curl -X POST --data '{"hostname":"server0"}'` and  that would return:

```
{
    "hostname":"server0",
    "CPU":{"mean":0.5030763415151056,"mode":[]},
    "Memory":{"mean":4.960234299570407,"mode":[]},
    "Disk":{"mean":25.77459401368746,"mode":[]},
    "Usage":0.3987051870811867
}
```

### Responses

The mean values of a response are given in the original processed units, that
is, the CPU is given in percentage, the memory is given in GB and the disk is
also given in GB.

The exception for that is the `Usage` value, that is given in an interval of [0,
1] representing the usage of the resources of that server.

# The development

I had to first tidy the `data.json` file, otherwise I could have problems with
files too big for a RAM to handle when parsing it. I created a script to extract
each JSON object from the array and wrote a file in a `data-<number>` format
insde the data folder under `bin`.

The problem of loading big files to RAM wasn't the only one. The only way that
structure could be parsed would be a tedious and quite big one because one would
have to use a lot of `map[string]interface{}` and then convert the data to the
one needed. It would increase code size by a lot and would be a headache to
maintain.

With this new architecture one could use concurrency to calculate the
statistical data for each server although this comes with its own problems as it
should be used with a thread safe map access using `sync.Map` or atomic
operations and it would possibly slow the application as a consequence. That
being said I chose to not go that way and used a simple map to map a hostname to
a server data.