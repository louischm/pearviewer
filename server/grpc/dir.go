package grpc

import (
	"context"
	"github.com/louischm/pkg/utils"
	"os"
	pb "pearviewer/generated"
	"pearviewer/server/service"
	"pearviewer/server/types"
)

func (s *dirServer) CreateDir(ctx context.Context, request *pb.CreateDirReq) (*pb.CreateDirRes, error) {
	log.Info("Received CreateDirReq: %s", request.String())
	res, err := service.CreateDir(request.GetDirName(), request.GetPathName())

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}

func (s *dirServer) RenameDir(ctx context.Context, request *pb.RenameDirReq) (*pb.RenameDirRes, error) {
	log.Info("Received RenameDirReq: %s", request.String())
	res, err := service.RenameDir(request)

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}

func (s *dirServer) DeleteDir(ctx context.Context, request *pb.DeleteDirReq) (*pb.DeleteDirRes, error) {
	log.Info("Received DeleteDirReq: %s", request.String())
	res, err := service.DeleteDir(request.GetDirName(), request.GetPathName())

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}

func (s *dirServer) MoveDir(ctx context.Context, request *pb.MoveDirReq) (*pb.MoveDirRes, error) {
	log.Info("Received MoveDirReq: %s", request.String())
	res, err := service.MoveDir(request.GetDirName(), request.GetOldPathName(), request.GetNewPathName())

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	// Delete oldDir
	oldName := utils.Joins(request.GetOldPathName(), request.GetDirName())
	err = os.RemoveAll(oldName)
	if err != nil {
		log.Info("DeleteDir failed: %s", err.Error())
	}
	return res, nil
}

func (s *dirServer) ListDir(ctx context.Context, request *pb.ListDirReq) (*pb.ListDirRes, error) {
	log.Info("Received ListDirReq: %s", request.String())
	res, err := service.ListDir(request)

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}

func (s *dirServer) GetRootPath(ctx context.Context, request *pb.GetRootPathReq) (*pb.GetRootPathRes, error) {
	log.Info("Received GetRootPathReq: %s", request.String())
	res, err := service.GetRootPath(request.UserName)

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}
