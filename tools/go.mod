module tools

go 1.11

require (
	github.com/go-swagger/go-swagger v0.21.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/golangci/golangci-lint v1.21.0
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/twitchtv/twirp v5.8.0+incompatible
	golang.org/x/tools v0.0.0-20191021201049-44c9a601ac60
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.1
