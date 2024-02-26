generate:
	make generate-term-api

generate-term-api:
	mkdir -p pkg/clientTerm
	protoc --proto_path api/clientTerm \
	--go_out=pkg/clientTerm --go_opt=paths=source_relative \
	--go-grpc_out=pkg/clientTerm --go-grpc_opt=paths=source_relative \
	api/clientTerm/clientTerm.proto
