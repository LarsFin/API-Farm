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
go test -coverprofile=coverage.out apifarm/src
goverreport -sort=block -order=desc -threshold=100
```

*Mocking*
```shell
mockgen -source <file-path> -destination mock/<file-name>_mock.go
```