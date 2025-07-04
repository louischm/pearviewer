package main

import (
	"github.com/louischm/pkg/logger"
	"pearviewer/client/ui"
)

var log = logger.NewLog()

func main() {
	// Log setup
	log.SetFileOutName("./log/main.log")
	log.SetFileErrName("./log/main.err")
	log.SetMaxSize(1e+9)

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

	// Run main window
	ui.CreateBody()
}
