package auction

import (
	"log"
	"strconv"

	"golang.org/x/net/context"

	"sync/atomic"
)

type Server struct {
	UnimplementedChatServiceServer
}

var clients []ChatService_JoinServer = make([]ChatService_JoinServer, 0)


func Broadcast(ctx context.Context, message *Message) (*Empty, error) {

	for _, client := range clients {
		client.Send(message)
	}

	if message.User != "" {
		log.Printf("(%s, %s) >> %s"), message.User, message.Content)
	} else {
		log.Printf("(%s) >> %s", message.Content)
	}
	return &Empty{}, nil
}

func (s *Server) Join(message *JoinMessage, stream ChatService_JoinServer) error {
	clients = append(clients, stream)

	msg := Message{
		User:    "",
		Content: "Participant " + message.User + " joined AuctionHouse ",
	}

	s.Broadcast(context.TODO(), &msg)

	for {
		select {
		case <-stream.Context().Done():
			msg := Message{
				User:    "",
				Content: "Participant " + message.User + " left the AuctionHouse",
			}
			for i, element := range clients {
				if element == stream {
					clients = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			s.Broadcast(context.TODO(), &msg)
			return nil
		}
	}

}
