package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/grpc"
	pb "CApI/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedQuoteGrabberServer
}

func (s *server) GrabQuote(ctx context.Context, in *pb.QuoteRequest) (*pb.QuoteReply, error) {
	log.Printf("Received: %v", in.GetTemp())
	return &pb.QuoteReply{
		Quote: "Boosh is the time of the day",
		Time:  timestamppb.Now(),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil { log.Fatalf("failed to listen: %v", err) }
	s := grpc.NewServer()

	pb.RegisterQuoteGrabberServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
