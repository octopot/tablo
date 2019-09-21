module tools

go 1.11

replace github.com/pelletier/go-toml => github.com/kamilsk/go-toml v1.4.0-asd-patch

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.0-20190921135421-dca3d7403570

require (
	github.com/go-swagger/go-swagger v0.20.1
	github.com/gogo/protobuf v1.3.0
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.3.2
	github.com/golangci/golangci-lint v1.18.0
	github.com/grpc-ecosystem/grpc-gateway v1.11.2
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/twitchtv/twirp v5.8.0+incompatible
	golang.org/x/tools v0.0.0-20190909030654-5b82db07426d
)
