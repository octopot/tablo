.PHONY: protobuf-generate
protobuf-generate: protobuf-generate-v1
protobuf-generate: protobuf-generate-v2
protobuf-generate: protobuf-generate-swagger

.PHONY: protobuf-generate-v1
protobuf-generate-v1:
	@protoc --proto_path=./api/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec/googleapis \
	        --go_out=paths=source_relative:./internal/generated/api \
	        --twirp_out=paths=source_relative:./internal/generated/api \
	        $(shell ls api/protobuf-spec/v1 | grep .proto | grep -v version | awk '{print "v1/"$$1}')
	@protoc --proto_path=./api/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec/googleapis \
	        --swagger_out=allow_merge=true,merge_file_name=tablo:./api/openapi-spec \
	        metadata.proto
	@protoc --proto_path=./api/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec/googleapis \
	        --swagger_out=allow_merge=true,merge_file_name=v1:./api/openapi-spec \
	        $(shell ls api/protobuf-spec/v1 | grep .proto | awk '{print "v1/"$$1}')

.PHONY: protobuf-generate-v2
protobuf-generate-v2:
	@protoc --proto_path=./api/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec/googleapis \
	        --go_out=paths=source_relative:./internal/generated/api \
	        --twirp_out=paths=source_relative:./internal/generated/api \
	        $(shell ls api/protobuf-spec/v2 | grep .proto | grep -v version | awk '{print "v2/"$$1}')
	@swagger mixin \
	         -o api/openapi-spec/v1.json \
	         api/openapi-spec/v1.swagger.json \
	         api/openapi-spec/tablo.swagger.json &>/dev/null
	@protoc --proto_path=./api/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec/googleapis \
	        --swagger_out=allow_merge=true,merge_file_name=v2:./api/openapi-spec \
	        $(shell ls api/protobuf-spec/v2 | grep .proto | awk '{print "v2/"$$1}')

.PHONY: protobuf-generate-swagger
protobuf-generate-swagger:
	@swagger mixin \
	         -o api/openapi-spec/v2.json \
	         api/openapi-spec/v2.swagger.json \
	         api/openapi-spec/tablo.swagger.json &>/dev/null
	@rm api/openapi-spec/*.swagger.json
	@for json in $$(ls api/openapi-spec/*.json); do echo >> $$json; done

.PHONY: protobuf-update
protobuf-update: VERSION = $(shell protoc --version | awk '{print $$2}')
protobuf-update: protobuf-update-google/protobuf

.PHONY: protobuf-update-google/protobuf
protobuf-update-google/protobuf:
	@rm -f bin/protobuf-$(VERSION).tar.gz
	@curl -sL -o bin/protobuf-$(VERSION).tar.gz \
	      https://github.com/protocolbuffers/protobuf/archive/v$(VERSION).tar.gz
	@rm -rf third_party/protobuf-spec/google/protobuf/*
	@tar -tf bin/protobuf-$(VERSION).tar.gz \
	| grep 'protobuf-$(VERSION)/src/google/protobuf/[^/]*.proto$$' \
	| grep -v 'unittest' \
	| grep -v '/test_' \
	| grep -v '_test' \
	| xargs tar --strip-components=4 \
	            -C third_party/protobuf-spec/google/protobuf \
	            -xf bin/protobuf-$(VERSION).tar.gz
	@rm -f bin/protobuf-$(VERSION).tar.gz
