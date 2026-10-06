

generate:
	protoc \
		--proto_path=./api/proto/random \
		--go_out=./gen/go/random \
		--go_opt=paths=source_relative \
		--go-grpc_out=./gen/go/random \
		--go-grpc_opt=paths=source_relative \
		api/proto/random/*.proto