package main

import (
	"github.com/louischm/pkg/logger"
	"github.com/louischm/pkg/utils"
	"os"
	"pearviewer/server/conf"
	"pearviewer/server/grpc"
)

var log = logger.NewLog()

func postConstruct(conf *conf.Conf) {
	if !utils.IsDirExist(conf.LogPath) {
		err := os.Mkdir(conf.LogPath, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating log directory: %v", err)
		}
	}

	if !utils.IsDirExist(conf.DataPath) {
		err := os.Mkdir(conf.DataPath, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating data directory: %v", err)
		}
	}
}

func main() {
	// Conf setup
	confData := conf.NewConf()
	postConstruct(confData)

	// Log setup
	log.SetFileOutName("./log/main.log")
	log.SetFileErrName("./log/main.err")
	log.SetMaxSize(1e+9)
	log.Info("Conf and log loaded")

	grpc.StartServer(confData)
}
