# Go > Mux

*Build*
```shell
go build -o bin/main.exe main.go
```

*Lint*
```shell
go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
golangci-lint run
```

*Test*
```shell
go get -u github.com/mcubik/goverreport
go test apifarm/src
```

*Mocking*
```shell
go get github.com/vektra/mockery/v2/.../
mockery --all --keeptree
```