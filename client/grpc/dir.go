package grpc

import (
	"context"
	"pearviewer/client/dto"
	"pearviewer/client/types"
	"pearviewer/client/utils"
	pb "pearviewer/generated"
)

func ListDir(dirName, pathName string) {
	client, conn := createDirClient()
	request := dto.CreateListDirReq(dirName, pathName)
	log.Info("List Dir Request created: " + request.String())
	listDirReq(*client, request)
	closeClient(conn)
}

func MoveDir(dirName, oldPathName, newPathName string) {
	client, conn := createDirClient()
	request := dto.CreateMoveDirReq(dirName, oldPathName, newPathName)
	log.Info("Move Dir Request created: " + request.String())
	moveDirReq(*client, request)
	closeClient(conn)
}

func UploadDir(dirName, oldPathName, newPathName string) {
	request := dto.CreateUploadDirTree(dirName, oldPathName, newPathName)
	log.Info("Upload Dir Tree created")
	uploadDirReq(request)
}

func CreateDir(dirName string, pathname string) {
	client, conn := createDirClient()
	request := dto.CreateDirReq(dirName, pathname)
	log.Info("Upload Dir Request created: " + request.String())
	createDirReq(*client, request)
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

func createDirReq(client pb.DirServiceClient, request *pb.CreateDirReq) {
	response, err := client.CreateDir(context.Background(), request)
	if err != nil {
		log.Error("Upload Dir Request error: " + err.Error())
	}
	log.Info("Upload Dir Response: " + response.String())
}

func renameDirReq(client pb.DirServiceClient, request *pb.RenameDirReq) {
	response, err := client.RenameDir(context.Background(), request)
	if err != nil {
		log.Error("Rename Dir Request error: " + err.Error())
	}
	log.Info("Rename Dir Response: " + response.String())
}

func deleteDirReq(client pb.DirServiceClient, request *pb.DeleteDirReq) {
	response, err := client.DeleteDir(context.Background(), request)
	if err != nil {
		log.Error("Delete Dir Request error: " + err.Error())
	}
	log.Info("Delete Dir Response: " + response.String())
}

func uploadDirReq(req *types.Dir) {
	// MkDir root dir
	CreateDir(req.Name(), req.NewPathName())
	// Upload Files
	for _, file := range req.Files() {
		oldName := utils.Joins(file.OldPathName(), file.Name())
		UploadFile(oldName, file.NewPathName())
	}

	// Upload next dir
	for _, dir := range req.Children() {
		uploadDirReq(dir)
	}
}

func moveDirReq(client pb.DirServiceClient, request *pb.MoveDirReq) {
	response, err := client.MoveDir(context.Background(), request)
	if err != nil {
		log.Error("Move Dir Request error: " + err.Error())
	}
	log.Info("Move Dir Response: " + response.String())
}

func listDirReq(client pb.DirServiceClient, request *pb.ListDirReq) {
	response, err := client.ListDir(context.Background(), request)
	if err != nil {
		log.Error("List Dir Request error: " + err.Error())
	}
	log.Info("List Dir Response: " + response.String())
}
