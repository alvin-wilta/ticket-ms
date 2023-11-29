.PHONY: proto

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
