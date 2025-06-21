package grpc

import (
	"context"
	"pearviewer/client/dto"
	. "pearviewer/generated/dir"
)

func UploadDir(dirName string, pathname string) {
	client, conn := createDirClient()
	request := dto.CreateUploadDirReq(dirName, pathname)
	log.Info("Upload Dir Request created: " + request.String())
	uploadDirReq(*client, request)
	closeClient(conn)
}

func RenameDir(oldName, newName, pathName string) {
	client, conn := createDirClient()
	request := dto.CreateRenameDirReq(oldName, newName, pathName)
	log.Info("Rename Dir Request created: " + request.String())
	renameDirReq(*client, request)
	closeClient(conn)
}

func DeleteDir(dirName, pathname string) {
	client, conn := createDirClient()
	request := dto.CreateDeleteDirReq(dirName, pathname)
	log.Info("Delete Dir request created: " + request.String())
	deleteDirReq(*client, request)
	closeClient(conn)
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

func deleteDirReq(client DirServiceClient, request *DeleteDirReq) {
	response, err := client.DeleteDir(context.Background(), request)
	if err != nil {
		log.Error("Delete Dir Request error: " + err.Error())
	}
	log.Info("Delete Dir Response: " + response.String())
}
