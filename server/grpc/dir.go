package grpc

import (
	"context"
	"os"
	pb "pearviewer/server/generated/dir"
)

func (s *dirServer) UploadDir(ctx context.Context, request *pb.UploadDirReq) (*pb.UploadDirRes, error) {
	var returnCode int32
	var message string

	log.Info("Received UploadDirReq: " + request.String())
	name := request.GetPathname() + request.Dir.GetName()
	if !isDirExist(name) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			returnCode = 500
			message = err.Error()
		}
		returnCode = 0
		message = "Directory Created: " + name
	} else {
		returnCode = 1
		message = "Directory Already Exists"
	}

	response := &pb.UploadDirRes{
		ReturnCode: returnCode,
		Message:    message,
	}
	return response, nil
}

func isDirExist(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}
