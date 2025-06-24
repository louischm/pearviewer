package grpc

import (
	"github.com/louischm/logger"
	"google.golang.org/grpc"
	"net"
	pdir "pearviewer/generated"
	pfile "pearviewer/generated"
	"pearviewer/server/conf"
)

var log = logger.NewLog()

type fileServer struct {
	pfile.FileServiceServer
}

type dirServer struct {
	pdir.DirServiceServer
}

func StartServer(confData *conf.Conf) {
	lis, err := net.Listen("tcp", confData.GetServerAddress())
	if err != nil {
		log.Fatal("Can't start server" + err.Error())
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pfile.RegisterFileServiceServer(grpcServer, &fileServer{})
	pdir.RegisterDirServiceServer(grpcServer, &dirServer{})
	errServer := grpcServer.Serve(lis)
	if errServer != nil {
		log.Fatal("Can't start server" + errServer.Error())
	}
}
