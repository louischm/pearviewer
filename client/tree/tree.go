package tree

import (
	"errors"
	"github.com/louischm/pkg/logger"
	"github.com/louischm/pkg/utils"
	"os"
	"pearviewer/client/types"
)

var log = logger.NewLog()

func CreateTree(dirName, oldPathName, newPathName string) (*types.Dir, error) {

	oldName := utils.Joins(oldPathName, dirName)
	newName := utils.Joins(newPathName, dirName)

	if !utils.IsDirExist(oldName) {
		log.Debug("Dir %s does not exist", oldName)
		return nil, errors.New("Dir %s does not exist", oldName)
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
