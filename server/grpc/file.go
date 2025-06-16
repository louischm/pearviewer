package grpc

import (
	"io"
	"os"
	. "pearviewer/server/generated/file"
)

func (s *fileServer) UploadFile(stream FileService_UploadFileServer) error {
	var returnCode int32
	var lastByte int64
	var message string

	for {
		upload, err := stream.Recv()
		if err == io.EOF {
			log.Info("Stream EOF")
			message = "Stream EOF"
			lastByte = upload.GetEndByte()
			return stream.SendAndClose(&UploadFileRes{
				ReturnCode: returnCode,
				LastByte:   lastByte,
				Message:    message,
			})
		}

		if err != nil {
			log.Error(err.Error())
		}

		// Create file if necessary
		if !isFileInDir(upload.File.GetName(), upload.GetPathname()) {
			createEmptyFile(upload.File.GetName(), upload.GetPathname())
		}

		filename := upload.GetPathname() + upload.File.GetName()

		// Write file
		file, errOpen := os.OpenFile(filename, os.O_WRONLY, 0666)
		if errOpen != nil {
			log.Error(errOpen.Error())
		}
		if _, errWrite := file.Write(upload.GetFile().GetData()); errWrite != nil {
			log.Error(errWrite.Error())
		}
		errClose := file.Close()
		if errClose != nil {
			log.Error(errClose.Error())
		}

	}
}

func createEmptyFile(fileName, dirName string) {
	_, err := os.OpenFile(dirName+fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Error(err.Error())
	}
}

func isFileInDir(fileName, dirName string) bool {
	entries, err := os.ReadDir(dirName)
	if err != nil {
		log.Error(err.Error())
	}

	for _, entry := range entries {
		if entry.Name() == fileName {
			return true
		}
	}
	return false
}
