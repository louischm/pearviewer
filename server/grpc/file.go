package grpc

import (
	"context"
	"io"
	pb "pearviewer/generated"
	"pearviewer/server/service"
	"pearviewer/server/types"
)

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
	log.Info("Received RenameFileReq: " + request.String())
	res, err := service.RenameFile(request)

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}

func (s *fileServer) DeleteFile(ctx context.Context, request *pb.DeleteFileReq) (*pb.DeleteFileRes, error) {
	log.Info("Received DeleteFileReq: " + request.String())
	res, err := service.DeleteFile(request)

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}

func (s *fileServer) MoveFile(ctx context.Context, request *pb.MoveFileReq) (*pb.MoveFileRes, error) {
	log.Info("Received MoveFileReq: " + request.String())
	res, err := service.MoveFile(request.GetFileName(), request.GetOldPathName(), request.GetNewPathName())

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}
