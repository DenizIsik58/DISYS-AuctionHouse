package main

import (
	"ChatService/auction"
	"bufio"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//	response, _ := client.SendMessage(context.Background(), &auction.Message{Content: "Hello from the client!"})
//	log.Printf("Response from server: %s", response.Content)

var name string
var client auction.AuctionHouseClient
var ctx context.Context
var connections = []string{"8000", "7000", "9000"}
var openConnections = make([]string, 0)

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

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomName(length int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


func main() {
	// Handle flags

	name = randomName(15)

	// Handle connection
	conn, err := grpc.Dial(":7000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect! %s", err)
		return
	}

	defer conn.Close()

	client = auction.NewAuctionHouseClient(conn)
	ctx = context.Background()

	go Join()


	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go Handle(strings.ToLower(scanner.Text()))
	}

}


func Handle(message string) {


	IsAlive("localhost", connections)


	if strings.Contains(message, "bid ") {
		bid := strings.Replace(message, "bid ", "", -1)
		intVar, _ := strconv.ParseInt(bid, 0, 64)

		for i, _ := range openConnections {

			conn, err := grpc.Dial(":" + openConnections[i], grpc.WithInsecure())

			if err != nil {
				log.Fatalf("Could not connect! %s", err)
				return
			}

			defer conn.Close()

			client = auction.NewAuctionHouseClient(conn)
			ctx = context.Background()

			Bid(intVar)
		}

	}

	if message == "result" {
		conn, err := grpc.Dial(":" + openConnections[0], grpc.WithInsecure())

		if err != nil {
			log.Fatalf("Could not connect! %s", err)
			return
		}

		defer conn.Close()

		client = auction.NewAuctionHouseClient(conn)
		ctx = context.Background()
			Result()
	}
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


func IsAlive(host string, ports []string) {
	openConnections = []string {}
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Println("Connecting error:", err)
		}
		if conn != nil {
			defer conn.Close()

			openConnections = append(openConnections, port)

		}
	}
}


