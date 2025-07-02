package grpc

import (
	"context"
	pb "pearviewer/generated"
	"pearviewer/server/service"
	"pearviewer/server/types"
)

func (s *userServer) SignIn(ctx context.Context, request *pb.SignInReq) (*pb.SignInRes, error) {
	log.Info("Received SignIn request: %s", request.String())
	res, err := service.SignIn(request)

	if err != nil && res.GetReturnCode() == types.ServerError {
		return nil, err
	}
	return res, nil
}
