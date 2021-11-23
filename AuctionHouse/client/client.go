package main

import (
	"ChatService/auction"
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

//	response, _ := client.SendMessage(context.Background(), &auction.Message{Content: "Hello from the client!"})
//	log.Printf("Response from server: %s", response.Content)

var name string
var client auction.AuctionHouseClient
var ctx context.Context

func Join() {
	stream, _ := client.Join(context.Background(), &auction.JoinMessage{User: name})

	for {
		response, err := stream.Recv()

		if err != nil {
			break
		}


		if response.User == "" {
			log.Default().Printf("%s", response.Content)
			continue
		}

		log.Default().Printf("(%s) >> %s", response.User, response.Content)
	}
}

func main() {
	// Handle flags
	nameFlag := flag.String("name", "", "")

	flag.Parse()
	name = *nameFlag

	// Handle connection
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect! %s", err)
		return
	}

	defer conn.Close()

	client = auction.NewAuctionHouseClient(conn)
	ctx = context.Background()

	go Join()

}

func Bid(amount int64){
	res,_ := client.Bid(ctx, &auction.BidMessage{Bid: amount, User: name})

	if !res.Valid{
		fmt.Printf("Your bid wasn't high enough, try again..")
		fmt.Printf("The highest bid is currently: %s", strconv.Itoa(int (res.HighestBid)))
	}

}

func Result(){
	res,_ :=client.Result(ctx, &auction.Empty{})

	fmt.Printf("%s has the highest bid with: %s", res.User, strconv.Itoa(int (res.Bid)))
}
