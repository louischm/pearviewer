package grpc

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	pb "pearviewer/client/generated/file"
)

func RenameFile(oldName, newName, pathName string) {
	client, conn := createFileClient()
	request := createRenameFileReq(oldName, newName, pathName)
	log.Info("Rename File request created:" + request.String())
	renameFileReq(*client, request)
	closeClient(conn)
}

func createRenameFileReq(oldName, newName, pathName string) *pb.RenameFileReq {
	request := &pb.RenameFileReq{
		OldName:  oldName,
		NewName:  newName,
		PathName: pathName,
	}
	return request
}

func renameFileReq(client pb.FileServiceClient, request *pb.RenameFileReq) {
	response, err := client.RenameFile(context.Background(), request)
	if err != nil {
		log.Error("Rename File Request error: " + err.Error())
	}
	log.Info("Rename File Response: " + response.String())
}

func uploadFile(fileName string, pathName string) {
	client, conn := createFileClient()
	request := createUploadFileReq(fileName, pathName)
	log.Info("File Data fetched")
	uploadFileReq(*client, request)
	closeClient(conn)
}

func createUploadFileReq(fileName string, pathName string) []*pb.UploadFileReq {
	// Fetch Data File
	var uploads []*pb.UploadFileReq
	var startByte int64 = 0

	fi, err := os.Open(fileName)
	if err != nil {
		log.Error("Failed to open file" + err.Error())
	}

	defer func() {
		if errClose := fi.Close(); errClose != nil {
			log.Error("Failed to close file" + errClose.Error())
		}
	}()

	for {
		file, errFileChunk := getFileChunk(fi, startByte)

		if errFileChunk != nil {
			log.Info("File read.")
			break
		}

		endByte := startByte + 1000
		upload := pb.UploadFileReq{
			File:      file,
			StartByte: startByte,
			EndByte:   endByte,
			PathName:  pathName,
		}
		uploads = append(uploads, &upload)
		startByte += 1000
	}
	return uploads
}

func uploadFileReq(client pb.FileServiceClient, request []*pb.UploadFileReq) {
	// Stream request
	log.Info("Start streaming")
	stream, err := client.UploadFile(context.Background())
	if err != nil {
		log.Fatal("Failed to start streaming" + err.Error())
	}

	for _, upload := range request {
		if errStream := stream.Send(upload); errStream != nil {
			log.Fatal("Failed to send streaming" + errStream.Error())
		}
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Failed to receive streaming" + err.Error())
	}
	log.Info("Upload finished" + response.GetMessage())
}

func getFileChunk(fi *os.File, startByte int64) (*pb.File, error) {
	// Open file

	buf := make([]byte, 1)
	var data []byte
	for {
		n, errReadAt := fi.ReadAt(buf, startByte)
		if n == 0 {
			break
		}
		if errReadAt != nil {
			log.Error("Failed to read file" + errReadAt.Error())
		}
		data = append(data, buf...)
		startByte++
	}

	if len(data) == 0 {
		errData := errors.New("File is empty")
		return nil, errData
	}
	fname := filepath.Base(fi.Name())
	file := pb.File{
		Name: fname,
		Data: data,
	}
	return &file, nil
}
