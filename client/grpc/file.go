package grpc

import (
	"context"
	"io"
	"pearviewer/client/dto"
	"pearviewer/client/types"
	pb "pearviewer/generated/file"
)

func UploadFile(fileName string, pathName string) {
	client, conn := createFileClient()
	request := dto.CreateUploadFileReq(fileName, pathName)
	log.Info("File Data fetched")
	uploadFileReq(*client, request)
	closeClient(conn)
}

func RenameFile(oldName, newName, pathName string) {
	client, conn := createFileClient()
	request := dto.CreateRenameFileReq(oldName, newName, pathName)
	log.Info("Rename File request created:" + request.String())
	renameFileReq(*client, request)
	closeClient(conn)
}

func DeleteFile(fileName, pathName string) {
	client, conn := createFileClient()
	request := dto.CreateDeleteFileReq(fileName, pathName)
	log.Info("Delete File request created: " + request.String())
	deleteFileReq(*client, request)
	closeClient(conn)
}

func MoveFile(fileName, oldPathName, newPathName string) {
	client, conn := createFileClient()
	request := dto.CreateMoveFileReq(fileName, oldPathName, newPathName)
	log.Info("Move File Request created: " + request.String())
	moveFileReq(*client, request)
	closeClient(conn)
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
