package response

import (
	pb "pearviewer/generated"
)

func CreateDirRes(returnCode int32, message string, err error) (*pb.CreateDirRes, error) {
	return &pb.CreateDirRes{ReturnCode: returnCode, Message: message}, err
}

func CreateRenameDirRes(returnCode int32, message string, err error) (*pb.RenameDirRes, error) {
	return &pb.RenameDirRes{
		ReturnCode: returnCode,
		Message:    message,
	}, err
}

func CreateDeleteDirRes(returnCode int32, message string, err error) (*pb.DeleteDirRes, error) {
	return &pb.DeleteDirRes{
		ReturnCode: returnCode,
		Message:    message,
	}, err
}

func CreateMoveDirRes(returnCode int32, message string, err error) (*pb.MoveDirRes, error) {
	return &pb.MoveDirRes{
		ReturnCode: returnCode,
		Message:    message,
	}, err
}

func CreateListDirRes(returnCode int32, message string, dir *pb.Dir, err error) (*pb.ListDirRes, error) {
	return &pb.ListDirRes{
		ReturnCode: returnCode,
		Message:    message,
		Dir:        dir,
	}, err
}

func CreateGetRootPathRes(returnCode int32, message string, pathName string, err error) (*pb.GetRootPathRes, error) {
	return &pb.GetRootPathRes{ReturnCode: returnCode, Message: message, PathName: pathName}, err
}
