package auction

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedAuctionHouseServer
}

var clients []AuctionHouse_JoinServer = make([]AuctionHouse_JoinServer, 0)
var highestBid int64
var NameOfHighestBidder string

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
