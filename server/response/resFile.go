package response

import (
	"github.com/louischm/logger"
	"io"
	pb "pearviewer/generated/file"
)

var log = logger.NewLog()

func CreateUploadFileRes(returnCode int32, message string, lastByte int64, err error) (*pb.UploadFileRes, error) {
	if err != nil && err != io.EOF {
		message += ": " + err.Error()
		log.Debug(message)
	} else {
		log.Info(message)
	}
	res := &pb.UploadFileRes{
		Message:    message,
		ReturnCode: returnCode,
		LastByte:   lastByte,
	}
	return res, err
}

func CreateRenameFileRes(returnCode int32, message string, err error) (*pb.RenameFileRes, error) {
	if err != nil {
		message += ": " + err.Error()
		log.Debug(message)
	} else {
		log.Info(message)
	}
	res := &pb.RenameFileRes{
		Message:    message,
		ReturnCode: returnCode,
	}
	return res, err
}

func CreateDeleteFileRes(returnCode int32, message string, err error) (*pb.DeleteFileRes, error) {
	if err != nil {
		message += ": " + err.Error()
		log.Debug(message)
	} else {
		log.Info(message)
	}

	res := &pb.DeleteFileRes{
		Message:    message,
		ReturnCode: returnCode,
	}
	return res, err
}

func CreateMoveFileRes(returnCode int32, message string, err error) (*pb.MoveFileRes, error) {
	if err != nil {
		message += ": " + err.Error()
		log.Debug(message)
	} else {
		log.Info(message)
	}

	res := &pb.MoveFileRes{
		Message:    message,
		ReturnCode: returnCode,
	}
	return res, err
}
