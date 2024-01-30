package handler

import (
	"context"
	crypto "crypto/rand"
	"log"
	"math/big"

	"github.com/drizzleent/chat-server/internal/lib/repository"
	"github.com/drizzleent/chat-server/internal/model"
	desc "github.com/drizzleent/chat-server/pkg/chat_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatGrpcServer struct {
	desc.UnimplementedChatV1Server
	log *log.Logger
	db  repository.Chatting
}

func NewChatGrpcServer(log *log.Logger, db repository.Chatting) *ChatGrpcServer {
	return &ChatGrpcServer{
		log: log,
		db:  db,
	}
}

func (s *ChatGrpcServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	s.log.Println("Recive create message")

	for i, user := range req.Usernames {
		s.log.Printf("\t#%d Username: %s\n", i, user)
	}

	if dline, ok := ctx.Deadline(); ok {
		s.log.Printf("Deadline: %v\n", dline)
	}

	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(100123))

	if err != nil {
		s.log.Printf("cant generate rand in %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	id := safeNum.Int64()

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
func (s *ChatGrpcServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	s.log.Println("Recive Delete message")

	if dline, ok := ctx.Deadline(); ok {
		s.log.Printf("Deadline: %v\n", dline)
	}

	return &emptypb.Empty{}, nil
}
func (s *ChatGrpcServer) SendMessage(ctx context.Context, req *desc.SendRequest) (*empty.Empty, error) {
	s.log.Println("Received SendMessage")

	s.log.Printf("Message:\n From: %v\nText: %v\nTime: %v", req.GetFrom(), req.GetText(), req.GetTimestamp())

	if dline, ok := ctx.Deadline(); ok {
		s.log.Printf("Deadline: %v\n", dline)
	}

	user := model.User{
		Username: req.From,
		Message: model.Message{
			Text: req.Text,
		},
	}

	err := s.db.SendMessage(ctx, user)

	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
