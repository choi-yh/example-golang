GOOGLEAPIS_DIR=$$GOPATH/src/googleapis

build-sample-proto:
	protoc -I . \
	-I $(GOOGLEAPIS_DIR) \
	--go_out . \
	--go_opt paths=source_relative \
	--go-grpc_out . \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out . \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	protos/sample/sample.proto