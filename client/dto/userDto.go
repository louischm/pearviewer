package dto

import pb "pearviewer/generated"

func CreateSignInReq(userName, password string) *pb.SignInReq {
	return &pb.SignInReq{
		User: &pb.User{
			UserName: userName,
			Password: password,
		},
	}
}
