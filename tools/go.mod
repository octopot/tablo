module go.octolab.org/ecosystem/tablo/tools

go 1.13

require (
	github.com/go-swagger/go-swagger v0.23.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.4.3
	github.com/golang/protobuf v1.4.2
	github.com/golangci/golangci-lint v1.27.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/kamilsk/egg v0.0.16
	github.com/mailru/easyjson v0.7.1
	github.com/twitchtv/twirp v5.11.0+incompatible
	golang.org/x/exp v0.0.0-20200513190911-00229845015e
	golang.org/x/tools v0.3.3
)

replace github.com/izumin5210/gex => github.com/kamilsk/gex v0.6.0-e4

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.3
