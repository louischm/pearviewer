package service

import (
	"errors"
	pb "pearviewer/generated"
	res "pearviewer/server/response"
	"pearviewer/server/types"
)

func SignIn(req *pb.SignInReq) (*pb.SignInRes, error) {
	userName := req.User.UserName
	password := req.User.Password

	if userName == "test" && password == "test" {
		return res.CreateSignInRes(types.Success, types.SignInSuccess, nil)
	} else {
		return res.CreateSignInRes(types.Fail, types.SignInError(userName), errors.New(types.SignInError(userName)))
	}
}
