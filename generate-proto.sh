rm -f ./server/generated/*
rm -f ./client/generated/*
protoc --go_out=./server/ --go-grpc_out=./server/ -I./proto ./proto/file.proto
protoc --go_out=./client/ --go-grpc_out=./client/ -I./proto ./proto/file.proto
protoc --go_out=./server/ --go-grpc_out=./server/ -I./proto ./proto/dir.proto
protoc --go_out=./client/ --go-grpc_out=./client/ -I./proto ./proto/dir.proto