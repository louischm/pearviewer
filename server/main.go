package main

import (
	"github.com/louischm/pkg/logger"
	"pearviewer/server/conf"
	"pearviewer/server/grpc"
)

var log = logger.NewLog()

func main() {
	// Log setup
	log.SetFileOutName("./log/main.log")
	log.SetFileErrName("./log/main.err")
	log.SetMaxSize(1e+9)

	// Conf setup
	confData := conf.GetConf()
	log.Info("Conf loaded")
	grpc.StartServer(confData)
}
