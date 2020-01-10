module go.octolab.org/ecosystem/tablo/tools

go 1.11

require (
	github.com/go-swagger/go-swagger v0.21.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/golangci/golangci-lint v1.22.2
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/kamilsk/egg v0.0.8
	github.com/twitchtv/twirp v5.8.0+incompatible
	golang.org/x/tools v0.2.2
)

replace github.com/izumin5210/gex => github.com/kamilsk/gex v0.6.0-e3

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.1
