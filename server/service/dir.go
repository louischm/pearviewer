package service

import (
	"errors"
	"github.com/louischm/pkg/logger"
	"github.com/louischm/pkg/utils"
	"os"
	pb "pearviewer/generated"
	"pearviewer/server/conf"
	res "pearviewer/server/response"
	"pearviewer/server/types"
	"strings"
)

var log = logger.NewLog()

func CreateDir(dirName, pathName string) (*pb.CreateDirRes, error) {
	name := utils.Joins(pathName, dirName)
	if !utils.IsDirExist(name) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			log.Debug(types.CreateDirSuccess(name))
			return res.CreateDirRes(types.ServerError, types.CreateDirError(name), err)
		}
		log.Info(types.CreateDirSuccess(name))
		return res.CreateDirRes(types.Success, types.CreateDirSuccess(name), nil)
	} else {
		log.Debug(types.DirAlreadyExists(name))
		return res.CreateDirRes(types.Fail, types.DirAlreadyExists(name),
			errors.New(types.DirAlreadyExists(name)))
	}
}

func RenameDir(request *pb.RenameDirReq) (*pb.RenameDirRes, error) {
	oldName := request.GetPathName() + request.GetOldName()
	newName := request.GetPathName() + request.GetNewName()

	if utils.IsDirExist(oldName) {
		err := os.Rename(oldName, newName)
		if err != nil {
			log.Debug(types.RenameDirError(oldName))
			return res.CreateRenameDirRes(types.Fail, types.RenameDirError(oldName), err)
		}
		log.Info(types.RenameDirSuccess(oldName, newName))
		return res.CreateRenameDirRes(types.Success, types.RenameDirSuccess(oldName, newName), nil)
	} else {
		log.Debug(types.DirNotFound(oldName))
		return res.CreateRenameDirRes(types.Fail, types.DirNotFound(oldName),
			errors.New(types.DirNotFound(oldName)))
	}
}

func DeleteDir(dirName, pathName string) (*pb.DeleteDirRes, error) {
	name := utils.Joins(pathName, dirName)

	if utils.IsDirExist(name) {
		err := os.Remove(name)
		if err != nil {
			log.Debug(types.DeleteDirError(name))
			return res.CreateDeleteDirRes(types.Fail, types.DeleteDirError(name), err)
		}
		log.Info(types.DeleteDirSuccess(name))
		return res.CreateDeleteDirRes(types.Success, types.DeleteDirSuccess(name), nil)
	} else {
		log.Debug(types.DirNotFound(name))
		return res.CreateDeleteDirRes(types.Fail, types.DirNotFound(name),
			errors.New(types.DirNotFound(name)))
	}
}

func MoveDir(dirName, oldPathName, newPathName string) (*pb.MoveDirRes, error) {
	oldName := utils.Joins(oldPathName, dirName)
	newName := utils.Joins(newPathName, dirName)
	if utils.IsDirExist(oldName) {
		files, err := os.ReadDir(oldName)
		if err != nil {
			log.Debug(types.ReadDirError(oldName))
			return res.CreateMoveDirRes(types.Fail, types.ReadDirError(oldName), err)
		}

		// Create dir
		resCreateDir, err := CreateDir(dirName, newPathName)
		if err != nil && resCreateDir.GetReturnCode() != types.ServerError {
			log.Debug(types.CreateDirError(oldName))
			return res.CreateMoveDirRes(types.Fail, types.CreateDirError(oldName), err)
		}
		for _, file := range files {
			if file.IsDir() {
				// Move child dir
				MoveDir(file.Name(), oldName, newName)
			} else {
				// Move files
				_, err = MoveFile(file.Name(), oldName, newName)
				if err != nil {
					log.Debug(types.FileMoveError(file.Name()))
					return res.CreateMoveDirRes(types.Fail, types.FileMoveError(file.Name()), err)
				}
			}
		}
		log.Info(types.DirMoveSuccess(newName))
		return res.CreateMoveDirRes(types.Success, types.DirMoveSuccess(newName), nil)
	} else {
		log.Debug(types.DirNotFound(oldName))
		return res.CreateMoveDirRes(types.Fail, types.DirNotFound(oldName),
			errors.New(types.DirNotFound(oldName)))
	}
}

func ListDir(request *pb.ListDirReq) (*pb.ListDirRes, error) {
	name := utils.Joins(request.GetPathName(), request.GetDirName())

	if utils.IsDirExist(name) {
		dir, err := createListDir(name, request.GetDirName(), request.GetPathName())
		if err != nil {
			log.Debug(types.ListDirError)
			return res.CreateListDirRes(types.ServerError, types.ListDirError, nil, err)
		}
		log.Info(types.ListDirSuccess)
		return res.CreateListDirRes(types.Success, types.ListDirSuccess, dir, nil)
	} else {
		log.Debug(types.DirNotFound(name))
		return res.CreateListDirRes(types.Fail, types.DirNotFound(name), nil,
			errors.New(types.DirNotFound(name)))
	}
}

func GetRootPath(username string) (*pb.GetRootPathRes, error) {
	confData := conf.NewConf()
	pathName := utils.Joins(confData.DataPath, username)

	if !utils.IsDirExist(pathName) {
		log.Info(types.DirNotFound(pathName))
		if err := os.Mkdir(pathName, os.ModePerm); err != nil {
			log.Info(types.CreateDirError(pathName))
			return res.CreateGetRootPathRes(types.Fail, types.CreateDirError(pathName), pathName, err)
		}
	}
	return res.CreateGetRootPathRes(types.Success, types.GetRootPathSuccess(username), pathName, nil)
}

func createFullName(name, fileName string) string {
	confData := conf.NewConf()
	rootName := utils.Joins(confData.DataPath, fileName)

	if strings.HasPrefix(name, rootName) {
		return strings.Replace(name, rootName, "", 1)
	}
	return name
}

func createListDir(name, dirName, pathName string) (*pb.Dir, error) {
	dir := &pb.Dir{
		DirName:  dirName,
		PathName: pathName,
		Type:     pb.Type_DirType,
		FullName: createFullName(name, dirName),
	}

	files, err := os.ReadDir(name)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			newName := utils.Joins(name, file.Name())
			newDir, errCreate := createListDir(newName, file.Name(), name)
			if errCreate != nil {
				return nil, errCreate
			}
			dir.Dir = append(dir.GetDir(), newDir)
		} else {
			newFile := &pb.File{
				Name:     file.Name(),
				PathName: name,
				Type:     pb.Type_FileType,
				FullName: createFullName(utils.Joins(name, file.Name()), file.Name()),
			}
			dir.File = append(dir.GetFile(), newFile)
		}
	}
	return dir, nil
}
