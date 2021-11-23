package main

import (
	"ChatService/auction"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Print("Loading Auction...")

	listener, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		log.Fatalf("TCP failed to listen... %s", err)
		return
	}

	log.Print("Listener registered - setting up server now...")

	s := auction.Server{}

	grpcServer := grpc.NewServer()

	auction.RegisterChatServiceServer(grpcServer, &s)

	log.Print("===============================================================================")
	log.Print("                            Welcome to AuctionHouse!                            ")
	log.Print("      Users can connect at any time and make bids for the current auction!       ")
	log.Print("===============================================================================")

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Failed to server gRPC serevr over port 9000")
	}
}
