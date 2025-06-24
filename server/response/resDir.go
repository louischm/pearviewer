package response

import (
	pb "pearviewer/generated"
)

func CreateDirRes(returnCode int32, message string, err error) (*pb.CreateDirRes, error) {
	if err != nil {
		message += ": " + err.Error()
	} else {
		log.Info(message)
	}
	return &pb.CreateDirRes{ReturnCode: returnCode, Message: message}, err
}

func CreateRenameDirRes(returnCode int32, message string, err error) (*pb.RenameDirRes, error) {
	if err != nil {
		message += ": " + err.Error()
		log.Debug(message)
	} else {
		log.Info(message)
	}
	return &pb.RenameDirRes{
		ReturnCode: returnCode,
		Message:    message,
	}, err
}

func CreateDeleteDirRes(returnCode int32, message string, err error) (*pb.DeleteDirRes, error) {
	if err != nil {
		message += ": " + err.Error()
		log.Debug(message)
	} else {
		log.Info(message)
	}
	return &pb.DeleteDirRes{
		ReturnCode: returnCode,
		Message:    message,
	}, err
}

func CreateMoveDirRes(returnCode int32, message string, err error) (*pb.MoveDirRes, error) {
	if err != nil {
		message += ": " + err.Error()
		log.Debug(message)
	} else {
		log.Info(message)
	}
	return &pb.MoveDirRes{
		ReturnCode: returnCode,
		Message:    message,
	}, err
}
