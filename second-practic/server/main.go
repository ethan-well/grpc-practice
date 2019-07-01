package main

import (
	pb "grpc-learn/second-practic/chat"
	"io"

	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

var answer1 = pb.Response{Answer: "answer1"}
var answer2 = pb.Response{Answer: "answer2"}
var answer3 = pb.Response{Answer: "answer3"}
var answers = [...]*pb.Response{&answer1, &answer2, &answer3}

func (s *server) QA(stream pb.Chat_QAServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		for _, w := range request.Question {
			if err := stream.Send(&pb.Response{Answer: string(w)}); err != nil {
				return err
			}
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:50001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
