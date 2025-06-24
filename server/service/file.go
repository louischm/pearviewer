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
			return res.CreateRenameFileRes(types.ServerError, "Error while renaming "+oldName+" to "+newName, err)
		}
		return res.CreateRenameFileRes(types.Success, "File Renamed "+oldName+" to "+newName, nil)
	} else {
		return res.CreateRenameFileRes(types.Fail, "File does not exist: "+oldName, errors.New("File does not exist: "+newName))
	}
}

func DeleteFile(request *pb.DeleteFileReq) (*pb.DeleteFileRes, error) {
	name := request.GetPathName() + request.GetFileName()
	if utils.IsFileInDir(request.GetFileName(), request.GetPathName()) {
		err := os.Remove(name)
		if err != nil {
			return res.CreateDeleteFileRes(types.ServerError, "Error while deleting "+name, err)
		}
		return res.CreateDeleteFileRes(types.Success, "File Deleted: "+name, nil)
	} else {
		return res.CreateDeleteFileRes(types.Fail, "File does not exist: "+name,
			errors.New("File does not exist: "+name))
	}
}

func MoveFile(fileName, oldPathName, newPathName string) (*pb.MoveFileRes, error) {
	oldName := utils.Joins(oldPathName, fileName)
	newName := utils.Joins(newPathName, fileName)
	if !utils.IsFileInDir(fileName, oldPathName) {
		return res.CreateMoveFileRes(types.Fail, "File does not exist: "+oldName,
			errors.New("File does not exist: "+fileName))
	} else {
		iFile, err := os.Open(oldName)
		if err != nil {
			return res.CreateMoveFileRes(types.Fail, "Error while opening file: "+oldName, err)
		}
		defer iFile.Close()

		oFile, errCreate := os.Create(newName)
		if errCreate != nil {
			return res.CreateMoveFileRes(types.Fail, "Error while creating file: "+newName, errCreate)
		}
		defer oFile.Close()

		_, err = io.Copy(oFile, iFile)
		if err != nil {
			return res.CreateMoveFileRes(types.Fail, "Error while copying file: "+newName, err)
		}
		iFile.Close()

		err = os.Remove(oldName)
		if err != nil {
			return res.CreateMoveFileRes(types.Fail, "Error while removing file: "+oldName, err)
		}
		return res.CreateMoveFileRes(types.Success, "File Moved: "+newName, nil)
	}
}

func writeFileChunk(fileName string, upload *pb.UploadFileReq) (string, error) {
	var message string

	// Write file
	file, errOpen := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if errOpen != nil {
		message = "Failed to open File: " + fileName
		return message, errOpen
	}
	if _, errWrite := file.Write(upload.GetFile().GetData()); errWrite != nil {
		message = "Failed to write to File: " + fileName
		return message, errWrite
	}
	errClose := file.Close()
	if errClose != nil {
		message = "Failed to close File: " + fileName
		return message, errClose
	}
	return "WriteFileChunk success: " + fileName, nil
}
