

install:
	go install ./cmd/protoc-gen-gorpc

generate-tests:
	protoc --go_out=./tests/generated  --gorpc_out=./tests/generated --proto_path=./tests/proto --include_imports tests/proto/xcorp/protobuf/parking/parking_service.proto