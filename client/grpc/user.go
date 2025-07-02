package grpc

import (
	"context"
	"pearviewer/client/dto"
	pb "pearviewer/generated"
)

func SignIn(userName, password string) (*pb.SignInRes, error) {
	client, conn := createUserClient()
	defer closeClient(conn)
	request := dto.CreateSignInReq(userName, password)
	log.Info("SignIn Request created: %s", request.String())
	return signInReq(*client, request)
}

func signInReq(client pb.UserServiceClient, request *pb.SignInReq) (*pb.SignInRes, error) {
	response, err := client.SignIn(context.Background(), request)
	if err != nil {
		log.Debug(err.Error())
		return response, err
	}
	log.Info("SignIn Response: %s", response.String())
	return response, nil
}
