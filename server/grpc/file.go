package grpc

import (
	"context"
	"github.com/louischm/pkg/utils"
	"io"
	pb "pearviewer/generated"
	"pearviewer/server/service"
	"pearviewer/server/types"
)

func (s *fileServer) DownloadFile(request *pb.DownloadFileReq, stream pb.FileService_DownloadFileServer) error {
	name := utils.Joins(request.PathName, request.FileName)
	log.Info("Start Download File: %s", name)
	uploads, err := service.DownloadFileStream(request)

	for _, upload := range uploads {
		if err = stream.Send(upload); err != nil {
			return err
		}
	}
	return nil
}

func (s *fileServer) UploadFile(stream pb.FileService_UploadFileServer) error {
	log.Info("New Upload File Streaming started.")
	for {
		res, err := service.UploadFileChunk(stream)

		// End Of Stream
		if err == io.EOF {
			return nil
		}

		// Server Error type
		if err != nil && res.GetReturnCode() == types.ServerError {
			return err
		}

		// Failure Error type
		if err != nil {
			if errSend := stream.Send(res); errSend != nil {
				return errSend
			}
		}

		// Success
		if errSend := stream.Send(res); errSend != nil {
			return errSend
		}
	}
}

func (s *fileServer) RenameFile(ctx context.Context, request *pb.RenameFileReq) (*pb.RenameFileRes, error) {
	log.Info("Received RenameFileReq: %s", request.String())
	res, err := service.RenameFile(request)

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}

func (s *fileServer) DeleteFile(ctx context.Context, request *pb.DeleteFileReq) (*pb.DeleteFileRes, error) {
	log.Info("Received DeleteFileReq: %s", request.String())
	res, err := service.DeleteFile(request)

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}

func (s *fileServer) MoveFile(ctx context.Context, request *pb.MoveFileReq) (*pb.MoveFileRes, error) {
	log.Info("Received MoveFileReq: %s", request.String())
	res, err := service.MoveFile(request.GetFileName(), request.GetOldPathName(), request.GetNewPathName())

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}

func (s *fileServer) GetFileSize(ctx context.Context, request *pb.GetFileSizeReq) (*pb.GetFileSizeRes, error) {
	log.Info("Received GetFileSizeReq: %s", request.String())
	res, err := service.GetFileSize(request.GetPathName(), request.GetFileName())
	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}
