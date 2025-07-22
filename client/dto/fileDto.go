package dto

import (
	"errors"
	"github.com/louischm/pkg/logger"
	"os"
	"path/filepath"
	pb "pearviewer/generated"
)

var log = logger.NewLog()

func CreateUploadFileReq(fileName string, pathName string) []*pb.UploadFileReq {
	// Fetch Data File
	var uploads []*pb.UploadFileReq
	var startByte int64 = 0

	fi, err := os.Open(fileName)
	if err != nil {
		log.Error("Failed to open file: %s", err.Error())
	}

	defer func() {
		if err = fi.Close(); err != nil {
			log.Error("Failed to close file: %s", err.Error())
		}
	}()

	for {
		file, errFileChunk := getFileChunk(fi, startByte, startByte+1000, pathName)

		if errFileChunk != nil && file != nil {
			log.Info("Empty file.")
			upload := pb.UploadFileReq{
				File:      file,
				StartByte: startByte,
				EndByte:   startByte,
			}
			uploads = append(uploads, &upload)
			break
		} else if errFileChunk != nil {
			log.Info("File read.")
			break
		}

		endByte := startByte + 1000
		upload := pb.UploadFileReq{
			File:      file,
			StartByte: startByte,
			EndByte:   endByte,
		}
		uploads = append(uploads, &upload)
		startByte += 1000
	}
	return uploads
}

func CreateRenameFileReq(oldName, newName, pathName string) *pb.RenameFileReq {
	request := &pb.RenameFileReq{
		OldName:  oldName,
		NewName:  newName,
		PathName: pathName,
	}
	return request
}

func CreateDeleteFileReq(fileName, pathName string) *pb.DeleteFileReq {
	request := &pb.DeleteFileReq{
		FileName: fileName,
		PathName: pathName,
	}
	return request
}

func CreateMoveFileReq(fileName, oldPathName, newPathName string) *pb.MoveFileReq {
	return &pb.MoveFileReq{
		FileName:    fileName,
		OldPathName: oldPathName,
		NewPathName: newPathName,
	}
}

func CreateDownloadFileReq(fileName, pathName string) *pb.DownloadFileReq {
	return &pb.DownloadFileReq{
		FileName: fileName,
		PathName: pathName,
	}
}

func CreateGetFileSizeReq(fileName, pathName string) *pb.GetFileSizeReq {
	return &pb.GetFileSizeReq{
		FileName: fileName,
		PathName: pathName,
	}
}

func getFileChunk(fi *os.File, startByte, endByte int64, pathName string) (*pb.File, error) {
	// Open file

	buf := make([]byte, 1)
	var data []byte
	for {
		n, errReadAt := fi.ReadAt(buf, startByte)
		if n == 0 || startByte == endByte {
			break
		}
		if errReadAt != nil {
			log.Error("Failed to read file: %s", errReadAt.Error())
		}
		data = append(data, buf...)
		startByte++
	}

	if len(data) == 0 && startByte == 0 {
		errData := errors.New("File is empty")
		fname := filepath.Base(fi.Name())
		file := &pb.File{
			Name:     fname,
			Data:     data,
			PathName: pathName,
		}
		return file, errData
	} else if len(data) == 0 {
		errData := errors.New("File is read")
		return nil, errData
	}

	fname := filepath.Base(fi.Name())
	file := pb.File{
		Name:     fname,
		Data:     data,
		PathName: pathName,
	}
	return &file, nil
}
