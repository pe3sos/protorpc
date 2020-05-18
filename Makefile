

install:
	go install ./cmd/protoc-gen-gorpc

process-test-proto:
	protoc --go_out=./test/generated  --gorpc_out=./test/generated -I ./test/proto --include_imports test/proto/parking/parking_service.proto

test:
	go test -v ./...