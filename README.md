1. api
2. testing
3. documentation
4. document my api (swagger)
5. handlers
6. utils
7. types/structs
8. linting

# tools

1. protobuf
2. godoc
3. swag

### Documentation

To be able to view the entire structure of the code as well the documentation, godocs package is used.
This is easily installed using

```bash
	go install golang.org/x/tools/cmd/godoc@latest
```

This can then be served locally using any port of choice as follows

```bash
	godoc -http=:6060
```

Navigate then to the link `http://localhost:6060` in the browser to explore the docs as well as any added examples from the source code as well as from external packages.
The documentation can be found at either http://localhost:8080/swagger/index.html#/ or http://localhost:80/swagger/index.html#/ depending on the desired port

## swagger documentation

The documentation is generated using the swaggo library in go
Some examples and further documetations can be found here

```bash
$ go install github.com/swaggo/swag/cmd/swag@latest
$ swag init
```

## install linter

binary will be $(go env GOPATH)/bin/golangci-lint

```bash
$ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.3
$ golangci-lint --version
```

# plugins
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

# To download the required protocol buffer binary, depending on your pc of choice
```bash
	PB_REL="https://github.com/protocolbuffers/protobuf/releases"
	curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
	unzip protoc-3.15.8-linux-x86_64.zip -d $HOME/Desktop/bin
	export PATH="$PATH:$HOME/Desktop/bin"
```

# Generate go types and attributes from the .proto files
```bash
	protoc -I=pb --go_out=. pb/\*.proto
```

# To install godoc
```bash
	go install golang.org/x/tools/cmd/godoc@latest
```

# To watch the docs locally
```bash
	godoc -http=:6060
```

# Documentation
To work with swagger documentation, we can downalod and generate documentation using the commands below.

```bash
	$ go install github.com/swaggo/swag/cmd/swag@latest
	$ swag init
```

Happy hacking!!!