package dto

import (
	"pearviewer/client/tree"
	"pearviewer/client/types"
	pb "pearviewer/generated"
)

func CreateUploadDirTree(dirName, oldPathName, newPathName string) *types.Dir {
	// Create dir tree
	data, err := tree.CreateTree(dirName, oldPathName, newPathName)
	if err != nil {
		log.Info(err.Error())
	}
	return data
}

func CreateDirReq(dirName, pathName string) *pb.CreateDirReq {
	request := &pb.CreateDirReq{
		DirName:  dirName,
		PathName: pathName,
	}
	return request
}

func CreateRenameDirReq(oldName, newName, pathName string) *pb.RenameDirReq {
	return &pb.RenameDirReq{
		OldName:  oldName,
		NewName:  newName,
		PathName: pathName,
	}
}

func CreateDeleteDirReq(dirName, pathName string) *pb.DeleteDirReq {
	return &pb.DeleteDirReq{
		DirName:  dirName,
		PathName: pathName,
	}
}

func CreateMoveDirReq(dirName, oldPathName, newPathName string) *pb.MoveDirReq {
	return &pb.MoveDirReq{
		DirName:     dirName,
		OldPathName: oldPathName,
		NewPathName: newPathName,
	}
}

func CreateListDirReq(dirName, pathName string) *pb.ListDirReq {
	return &pb.ListDirReq{
		DirName:  dirName,
		PathName: pathName,
	}
}

func CreateGetRootPathReq(username string) *pb.GetRootPathReq {
	return &pb.GetRootPathReq{
		UserName: username,
	}
}

func CreateGetFileNumberReq(dirname, pathname string) *pb.GetFileNumberReq {
	return &pb.GetFileNumberReq{
		DirName:  dirname,
		PathName: pathname,
	}
}

func CreateFileSearchReq(search, pathName, dirName string) *pb.SearchFileReq {
	return &pb.SearchFileReq{
		Search:   search,
		PathName: pathName,
		DirName:  dirName,
	}
}
