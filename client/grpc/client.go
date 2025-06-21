package grpc

import (
	"github.com/louischm/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	. "pearviewer/client/conf"
	pdir "pearviewer/generated/dir"
	pfile "pearviewer/generated/file"
)

var log = logger.NewLog()

func getConn() *grpc.ClientConn {
	// Conf setup
	confData := GetConf()
	log.Info("Conf loaded")
	// Client startup
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(confData.GetServerAddress(), opts...)
	if err != nil {
		log.Fatal("Fatal error connecting to server" + err.Error())
	}
	return conn
}

func createFileClient() (*pfile.FileServiceClient, *grpc.ClientConn) {
	conn := getConn()
	client := pfile.NewFileServiceClient(conn)
	log.Info("Starting File client")
	return &client, conn
}

func createDirClient() (*pdir.DirServiceClient, *grpc.ClientConn) {
	conn := getConn()
	client := pdir.NewDirServiceClient(conn)
	log.Info("Starting Dir client")
	return &client, conn
}

func closeClient(conn *grpc.ClientConn) {
	if errClose := conn.Close(); errClose != nil {
		log.Fatal("Failed to close connection" + errClose.Error())
	}
}
