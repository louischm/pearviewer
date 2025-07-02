package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/louischm/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"pearviewer/client/conf"
	pb "pearviewer/generated"
)

var log = logger.NewLog()

func getConn() *grpc.ClientConn {
	var confFile = conf.NewConf()

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("Cannot load TLS credentials: %v", err)
	}
	// Client startup
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
	conn, err := grpc.NewClient(confFile.GetServerAddress(), grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatal("Fatal error connecting to server: %s", err.Error())
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

func createUserClient() (*pb.UserServiceClient, *grpc.ClientConn) {
	conn := getConn()
	client := pb.NewUserServiceClient(conn)
	log.Info("Starting User client")
	return &client, conn
}

func closeClient(conn *grpc.ClientConn) {
	if errClose := conn.Close(); errClose != nil {
		log.Fatal("Failed to close connection: %s", errClose.Error())
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	var conf = conf.NewConf()
	pemServerCA, err := os.ReadFile(conf.CaCertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, errors.New("failed to add server CA's certificate")
	}

	clientCert, err := tls.LoadX509KeyPair(conf.ClientCert, conf.ClientKey)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		ServerName:   "localhost",
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientCert},
	}
	return credentials.NewTLS(config), nil
}
