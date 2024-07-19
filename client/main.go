package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "CApI/proto"
)

const (
	defaultQuery =  "hello!"
)

var (
	addr  = flag.String("addr", "localhost:50051", "address to connect to")
	query = flag.String("query", defaultQuery, "query to use")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewQuoteGrabberClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GrabQuote(ctx, &pb.QuoteRequest{ Temp: *query })

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("The quote is: %s", r.GetQuote())
}
