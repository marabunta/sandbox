package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/marabunta/sandbox/mTLS/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certificate, err := tls.LoadX509KeyPair(
		"./certs/client.crt",
		"./certs/client.key",
	)

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile("./certs/CA.pem")
	if err != nil {
		log.Fatalf("failed to read ca cert: %s", err)
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		log.Fatal("failed to append certs")
	}

	transportCreds := credentials.NewTLS(&tls.Config{
		ServerName:   "server.example.com",
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	dialOption := grpc.WithTransportCredentials(transportCreds)
	conn, err := grpc.Dial("localhost:1415", dialOption)
	if err != nil {
		log.Fatalf("failed to dial server: %s", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Greting: %s", r.Message)
}
