package grpc

import (
	"github.com/louischm/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pearviewer/client/conf"
	pb "pearviewer/generated"
)

var log = logger.NewLog()

func getConn() *grpc.ClientConn {
	var confFile = conf.NewConf()
	// Client startup
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(confFile.GetServerAddress(), opts...)
	if err != nil {
		log.Fatal("Fatal error connecting to server" + err.Error())
	}
	return conn
}

func createFileClient() (*pb.FileServiceClient, *grpc.ClientConn) {
	conn := getConn()
	client := pb.NewFileServiceClient(conn)
	log.Info("Starting File client")
	return &client, conn
}

func createDirClient() (*pb.DirServiceClient, *grpc.ClientConn) {
	conn := getConn()
	client := pb.NewDirServiceClient(conn)
	log.Info("Starting Dir client")
	return &client, conn
}

func closeClient(conn *grpc.ClientConn) {
	if errClose := conn.Close(); errClose != nil {
		log.Fatal("Failed to close connection" + errClose.Error())
	}
}
