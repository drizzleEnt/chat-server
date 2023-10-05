package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/drizzleent/chat-server/internal/config"
	"github.com/drizzleent/chat-server/internal/lib/handler"
	"github.com/drizzleent/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.MustConfig()

	aLog := log.New(os.Stdout, "[INFO]:", log.Flags())

	url := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	lis, err := net.Listen("tcp", url)

	if err != nil {
		aLog.Fatalf("Failed to listen server %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	rpcSrv := handler.NewChatGrpcServer(aLog)
	chat_v1.RegisterChatV1Server(s, rpcSrv)

	done := make(chan os.Signal, 1)

	go func() {
		if err := s.Serve(lis); err != nil {
			aLog.Fatalf("Failed to serve %v", err)
		}

	}()

	aLog.Println("Server Started")
	<-done
	s.GracefulStop()
	aLog.Println("Server stopped")
}
