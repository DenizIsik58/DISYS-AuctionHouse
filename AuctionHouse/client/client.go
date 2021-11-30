package main

import (
	"AuctionHouse/auction"
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var name string

var client auction.AuctionHouseClient

var ctx context.Context = context.Background()
var ports = []string{"8000", "7000", "9000"}
var openConnections = make([]string, 0)

func Join(port string) {
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect! %s", err)
		return
	}

	defer conn.Close()

	client := auction.NewAuctionHouseClient(conn)
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

func RandomName(length int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	// Handle flags
	name = RandomName(15)

	// Handle connection
	IsAlive("localhost", ports)

	for _, connection := range openConnections {
		go Join(connection)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go Handle(strings.ToLower(scanner.Text()))
	}
}

func Handle(message string) {
	if !strings.HasPrefix(message, "bid ") && message != "result" {
		log.Printf(">> Unknown command, please use bid or result")
		return
	}

	IsAlive("localhost", ports)

	for _, connection := range openConnections {
		conn, err := grpc.Dial(":"+connection, grpc.WithInsecure())

		if err != nil {
			log.Fatalf("Could not connect! %s", err)
		}

		defer conn.Close()

		client = auction.NewAuctionHouseClient(conn)
		ctx = context.Background()

		if message == "result" {
			Result()
		} else {
			bid := strings.Replace(message, "bid ", "", -1)
			intVar, _ := strconv.ParseInt(bid, 0, 64)
			Bid(intVar)
		}
	}
}

func Bid(amount int64) {
	res, _ := client.Bid(ctx, &auction.BidMessage{Bid: amount, User: name})

	if !res.Valid {
		log.Printf("Your bid wasn't high enough, try again..")
		log.Printf("The highest bid is currently: %s", strconv.Itoa(int(res.HighestBid)))
	}
}

func Result() {
	res, _ := client.Result(ctx, &auction.Empty{})

	if res.User == "" {
		log.Println("There is no active bid right now - be the first by typing bid <amount>")
		return
	}

	log.Printf("%s has the highest bid with: %s", res.User, strconv.Itoa(int(res.Bid)))
}

func IsAlive(host string, ports []string) {
	openConnections = []string{}

	for _, port := range ports {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second)

		if err != nil {
			fmt.Println("Connecting error:", err)
		}

		if conn != nil {
			openConnections = append(openConnections, port)
		}
	}
}
