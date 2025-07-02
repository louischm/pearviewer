package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/louischm/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	pb "pearviewer/generated"
	"pearviewer/server/conf"
)

var log = logger.NewLog()

type fileServer struct {
	pb.FileServiceServer
}

type dirServer struct {
	pb.DirServiceServer
}

type userServer struct {
	pb.UserServiceServer
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	var conf = conf.NewConf()
	pemClientCA, err := os.ReadFile(conf.CaCertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, errors.New("failed to append client certs")
	}

	serverCert, err := tls.LoadX509KeyPair(conf.ServerCert, conf.ServerKey)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func StartServer(confData *conf.Conf) {
	lis, err := net.Listen("tcp", confData.GetServerAddress())
	if err != nil {
		log.Fatal("Can't start server: %s", err.Error())
	}

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("Failed to load TLS credentials: %v", err)
	}

	var opts = []grpc.ServerOption{
		grpc.Creds(tlsCredentials),
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFileServiceServer(grpcServer, &fileServer{})
	pb.RegisterDirServiceServer(grpcServer, &dirServer{})
	pb.RegisterUserServiceServer(grpcServer, &userServer{})
	errServer := grpcServer.Serve(lis)
	if errServer != nil {
		log.Fatal("Can't start server: %s", errServer.Error())
	}
}
