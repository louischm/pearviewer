package grpc

import (
	"context"
	. "pearviewer/client/generated/dir"
)

func UploadDir(dirName string, pathname string) {
	client, conn := createDirClient()
	request := createUploadDirReq(dirName, pathname)
	log.Info("Upload Dir Request construct: " + request.String())
	uploadDirReq(*client, request)
	closeClient(conn)
}

func RenameDir(oldName, newName, pathName string) {
	client, conn := createDirClient()
	request := createRenameDirReq(oldName, newName, pathName)
	log.Info("Rename Dir Request construct: " + request.String())
	renameDirReq(*client, request)
	closeClient(conn)
}

func createUploadDirReq(dirName, pathName string) *UploadDirReq {
	dir := Dir{
		Name: dirName,
	}
	request := &UploadDirReq{
		Dir:      &dir,
		Pathname: pathName,
	}
	return request
}

func createRenameDirReq(oldName, newName, pathName string) *RenameDirReq {
	request := &RenameDirReq{
		OldName:  oldName,
		NewName:  newName,
		PathName: pathName,
	}
	return request
}

func uploadDirReq(client DirServiceClient, request *UploadDirReq) {
	response, err := client.UploadDir(context.Background(), request)
	if err != nil {
		log.Error("Upload Dir Request error: " + err.Error())
	}
	log.Info("Upload Dir Response: " + response.String())
}

func renameDirReq(client DirServiceClient, request *RenameDirReq) {
	response, err := client.RenameDir(context.Background(), request)
	if err != nil {
		log.Error("Rename Dir Request error: " + err.Error())
	}
	log.Info("Rename Dir Response: " + response.String())
}
