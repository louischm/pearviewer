package grpc

import (
	"context"
	"os"
	pb "pearviewer/generated/dir"
	"pearviewer/server/utils"
)

func (s *dirServer) UploadDir(ctx context.Context, request *pb.UploadDirReq) (*pb.UploadDirRes, error) {
	var returnCode int32
	var message string

	log.Info("Received UploadDirReq: " + request.String())
	name := request.GetPathname() + request.GetDir().GetName()
	if !utils.IsDirExist(name) {
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

func (s *dirServer) RenameDir(ctx context.Context, request *pb.RenameDirReq) (*pb.RenameDirRes, error) {
	var returnCode int32
	var message string

	log.Info("Received RenameDirReq: " + request.String())
	oldName := request.GetPathName() + request.GetOldName()
	newName := request.GetPathName() + request.GetNewName()
	if utils.IsDirExist(oldName) {
		err := os.Rename(oldName, newName)
		if err != nil {
			returnCode = 500
			message = err.Error()
			log.Error(message)
		}
		message = "Directory Renamed: " + oldName + " to " + newName
		log.Info(message)
		returnCode = 0
	} else {
		returnCode = 1
		message = "Directory: " + oldName + " doesn't exists"
		log.Info(message)
	}

	response := &pb.RenameDirRes{
		ReturnCode: returnCode,
		Message:    message,
	}
	return response, nil
}

func (s *dirServer) DeleteDir(ctx context.Context, request *pb.DeleteDirReq) (*pb.DeleteDirRes, error) {
	var returnCode int32
	var message string

	log.Info("Received DeleteDirReq: " + request.String())
	name := request.GetPathName() + request.GetDirName()
	if utils.IsDirExist(name) {
		err := os.Remove(name)
		if err != nil {
			returnCode = 500
			message = "Dir: " + name + " not deleted: " + err.Error()
			log.Error(message)
		}
		returnCode = 0
		message = "Directory Deleted: " + name
		log.Info(message)
	} else {
		returnCode = 1
		message = "Directory Not Found: " + name
		log.Info(message)
	}

	response := &pb.DeleteDirRes{
		ReturnCode: returnCode,
		Message:    message,
	}
	return response, nil
}
