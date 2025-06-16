package main

import (
	"github.com/louischm/logger"
	"pearviewer/client/grpc"
)

var log = logger.NewLog()

func main() {
	// Log setup
	log.SetFileOutName("./log/main.log")
	log.SetFileErrName("./log/main.err")
	log.SetMaxSize(1e+9)

	grpc.RenameFile("test.txt", "test1.csv", "./filetest/")
}
