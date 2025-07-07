package response

import pb "pearviewer/generated"

func CreateSignInRes(returnCode int32, message string, err error) (*pb.SignInRes, error) {
	res := &pb.SignInRes{
		ReturnCode: returnCode,
		Message:    message,
	}
	return res, err
}

func CreateUserRes(returnCode int32, message string, err error) (*pb.CreateUserRes, error) {
	res := &pb.CreateUserRes{
		ReturnCode: returnCode,
		Message:    message,
	}
	return res, err
}
