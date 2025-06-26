package response

import (
	pb "pearviewer/generated"
)

func CreateUploadFileRes(returnCode int32, message string, lastByte int64, err error) (*pb.UploadFileRes, error) {
	res := &pb.UploadFileRes{
		Message:    message,
		ReturnCode: returnCode,
		LastByte:   lastByte,
	}
	return res, err
}

func CreateRenameFileRes(returnCode int32, message string, err error) (*pb.RenameFileRes, error) {
	res := &pb.RenameFileRes{
		Message:    message,
		ReturnCode: returnCode,
	}
	return res, err
}

func CreateDeleteFileRes(returnCode int32, message string, err error) (*pb.DeleteFileRes, error) {
	res := &pb.DeleteFileRes{
		Message:    message,
		ReturnCode: returnCode,
	}
	return res, err
}

func CreateMoveFileRes(returnCode int32, message string, err error) (*pb.MoveFileRes, error) {
	res := &pb.MoveFileRes{
		Message:    message,
		ReturnCode: returnCode,
	}
	return res, err
}
