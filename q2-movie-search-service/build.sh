go generate ./...
mockgen -destination=internal/mocks/mock_context.go -package=mocks github.com/labstack/echo/v4 Context

protoc --go_out=. protos/movie.proto
protoc --go-grpc_out=. protos/movie.proto

go mod tidy
go fmt ./...
