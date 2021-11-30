package auction

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"strconv"
	"sync"
)

type Server struct {
	UnimplementedAuctionHouseServer
}

var clients []AuctionHouse_JoinServer = make([]AuctionHouse_JoinServer, 0)
var highestBid int64
var NameOfHighestBidder string
var isOver bool
var lock sync.Mutex

func Broadcast(ctx context.Context, message *Message) (*Empty, error) {

	for _, client := range clients {
		client.Send(message)
	}

	if message.User != "" {
		log.Printf("(%s) >> %s", message.User, message.Content)
	} else {
		log.Printf("%s", message.Content)
	}
	return &Empty{}, nil
}

func (s *Server) Join(message *JoinMessage, stream AuctionHouse_JoinServer) error {
	clients = append(clients, stream)

	msg := Message{
		User:    "",
		Content: "Participant " + message.User + ", has entered the AuctionHouse ",
	}

	Broadcast(context.TODO(), &msg)

	for {
		select {
		case <-stream.Context().Done():
			msg := Message{
				User:    "",
				Content: "Participant " + message.User + ", has left the AuctionHouse",
			}
			for i, element := range clients {
				if element == stream {
					clients = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			Broadcast(context.TODO(), &msg)
			return nil
		}
	}
}

func (s *Server) Bid(ctx context.Context, bidmsg *BidMessage) (*BidResponse, error){
lock.Lock()

	if bidmsg.Bid <= highestBid{
		lock.Unlock()
		return &BidResponse{Valid: false, HighestBid: highestBid}, nil;
	}

	highestBid = bidmsg.Bid
	NameOfHighestBidder = bidmsg.User
lock.Unlock()
	Broadcast(ctx, &Message{User: bidmsg.User, Content: fmt.Sprintf("%s, has the highest bid of %s)", bidmsg.User, strconv.Itoa(int(bidmsg.Bid)))})
	return &BidResponse{Valid: true}, nil

}
func (s *Server) Result(ctx context.Context, emp *Empty) (*BidMessage, error){
	return &BidMessage{Bid: highestBid, User: NameOfHighestBidder}, nil
}
