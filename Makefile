include .env
LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
generate:
	make generate-term-api

generate-term-api:
	mkdir -p pkg/clientTerm
	protoc --proto_path api/clientTerm \
	--go_out=pkg/clientTerm --go_opt=paths=source_relative \
	--go-grpc_out=pkg/clientTerm --go-grpc_opt=paths=source_relative \
	api/clientTerm/clientTerm.proto
