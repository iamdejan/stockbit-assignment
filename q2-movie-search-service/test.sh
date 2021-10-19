go clean -testcache

CVPKG=$(go list ./... | grep -v mocks | grep -v pb | tr '\n' ',')
go test -race -coverpkg $CVPKG -coverprofile=coverage.out ./...
