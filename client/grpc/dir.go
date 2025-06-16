package grpc

import (
	"context"
	pb "pearviewer/client/generated/dir"
)

func UploadDir(dirName string, pathname string) {
	client, conn := createDirClient()
	request := createUploadDirReq(dirName, pathname)
	log.Info("Upload Dir Request construct: " + request.String())
	uploadDirReq(*client, request)
	closeClient(conn)
}

func createUploadDirReq(dirName, pathName string) *pb.UploadDirReq {
	dir := pb.Dir{
		Name: dirName,
	}
	request := &pb.UploadDirReq{
		Dir:      &dir,
		Pathname: pathName,
	}
	return request
}

func uploadDirReq(client pb.DirServiceClient, request *pb.UploadDirReq) {
	response, err := client.UploadDir(context.Background(), request)
	if err != nil {
		log.Error("Upload Dir Request error: " + err.Error())
	}
	log.Info("Upload Dir Response: " + response.String())
}
