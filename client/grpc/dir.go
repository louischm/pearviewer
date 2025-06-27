package grpc

import (
	"context"
	"github.com/louischm/pkg/utils"
	"os"
	"pearviewer/client/dto"
	"pearviewer/client/types"
	pb "pearviewer/generated"
)

func ListDir(dirName, pathName string) *pb.ListDirRes {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateListDirReq(dirName, pathName)
	log.Info("List Dir Request created: " + request.String())
	res := listDirReq(*client, request)
	return res
}

func MoveDir(dirName, oldPathName, newPathName string) {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateMoveDirReq(dirName, oldPathName, newPathName)
	log.Info("Move Dir Request created: " + request.String())
	moveDirReq(*client, request)
}

func UploadDir(dirName, oldPathName, newPathName string) {
	request := dto.CreateUploadDirTree(dirName, oldPathName, newPathName)
	log.Info("Upload Dir Tree created")
	uploadDirReq(request)
}

func DownloadDir(dirName, sourcePathName, destPathName string) {
	listDir := ListDir(dirName, sourcePathName)
	log.Info("Download Dir Request created: " + listDir.String())
	downloadDirReq(listDir.Dir, sourcePathName, destPathName)
}

func CreateDir(dirName string, pathname string) {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateDirReq(dirName, pathname)
	log.Info("Upload Dir Request created: " + request.String())
	createDirReq(*client, request)
}

func RenameDir(oldName, newName, pathName string) {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateRenameDirReq(oldName, newName, pathName)
	log.Info("Rename Dir Request created: " + request.String())
	renameDirReq(*client, request)
}

func DeleteDir(dirName, pathname string) {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateDeleteDirReq(dirName, pathname)
	log.Info("Delete Dir request created: " + request.String())
	deleteDirReq(*client, request)
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

func listDirReq(client pb.DirServiceClient, request *pb.ListDirReq) *pb.ListDirRes {
	response, err := client.ListDir(context.Background(), request)
	if err != nil {
		log.Error("List Dir Request error: " + err.Error())
	}
	log.Info("List Dir Response: " + response.String())
	return response
}

func downloadDirReq(dir *pb.Dir, sourcePathName, destPathName string) {
	createSourceDir(dir, destPathName)
	sourceName := utils.Joins(sourcePathName, dir.DirName)
	destName := utils.Joins(destPathName, dir.DirName)
	for _, child := range dir.GetDir() {
		downloadDirReq(child, sourceName, destName)
	}

	for _, file := range dir.GetFile() {
		DownloadFile(file.Name, sourceName, destName)
	}
}

func createSourceDir(dir *pb.Dir, destPathName string) {
	name := utils.Joins(destPathName, dir.DirName)

	if !utils.IsDirExist(name) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			log.Debug("Create Dir error: " + err.Error())
		}
		log.Info("Create Dir: " + name)
	} else {
		log.Debug("Dir already exist: " + name)
	}
}
