package response

import pb "pearviewer/generated"

func CreateSignInRes(returnCode int32, message string, err error) (*pb.SignInRes, error) {
	res := &pb.SignInRes{
		ReturnCode: returnCode,
		Message:    message,
	}
	return res, err
}
