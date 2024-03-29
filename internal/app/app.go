package app

import (
	"context"
	"log"
	"net"

	"github.com/drizzleent/chat-server/internal/config"
	"github.com/drizzleent/chat-server/internal/interceptor"
	desc "github.com/drizzleent/chat-server/pkg/chat_v1"
	"github.com/drizzleent/chat-server/pkg/closer"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	err := a.runGRPCServer()
	if err != nil {
		log.Fatalf("Failed to run grpc server %s", err.Error())
	}

	return nil
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDebs(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) initDebs(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {

	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				interceptor.AuthInterceptor,
			),
		))

	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)

	if err != nil {
		return err
	}

	return nil
}
