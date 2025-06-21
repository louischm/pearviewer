package dto

import (
	pb "pearviewer/generated/dir"
)

func CreateUploadDirReq(dirName, pathName string) *pb.UploadDirReq {
	dir := pb.Dir{
		Name: dirName,
	}
	request := &pb.UploadDirReq{
		Dir:      &dir,
		Pathname: pathName,
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
