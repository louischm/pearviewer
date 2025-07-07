package service

import (
	"errors"
	"github.com/louischm/pkg/utils"
	"os"
	pb "pearviewer/generated"
	"pearviewer/server/conf"
	"pearviewer/server/db"
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

func CreateUser(req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	connDb, err := db.OpenDB()
	if err != nil {
		log.Debug("Can't connect to database: %v", err)
		return res.CreateUserRes(types.ServerError, types.DBConnectError, err)
	}

	if user := db.GetUserByUserName(connDb, req.User.UserName); user.Username != "" {
		log.Debug("User %s already exists", req.User.UserName)
		return res.CreateUserRes(types.Fail, types.UserAlreadyExists(req.User.UserName),
			errors.New("user already exists"))
	} else {
		db.AddUser(connDb, &db.User{
			Username: req.User.UserName,
			Password: req.User.Password,
		})
		err = createUserDir(req)
		if err != nil {
			return res.CreateUserRes(types.ServerError, types.CreateDirError(req.User.UserName), err)
		}
		return res.CreateUserRes(types.Success, types.UserCreated(req.User.UserName), nil)
	}
}

func createUserDir(req *pb.CreateUserReq) error {
	confData := conf.NewConf()
	dirName := utils.Joins(confData.DataPath, req.User.UserName)
	if !utils.IsDirExist(dirName) {
		log.Info("Creating user directory: %s", dirName)
		if err := os.Mkdir(dirName, os.ModePerm); err != nil {
			log.Debug("Can't create user directory %s: %v", dirName, err)
			return err
		}
	} else {
		log.Info("User directory already exists: %s", dirName)
	}
	return nil
}
