package service

import (
	"errors"
	"github.com/louischm/logger"
	"io"
	"os"
	pb "pearviewer/generated"
	res "pearviewer/server/response"
	"pearviewer/server/types"
	"pearviewer/server/utils"
)

var log = logger.NewLog()

func UploadFileChunk(stream pb.FileService_UploadFileServer) (*pb.UploadFileRes, error) {
	var lastByte int64

	upload, err := stream.Recv()
	lastByte = upload.GetEndByte()
	log.Info("Upload file chunk start for: " + upload.GetPathName())
	if err == io.EOF {
		log.Info("File upload EOF")
		return nil, err
	}

	if err != nil {
		return res.CreateUploadFileRes(types.ServerError, "Upload File Stream Read Error", lastByte, err)
	}

	// Create file if necessary
	if !utils.IsFileInDir(upload.File.GetName(), upload.GetPathName()) {
		utils.CreateEmptyFile(upload.File.GetName(), upload.GetPathName())
	}
	filename := utils.Joins(upload.GetPathName(), upload.File.GetName())
	message, errWrite := writeFileChunk(filename, upload)
	if errWrite != nil {
		return res.CreateUploadFileRes(types.Fail, message, lastByte, errWrite)
	}
	return res.CreateUploadFileRes(types.Success, message, lastByte, nil)
}

func RenameFile(request *pb.RenameFileReq) (*pb.RenameFileRes, error) {
	oldName := request.GetPathName() + request.GetOldName()
	newName := request.GetPathName() + request.GetNewName()

	if utils.IsFileInDir(request.GetOldName(), request.GetPathName()) {
		err := os.Rename(oldName, newName)
		if err != nil {
			log.Debug(types.RenameFileError(oldName, newName))
			return res.CreateRenameFileRes(types.ServerError, types.RenameFileError(oldName, newName), err)
		}
		log.Info(types.RenameFileSuccess(oldName, newName))
		return res.CreateRenameFileRes(types.Success, types.RenameFileSuccess(oldName, newName), nil)
	} else {
		log.Debug(types.FileNotFound(oldName))
		return res.CreateRenameFileRes(types.Fail, types.FileNotFound(oldName),
			errors.New(types.FileNotFound(oldName)))
	}
}

func DeleteFile(request *pb.DeleteFileReq) (*pb.DeleteFileRes, error) {
	name := request.GetPathName() + request.GetFileName()
	if utils.IsFileInDir(request.GetFileName(), request.GetPathName()) {
		err := os.Remove(name)
		if err != nil {
			log.Debug(types.DeleteFileError(name))
			return res.CreateDeleteFileRes(types.ServerError, types.DeleteFileSuccess(name), err)
		}
		log.Info(types.DeleteFileSuccess(name))
		return res.CreateDeleteFileRes(types.Success, types.DeleteFileSuccess(name), nil)
	} else {
		log.Debug(types.FileNotFound(name))
		return res.CreateDeleteFileRes(types.Fail, types.FileNotFound(name),
			errors.New(types.FileNotFound(name)))
	}
}

func MoveFile(fileName, oldPathName, newPathName string) (*pb.MoveFileRes, error) {
	oldName := utils.Joins(oldPathName, fileName)
	newName := utils.Joins(newPathName, fileName)
	if !utils.IsFileInDir(fileName, oldPathName) {
		log.Debug(types.FileNotFound(oldName))
		return res.CreateMoveFileRes(types.Fail, types.FileNotFound(oldName),
			errors.New(types.FileNotFound(oldName)))
	} else {
		iFile, err := os.Open(oldName)
		if err != nil {
			log.Debug(types.OpenFileError(oldName))
			return res.CreateMoveFileRes(types.Fail, types.OpenFileError(oldName), err)
		}
		defer iFile.Close()

		oFile, errCreate := os.Create(newName)
		if errCreate != nil {
			log.Debug(types.CreateFileError(newName))
			return res.CreateMoveFileRes(types.Fail, types.CreateFileError(newName), errCreate)
		}
		defer oFile.Close()

		_, err = io.Copy(oFile, iFile)
		if err != nil {
			log.Debug(types.CopyFileError(newName))
			return res.CreateMoveFileRes(types.Fail, types.CopyFileError(newName), err)
		}
		iFile.Close()

		err = os.Remove(oldName)
		if err != nil {
			log.Debug(types.DeleteFileError(oldName))
			return res.CreateMoveFileRes(types.Fail, types.DeleteFileError(oldName), err)
		}
		log.Info(types.MoveFileSuccess(newName))
		return res.CreateMoveFileRes(types.Success, types.MoveFileSuccess(newName), nil)
	}
}

func writeFileChunk(fileName string, upload *pb.UploadFileReq) (string, error) {
	var message string

	// Write file
	file, errOpen := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if errOpen != nil {
		log.Debug(types.OpenFileError(fileName))
		message = types.OpenFileError(fileName)
		return message, errOpen
	}
	if _, errWrite := file.Write(upload.GetFile().GetData()); errWrite != nil {
		log.Debug(types.WriteFileError(fileName))
		message = types.WriteFileError(fileName)
		return message, errWrite
	}
	errClose := file.Close()
	if errClose != nil {
		log.Debug(types.CloseFileError(fileName))
		message = types.CloseFileError(fileName)
		return message, errClose
	}
	log.Info(types.WriteFileChunkSuccess(fileName))
	return types.WriteFileChunkSuccess(fileName), nil
}
