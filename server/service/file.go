package service

import (
	"errors"
	"github.com/louischm/pkg/utils"
	"io"
	"os"
	pb "pearviewer/generated"
	res "pearviewer/server/response"
	"pearviewer/server/types"
)

func DownloadFileStream(request *pb.DownloadFileReq) ([]*pb.DownloadFileRes, error) {
	// Fetch Data file
	var downloads []*pb.DownloadFileRes
	var startByte int64 = 0
	fileName := utils.Joins(request.PathName, request.FileName)

	fi, err := os.Open(fileName)
	if err != nil {
		log.Error("Failed to open file: %s", fileName)
		return nil, err
	}
	defer func() {
		if err = fi.Close(); err != nil {
			log.Error("Failed to close file: %s", fileName)
		}
	}()

	for {
		file, errFileChunk := getFileChunk(fi, startByte, startByte+1000, request.FileName)

		if errFileChunk != nil && file != nil {
			log.Info("Empty File: %s", fileName)
			// A refractor
			download := &pb.DownloadFileRes{
				/*				ReturnCode: returnCode,
								Message: message,
				*/File:    file,
				StartByte: startByte,
				EndByte:   startByte,
			}
			downloads = append(downloads, download)
			break
		} else if errFileChunk != nil {
			log.Info("File read.")
			break
		}

		endByte := startByte + 1000
		download := &pb.DownloadFileRes{
			File:      file,
			StartByte: startByte,
			EndByte:   endByte,
		}
		downloads = append(downloads, download)
		startByte += 1000
	}
	return downloads, nil
}

func UploadFileChunk(stream pb.FileService_UploadFileServer) (*pb.UploadFileRes, error) {
	var lastByte int64

	upload, err := stream.Recv()
	lastByte = upload.GetEndByte()
	log.Info("Upload file chunk start for: %s", upload.GetFile().GetPathName())
	if err == io.EOF {
		log.Info("File upload EOF")
		return nil, err
	}

	if err != nil {
		return res.CreateUploadFileRes(types.ServerError, "Upload File Stream Read Error", lastByte, err)
	}

	// Create file if necessary
	if !utils.IsFileInDir(upload.File.GetName(), upload.File.PathName) {
		utils.CreateEmptyFile(upload.File.GetName(), upload.File.PathName)
	}
	filename := utils.Joins(upload.File.PathName, upload.File.GetName())
	message, errWrite := writeFileChunk(filename, upload)
	if errWrite != nil {
		return res.CreateUploadFileRes(types.Fail, message, lastByte, errWrite)
	}
	return res.CreateUploadFileRes(types.Success, message, lastByte, nil)
}

func RenameFile(request *pb.RenameFileReq) (*pb.RenameFileRes, error) {
	oldName := utils.Joins(request.GetPathName(), request.GetOldName())
	newName := utils.Joins(request.GetPathName(), request.GetNewName())

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
	name := utils.Joins(request.GetPathName(), request.GetFileName())
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

func GetFileSize(pathName, fileName string) (*pb.GetFileSizeRes, error) {
	name := utils.Joins(pathName, fileName)

	if !utils.IsFileInDir(fileName, pathName) {
		log.Debug(types.FileNotFound(pathName))
		return res.CreateGetFileSizeRes(types.Fail, types.FileNotFound(name), -1, nil)
	}
	fi, err := os.Stat(name)
	if err != nil {
		log.Debug("Error getting file size: %v", err)
		return res.CreateGetFileSizeRes(types.ServerError, "Stat() error", -1, err)
	}
	size := fi.Size()
	return res.CreateGetFileSizeRes(types.Success, types.FileSizeSuccess(name), size, nil)
}

func writeFileChunk(fileName string, upload *pb.UploadFileReq) (string, error) {
	var message string

	// Write file
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Debug(types.OpenFileError(fileName))
		message = types.OpenFileError(fileName)
		return message, err
	}
	if _, err = file.Write(upload.GetFile().GetData()); err != nil {
		log.Debug(types.WriteFileError(fileName))
		message = types.WriteFileError(fileName)
		return message, err
	}
	err = file.Close()
	if err != nil {
		log.Debug(types.CloseFileError(fileName))
		message = types.CloseFileError(fileName)
		return message, err
	}
	log.Info(types.WriteFileChunkSuccess(fileName))
	return types.WriteFileChunkSuccess(fileName), nil
}

func getFileChunk(fi *os.File, startByte, endByte int64, fileName string) (*pb.File, error) {
	buf := make([]byte, 1)
	var data []byte
	for {
		n, err := fi.ReadAt(buf, startByte)
		if n == 0 || startByte == endByte {
			break
		}
		if err != nil {
			log.Error("Failed to read file: %s", err.Error())
		}

		data = append(data, buf...)
		startByte++
	}

	if len(data) == 0 && startByte == 0 {
		err := errors.New("File is empty")
		file := &pb.File{
			Name: fileName,
			Data: data,
		}
		return file, err
	} else if len(data) == 0 {
		err := errors.New("File is read")
		return nil, err
	}

	file := &pb.File{
		Name: fileName,
		Data: data,
	}
	return file, nil
}
