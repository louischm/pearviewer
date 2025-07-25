package grpc

import (
	"context"
	"github.com/louischm/pkg/utils"
	"os"
	"pearviewer/client/dto"
	"pearviewer/client/types"
	pb "pearviewer/generated"
)

func GetRootPath(userName string) *pb.GetRootPathRes {
	client, coon := createDirClient()
	defer closeClient(coon)
	request := dto.CreateGetRootPathReq(userName)
	log.Info("GetRootPath request: %s", request.String())
	res := getRootPathReq(*client, request)
	return res
}

func ListDir(dirName, pathName string) *pb.ListDirRes {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateListDirReq(dirName, pathName)
	log.Info("List Dir Request created: %s", request.String())
	res := listDirReq(*client, request)
	return res
}

func MoveDir(dirName, oldPathName, newPathName string) {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateMoveDirReq(dirName, oldPathName, newPathName)
	log.Info("Move Dir Request created: %s", request.String())
	moveDirReq(*client, request)
}

func UploadDir(dirName, oldPathName, newPathName string) {
	request := dto.CreateUploadDirTree(dirName, oldPathName, newPathName)
	log.Info("Upload Dir Tree created")
	uploadDirReq(request)
}

/*
func DownloadDir(dirName, sourcePathName, destPathName string, done chan bool, fileDwnld chan float64) {
	listDir := ListDir(dirName, sourcePathName)
	log.Info("Download Dir Request created: %s", listDir.String())
	downloadDirReq(listDir.Dir, sourcePathName, destPathName, fileDwnld, 0)
	done <- true
}*/

func CreateDir(dirName string, pathname string) {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateDirReq(dirName, pathname)
	log.Info("Upload Dir Request created: %s", request.String())
	createDirReq(*client, request)
}

func RenameDir(oldName, newName, pathName string) {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateRenameDirReq(oldName, newName, pathName)
	log.Info("Rename Dir Request created: %s", request.String())
	renameDirReq(*client, request)
}

func DeleteDir(dirName, pathname string) {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateDeleteDirReq(dirName, pathname)
	log.Info("Delete Dir request created: %s", request.String())
	deleteDirReq(*client, request)
}

func GetFileNumber(dirName, pathName string) (*pb.GetFileNumberRes, error) {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateGetFileNumberReq(dirName, pathName)
	log.Info("GetFileNUmber request created: %s", request.String())
	return getFileNumberReq(*client, request), nil
}

func SearchFile(search, pathName, dirName string) *pb.ListDirRes {
	client, conn := createDirClient()
	defer closeClient(conn)
	request := dto.CreateFileSearchReq(search, pathName, dirName)
	log.Info("Search request created: %s", request.String())
	return searchFileReq(*client, request)
}

func createDirReq(client pb.DirServiceClient, request *pb.CreateDirReq) {
	response, err := client.CreateDir(context.Background(), request)
	if err != nil {
		log.Error("Upload Dir Request error: %s", err.Error())
	}
	log.Info("Upload Dir Response: %s", response.String())
}

func renameDirReq(client pb.DirServiceClient, request *pb.RenameDirReq) {
	response, err := client.RenameDir(context.Background(), request)
	if err != nil {
		log.Error("Rename Dir Request error: %s", err.Error())
	}
	log.Info("Rename Dir Response: %s", response.String())
}

func deleteDirReq(client pb.DirServiceClient, request *pb.DeleteDirReq) {
	response, err := client.DeleteDir(context.Background(), request)
	if err != nil {
		log.Error("Delete Dir Request error: " + err.Error())
	}
	log.Info("Delete Dir Response: %s", response.String())
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
		log.Error("Move Dir Request error: %s", err.Error())
	}
	log.Info("Move Dir Response: %s", response.String())
}

func listDirReq(client pb.DirServiceClient, request *pb.ListDirReq) *pb.ListDirRes {
	response, err := client.ListDir(context.Background(), request)
	if err != nil {
		log.Error("List Dir Request error: %s", err.Error())
	}
	log.Info("List Dir Response: %s", response.String())
	return response
}

func CreateSourceDir(dir *pb.Dir, destPathName string) {
	name := utils.Joins(destPathName, dir.DirName)

	if !utils.IsDirExist(name) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			log.Debug("Create Dir error: %s", err.Error())
		}
		log.Info("Create Dir: %s", name)
	} else {
		log.Debug("Dir already exist: %s", name)
	}
}

func getRootPathReq(client pb.DirServiceClient, request *pb.GetRootPathReq) *pb.GetRootPathRes {
	response, err := client.GetRootPath(context.Background(), request)
	if err != nil {
		log.Debug("Get Root Path Request error: %s", err.Error())
		return nil
	}
	log.Info("Get Root Path Response: %s", response.String())
	return response
}

func getFileNumberReq(client pb.DirServiceClient, request *pb.GetFileNumberReq) *pb.GetFileNumberRes {
	response, err := client.GetFileNumber(context.Background(), request)
	if err != nil {
		log.Debug("Get File Number Request error: %s", err.Error())
		return nil
	}
	log.Info("Get File Number Response: %s", response.String())
	return response
}

func searchFileReq(client pb.DirServiceClient, request *pb.SearchFileReq) *pb.ListDirRes {
	response, err := client.SearchFile(context.Background(), request)
	if err != nil {
		log.Debug("Search File Request error: %s", err.Error())
		return nil
	}
	log.Info("Search File Response: %s", response.String())
	return response
}
