submodule_build_proto:
	protoc --go_out=./proto/. --go-grpc_out=./proto/. ./proto/*.proto
	protoc --go_out=./proto/. --go-grpc_out=./proto/. --go_opt=Mproto/Match.proto=github.com/WolffunService/theta-shared-common/proto/coreproto ./proto/service_rivals.proto
