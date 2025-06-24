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
