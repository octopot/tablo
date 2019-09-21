module tools

go 1.11

replace github.com/pelletier/go-toml => github.com/kamilsk/go-toml v1.4.0-asd-patch

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.0-20190618115843-d350ce7f7a97

require (
	github.com/go-swagger/go-swagger v0.20.1
	github.com/gogo/protobuf v1.2.1
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/golangci/golangci-lint v1.17.1
	github.com/goreleaser/godownloader v0.0.0-20190907185828-93b2b793cd90
	github.com/goreleaser/goreleaser v0.110.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.6
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/twitchtv/twirp v5.8.0+incompatible
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	golang.org/x/sys v0.0.0-20190907184412-d223b2b6db03 // indirect
	golang.org/x/tools v0.0.0-20190908122318-4c50eace5a90
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)
