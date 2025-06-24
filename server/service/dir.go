package service

import (
	"errors"
	"os"
	pb "pearviewer/generated"
	res "pearviewer/server/response"
	"pearviewer/server/types"
	"pearviewer/server/utils"
)

func CreateDir(dirName, pathName string) (*pb.CreateDirRes, error) {
	name := utils.Joins(pathName, dirName)
	if !utils.IsDirExist(name) {
		if err := os.Mkdir(name, os.ModePerm); err != nil {
			return res.CreateDirRes(types.ServerError, "Error creating directory "+name, err)
		}
		return res.CreateDirRes(types.Success, "Directory Created: "+name, nil)
	} else {
		return res.CreateDirRes(types.Fail, "Directory Already Exists",
			errors.New("Directory Already Exists: "+name))
	}
}

func RenameDir(request *pb.RenameDirReq) (*pb.RenameDirRes, error) {
	oldName := request.GetPathName() + request.GetOldName()
	newName := request.GetPathName() + request.GetNewName()

	if utils.IsDirExist(oldName) {
		err := os.Rename(oldName, newName)
		if err != nil {
			return res.CreateRenameDirRes(types.Fail, "Failed to rename directory: "+oldName, err)
		}
		return res.CreateRenameDirRes(types.Success, "Directory Renamed: "+oldName+" to "+newName, nil)
	} else {
		return res.CreateRenameDirRes(types.Fail, "Directory: "+oldName+" doesn't exists",
			errors.New("Directory: "+oldName+" doesn't exists"))
	}
}

func DeleteDir(dirName, pathName string) (*pb.DeleteDirRes, error) {
	name := utils.Joins(pathName, dirName)

	if utils.IsDirExist(name) {
		err := os.Remove(name)
		if err != nil {
			return res.CreateDeleteDirRes(types.Fail, "Dir: "+name+" not deleted", err)
		}
		return res.CreateDeleteDirRes(types.Success, "Directory Deleted: "+name, nil)
	} else {
		return res.CreateDeleteDirRes(types.Fail, "Directory Not Found: "+name,
			errors.New("Directory Not Found: "+name))
	}
}

func MoveDir(dirName, oldPathName, newPathName string) (*pb.MoveDirRes, error) {
	oldName := utils.Joins(oldPathName, dirName)
	newName := utils.Joins(newPathName, dirName)
	if utils.IsDirExist(oldName) {
		files, err := os.ReadDir(oldName)
		if err != nil {
			return res.CreateMoveDirRes(types.Fail, "Failed to read directory: "+oldName, err)
		}

		// Create dir
		resCreateDir, err := CreateDir(dirName, newPathName)
		if err != nil && resCreateDir.GetReturnCode() != types.ServerError {
			return res.CreateMoveDirRes(types.Fail, "Failed to create directory: "+oldName, err)
		}
		for _, file := range files {
			if file.IsDir() {
				// Move child dir
				MoveDir(file.Name(), oldName, newName)
			} else {
				// Move files
				_, err = MoveFile(file.Name(), oldName, newName)
				if err != nil {
					return res.CreateMoveDirRes(types.Fail, "Failed to move file: "+file.Name(), err)
				}
			}
		}
		return res.CreateMoveDirRes(types.Success, "Directory Moved: "+newName, nil)
	} else {
		return res.CreateMoveDirRes(types.Fail, "Directory Not Found: "+oldName,
			errors.New("Directory Not Found: "+oldName))
	}
}
