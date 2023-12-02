.PHONY: proto
.PHONY: gql

OUTDIR = proto

proto: clean proto-generate

tidy:
	@go mod tidy
	@echo " > Go package tidy completed"

clean:
	@rm -f $(OUTDIR)/*.pb.go || true

proto-generate:
	@echo " > Generating proto files"
# require_unimplemented_servers=false:
	@protoc -I $(OUTDIR)/ --go_out=. --go-grpc_out=. ./$(OUTDIR)/*.proto
	@echo " > Done generating proto files"

gql:
	@echo " > Generating graphql files"
	@cd proxy_service/ && go run github.com/99designs/gqlgen generate
	@echo " > Done generating graphql files"