package main

import (
	"errors"
	pb "grpc-demo/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return stream.SendAndClose(&pb.MessagesList{})
		}
		if err != nil {
			return err
		}

		log.Printf("Got request with name: %s", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
