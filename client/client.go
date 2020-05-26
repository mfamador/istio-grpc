package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

const (
	defaultName = "world"
	timeout     = 5 * time.Second
)

func main() {
	address := flag.String("address", "localhost:50051", "host:port of gRPC server")
	cert := flag.String("cert", "/data/cert.pem", "path to TLS certificate")
	repeat := flag.Int("repeat", 99, "number of unary gRPC requests to send")
	insecure := flag.Bool("insecure", false, "connect without TLS")
	flag.Parse()

	// Set up a connection to the server.
	var conn *grpc.ClientConn
	var err error
	if *insecure {
		log.Printf("parsed address %s", *address)
		conn, err = grpc.Dial(*address, grpc.WithInsecure())
	} else {
		tc, err := credentials.NewClientTLSFromFile(*cert, "")
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		conn, err = grpc.Dial(*address, grpc.WithTransportCredentials(tc))
	}
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Determine name to send to server.
	name := defaultName
	nonFlagArgs := make([]string, 0)
	for _, arg := range os.Args {
		if !strings.HasPrefix(arg, "--") {
			nonFlagArgs = append(nonFlagArgs, arg)
		}
	}
	if len(nonFlagArgs) > 1 {
		name = nonFlagArgs[1]
	}

	// Contact the server and print out its response multiple times.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for i := 0; i < *repeat; i++ {
		var header metadata.MD
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.Header(&header))
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		hostname := "unknown"
		// [START istio_sample_apps_grpc_greeter_go_client_hostname]
		if len(header["hostname"]) > 0 {
			hostname = header["hostname"][0]
		}
		log.Printf("%s from %s", r.Message, hostname)
		// [END istio_sample_apps_grpc_greeter_go_client_hostname]
	}
}
