package grpc

import (
	"context"
	"errors"
	"pearviewer/client/dto"
	"pearviewer/client/types"
	pb "pearviewer/generated"
)

func SignIn(userName, password string) (*pb.SignInRes, error) {
	client, conn := createUserClient()
	defer closeClient(conn)
	request := dto.CreateSignInReq(userName, password)
	log.Info("SignIn Request created: %s", request.User.UserName)
	return signInReq(*client, request)
}

func CreateUser(userName, password string) (*pb.CreateUserRes, error) {
	client, conn := createUserClient()
	defer closeClient(conn)
	request := dto.CreateUserReq(userName, password)
	log.Info("CreateUser request created: %s", request.User.UserName)
	return createUserReq(*client, request)
}

func signInReq(client pb.UserServiceClient, request *pb.SignInReq) (*pb.SignInRes, error) {
	response, err := client.SignIn(context.Background(), request)
	if err != nil {
		log.Debug(err.Error())
		return response, err
	} else if response.ReturnCode != types.Success {
		return response, errors.New(response.Message)
	}
	log.Info("SignIn Response: %s", response.String())
	return response, nil
}

func createUserReq(client pb.UserServiceClient, request *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	response, err := client.CreateUser(context.Background(), request)
	if err != nil {
		log.Debug(err.Error())
		return response, err
	}
	log.Info("CreateUser Response: %s", response.String())
	return response, nil
}
