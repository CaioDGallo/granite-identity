package grpc

import (
	"fmt"
	"log/slog"
	"net"
	"sync"

	"github.com/CaioDGallo/granite-identity/internal/config"
	pb "github.com/CaioDGallo/granite-identity/internal/grpc"
	"github.com/CaioDGallo/granite-identity/internal/logger"
	"google.golang.org/grpc"
)

type GraniteGRPCServer struct {
	pb.UnimplementedAccountServiceServer
}

func (s *GraniteGRPCServer) StartListening(cfg *config.Config, wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
	if err != nil {
		logger.GetLogger().Error("gRPC failed to listen: ", slog.String("error", err.Error()))
	}
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &GraniteGRPCServer{})
	logger.GetLogger().Info("gRPC server listening at ", slog.String("port", lis.Addr().String()))
	if err := server.Serve(lis); err != nil {
		logger.GetLogger().Error("gRPC failed to serve: ", slog.String("error", err.Error()))
	}
}
