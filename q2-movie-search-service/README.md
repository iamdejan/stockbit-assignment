# Movie Search Service

This server is for searching movies from [OMBD Api](https://www.omdbapi.com/).

## Getting started

### Prerequisites

1) Go language compiler (this server uses version 1.17.1).
    - After Go language compiler is installed, you have to set `GOROOT` and `GOPATH`.
2) Protobuf compiler.
    - for Mac and Linux users, install using Homebrew: `$ brew install protobuf`
3) Protobuf generator for Go language (`protoc-gen-go` and `protoc-gen-go-grpc`).
    - for Mac and Linux users, install using Homebrew: `$ brew install protoc-gen-go protoc-gen-go-grpc`
4) `mockgen` library installed.
    - for Mac and Linux users, install with this command: `$ go install github.com/golang/mock/mockgen`

### Build

Run `$ sh build.sh`. This script will:

1) generate Go files based on Protobuf spec;
2) download necessary dependencies; and
3) reformat the code.

### Run the service

Run `$ sh run-server.sh`. This command will run **both** HTTP server and gRPC server.

### Run gRPC client for testing RPC call

Run `$ sh run-grpc-client.sh`.

### Run unit tests

Run `$ sh test.sh`. This command will generate `coverage.out` file, which can be used to visualize the code coverage
line-by-line. To see the the line-by-line coverage in HTML format, run `$ go tool cover -html=coverage.out` in laptop,
not in server.
