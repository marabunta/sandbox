package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/marabunta/sandbox/mTLS/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {

	certificate, err := tls.LoadX509KeyPair(
		"./certs/server.crt",
		"./certs/server.key",
	)

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile("./certs/CA.pem")
	if err != nil {
		log.Fatalf("failed to read client ca cert: %s", err)
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		log.Fatal("failed to append client certs")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	}

	serverOption := grpc.Creds(credentials.NewTLS(tlsConfig))
	s := grpc.NewServer(serverOption)
	pb.RegisterGreeterServer(s, &server{})

	conn, err := net.Listen("tcp", ":1415")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
