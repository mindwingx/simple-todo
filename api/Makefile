.PHONY: env
env:
	@cp .env.example .env
	@echo ".env generated"

.PHONY: swag
swag:
	@echo "swagger generating..."
	@swag init -g ./cmd/main.go -q -o ./docs
	@echo "swagger generated"

proto:
	@mkdir -p pb/gen/go

	@protoc -I pb/proto \
	--go_out pb/gen/go \
	--go_opt paths=source_relative \
	--go-grpc_out pb/gen/go \
	--go-grpc_opt paths=source_relative \
	  pb/proto/*/*.proto

	@echo "proto files generated"