module tools

go 1.12

replace github.com/pelletier/go-toml => github.com/kamilsk/go-toml v1.4.0-asd-patch

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.0-20190618115843-d350ce7f7a97

replace git.apache.org/thrift.git => github.com/apache/thrift v0.12.0

require (
	github.com/go-swagger/go-swagger v0.20.1
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/golangci/golangci-lint v1.17.1
	github.com/goreleaser/godownloader v0.0.0-20190803193356-7ef626c90bb6
	github.com/goreleaser/goreleaser v0.110.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.6
	github.com/twitchtv/twirp v5.8.0+incompatible
	golang.org/x/tools v0.0.0-20190905173453-6b3d1c9ba8bf
)
