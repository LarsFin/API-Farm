# Go > Mux

The package `gorilla/mux` provides a router for managing requests made to a server.

## Dependencies

Golang uses packages which are installed at your `GOPATH` in your `pkg` directory. With the addition of golang modules, it is possible to retrieve all dependencies referenced in the `go.mod` file using the command below;

```shell
go mod download
```

## Running

After installing the necessary dependencies build the executable;

```shell
go build
```

Then run the output executable file, which should be `apifarm` or `apifarm.exe` on Windows.

```shell
./apifarm
```

You can verify that the application is running successfully with a quick curl request to the `/ping` endpoint.

```console
C:\> curl http://localhost:8080/ping
pong
```

### Configuration

A configuration file exists to determine the address to bind on. As with other Api Farm api's, a separate config file exists for each environment under the naming pattern; `config.<environment>.json`.

Below, are the current settings of the application.

```json
{
    "host": "0.0.0.0",
    "port": 8080
}
```

## Testing

Ensure you have the necessary testing dependencies installed and run the command below to run the test suite;

```shell
go test apifarm/src
```

To utilise mock types and make more informative assertions, I've used the package `stretchr/testify`. It should be noted, that this package will not generate mocked types for you. To generate these, look below.

To mock interfaces, I've used the tool `vektra/mockery`. This generates a `mocks` directory at the root of the project and contains mock types implementing the various interfaces of the source code. To generate these mocks after changing an interface you can use the command below;

```shell
mockery --all --keeptree
```

To install `vektra/mockery` on your local, get the package like so;

```shell
go get github.com/vektra/mockery/v2/.../
```

## Linting

You must have installed the golangci-lint tool to perform a scan of the code base. The command below will install the necessary dependency;

```shell
go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
```

Once installed, you can run it;

```
golangci-lint run
```

`golangci-lint` runs a suite of linting tools. Which tools you wish to use and their own rules can be defined within the `.golangci.yaml` file.

## Docker

This API is supported with Docker. You can check out how to run it by following the instructions in the root README [here](https://github.com/LarsFin/API-Farm).

## Resource Documentation

- **golang:** https://golang.org/doc/
- **gorilla/mux:** https://github.com/gorilla/mux
- **stretchr/testify:** https://github.com/stretchr/testify
- **vektra/mockery:** https://github.com/vektra/mockery