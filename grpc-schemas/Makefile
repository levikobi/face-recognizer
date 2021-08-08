generate:
	protoc --go_out=golang/face-recognition --go_opt=paths=source_relative \
        --go-grpc_out=golang/face-recognition --go-grpc_opt=paths=source_relative \
        protos/face-recognition-service.proto

protos:
	protoc -I protos/ protos/face-recognition-service.proto --go_out=plugins=grpc:golang/face-recognition