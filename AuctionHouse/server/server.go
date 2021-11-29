package main

import (
	"ChatService/auction"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Print("Loading Auction...")

	//auctionItem := flag.String("item", "", "")
	//flag.Parse()
	portFlag := flag.String("port", "", "")

	flag.Parse()
	port := *portFlag

	listener, err := net.Listen("tcp", "localhost:" + port)

	if err != nil {
		log.Fatalf("TCP failed to listen... %s", err)
		return
	}

	log.Print("Listener registered - setting up server now...")

	s := auction.Server{}

	grpcServer := grpc.NewServer()

	auction.RegisterAuctionHouseServer(grpcServer, &s)

	log.Print("===============================================================================")
	log.Print("                            Welcome to AuctionHouse!                            ")
	log.Print("      Users can connect at any time and make bids for the current auction!       ")
	log.Print("===============================================================================")

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Failed to server gRPC serevr over port 9000")
	}
}
