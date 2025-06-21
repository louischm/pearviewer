rm -rf ./generated
mkdir ./generated
cd ./generated
go mod init generated
cd ..

protoc --go_out=./ --go-grpc_out=./ -I./proto ./proto/file.proto
protoc --go_out=./ --go-grpc_out=./ -I./proto ./proto/dir.proto

cd ./generated
go mod tidy