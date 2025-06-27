package grpc

import (
	"context"
	"github.com/louischm/pkg/utils"
	"io"
	"os"
	"pearviewer/client/dto"
	"pearviewer/client/types"
	pb "pearviewer/generated"
)

func DownloadFile(fileName, sourcePathName, destPathName string) {
	client, conn := createFileClient()
	defer closeClient(conn)
	request := dto.CreateDownloadFileReq(fileName, sourcePathName)
	log.Info("Download File request created: " + request.String())
	downloadFileReq(*client, request, destPathName)
}

func UploadFile(fileName, pathName string) {
	client, conn := createFileClient()
	defer closeClient(conn)
	request := dto.CreateUploadFileReq(fileName, pathName)
	log.Info("File Data fetched")
	uploadFileReq(*client, request)
}

func RenameFile(oldName, newName, pathName string) {
	client, conn := createFileClient()
	defer closeClient(conn)
	request := dto.CreateRenameFileReq(oldName, newName, pathName)
	log.Info("Rename File request created:" + request.String())
	renameFileReq(*client, request)
}

func DeleteFile(fileName, pathName string) {
	client, conn := createFileClient()
	defer closeClient(conn)
	request := dto.CreateDeleteFileReq(fileName, pathName)
	log.Info("Delete File request created: " + request.String())
	deleteFileReq(*client, request)
}

func MoveFile(fileName, oldPathName, newPathName string) {
	client, conn := createFileClient()
	defer closeClient(conn)
	request := dto.CreateMoveFileReq(fileName, oldPathName, newPathName)
	log.Info("Move File Request created: " + request.String())
	moveFileReq(*client, request)
}

func uploadFileReq(client pb.FileServiceClient, request []*pb.UploadFileReq) {
	// Stream request
	log.Info("Start streaming")
	stream, err := client.UploadFile(context.Background())
	if err != nil {
		log.Fatal("Failed to start streaming" + err.Error())
	}
	waitc := make(chan struct{})

	// Recv responses
	go func() {
		for {
			res, errRecv := stream.Recv()
			if errRecv == io.EOF {
				log.Info("Streaming finished")
				close(waitc)
				return
			}
			if errRecv != nil && res.GetReturnCode() == types.ServerError {
				log.Debug("Failed to receive response: " + errRecv.Error())
			}
			if errRecv != nil {
				log.Debug("Failure during streaming: " + errRecv.Error())
			}
			log.Info("Response received: " + res.String())
		}
	}()

	// Send req
	for _, upload := range request {
		if errStream := stream.Send(upload); errStream != nil {
			log.Debug("Failed to send streaming" + errStream.Error())
		}
		log.Info("Request streaming sended.")
	}
	log.Info("Upload finished")
	_ = stream.CloseSend()
	<-waitc
}

func renameFileReq(client pb.FileServiceClient, request *pb.RenameFileReq) {
	response, err := client.RenameFile(context.Background(), request)
	if err != nil {
		log.Error("Rename File Request error: " + err.Error())
	}
	log.Info("Rename File Response: " + response.String())
}

func deleteFileReq(client pb.FileServiceClient, request *pb.DeleteFileReq) {
	response, err := client.DeleteFile(context.Background(), request)
	if err != nil {
		log.Error("Delete File Request error: " + err.Error())
	}
	log.Info("Delete File response: " + response.String())
}

func moveFileReq(client pb.FileServiceClient, request *pb.MoveFileReq) {
	response, err := client.MoveFile(context.Background(), request)
	if err != nil {
		log.Error("Move File request error: " + err.Error())
	}
	log.Info("Move File response: " + response.String())
}

func downloadFileReq(client pb.FileServiceClient, request *pb.DownloadFileReq, pathName string) {
	stream, err := client.DownloadFile(context.Background(), request)
	if err != nil {
		log.Error("Download File request error: " + err.Error())
	}

	for {
		res, errStream := stream.Recv()
		if errStream == io.EOF {
			log.Info("Download File finished")
			break
		}

		if errStream != nil {
			log.Error("Download File request error: " + errStream.Error())
		}
		log.Info("Response received")
		saveFileChunk(res, pathName)
	}
}

func saveFileChunk(res *pb.DownloadFileRes, pathName string) {
	// Create file if necessary
	if !utils.IsFileInDir(res.File.Name, pathName) {
		utils.CreateEmptyFile(res.File.Name, pathName)
	}
	fileName := utils.Joins(pathName, res.File.Name)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error("Failed to open file: " + fileName)
		return
	}
	if _, err = file.Write(res.File.Data); err != nil {
		log.Error("Failed to write file: " + fileName)
		return
	}
	err = file.Close()
	if err != nil {
		log.Error("Failed to close file: " + fileName)
		return
	}
	log.Info("Chunk File write successfully: " + fileName)
	return
}
