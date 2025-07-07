package main

import (
	"github.com/louischm/pkg/logger"
	"github.com/louischm/pkg/utils"
	"os"
	"pearviewer/client/conf"
	"pearviewer/client/ui"
)

var log = logger.NewLog()

func postConstruct(conf *conf.Conf) {
	if !utils.IsDirExist(conf.LogPath) {
		err := os.Mkdir(conf.LogPath, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating log directory: %v", err)
		}
	}
}

func main() {

	/*	grpc.UploadFile("./filetest/test.txt", "./filetest/")*/
	//	grpc.RenameFile("test.txt", "test.csv", "./filetest/")
	// grpc.DeleteFile("test.csv", "./filetest/")
	//grpc.MoveFile("test.csv", "./filetest/", "./newFiletest/")
	//grpc.CreateDir("newDir", "./filetest/")
	//grpc.UploadDir("test", "./filetest/", "./newFileTest/")
	//grpc.MoveDir("test", "./newFileTest/", "./filetest/")
	//grpc.ListDir("test", "./filetest/")
	//grpc.DownloadFile("test.txt", "./filetest/", "./filetest/")
	//grpc.DownloadDir("test", "./filetest/", "./filetest/")

	// Log setup
	confData := conf.NewConf()
	postConstruct(confData)
	log.SetFileOutName("./log/main.log")
	log.SetFileErrName("./log/main.err")
	log.SetMaxSize(1e+9)

	// Run main window
	ui.Run()
}
