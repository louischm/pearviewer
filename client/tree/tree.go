package tree

import (
	"errors"
	"github.com/louischm/logger"
	"os"
	"pearviewer/client/types"
	"pearviewer/client/utils"
)

var log = logger.NewLog()

func CreateTree(dirName, oldPathName, newPathName string) (*types.Dir, error) {

	oldName := utils.Joins(oldPathName, dirName)
	newName := utils.Joins(newPathName, dirName)

	if !utils.IsDirExist(oldName) {
		log.Debug("Dir " + oldName + " does not exist")
		return nil, errors.New("Dir " + oldName + " does not exist")
	}

	files, err := os.ReadDir(oldName)
	if err != nil {
		return nil, err
	}

	dir := types.NewDir(dirName, oldPathName, newPathName)
	for _, f := range files {
		if f.IsDir() {
			child, errRec := CreateTree(f.Name(), oldName, newName)
			if errRec != nil {
				return nil, errRec
			}
			dir.SetChildren(append(dir.Children(), child))
		} else {
			AddFile(f.Name(), oldName, newName, dir)
		}
	}
	return dir, nil
}

func AddFile(name string, oldPathName, newPathName string, dir *types.Dir) {
	file := types.NewFile(name, oldPathName, newPathName)
	dir.SetFiles(append(dir.Files(), file))
}
