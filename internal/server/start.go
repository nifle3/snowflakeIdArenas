package server

import (
	"fmt"
	"log/slog"
	"net"

	pb "github.com/nifle3/goarenas-snowflakeClone/gen/api/proto"
	"github.com/nifle3/goarenas-snowflakeClone/internal/domain/snowflakeid"
	"github.com/nifle3/goarenas-snowflakeClone/internal/server/handlers"
	"github.com/nifle3/goarenas-snowflakeClone/internal/server/middleware"
	"google.golang.org/grpc"
)

func MustStart(service snowflakeid.Service) {
	cfg := mustNewConfig()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		slog.Error("Open port failed", slog.String("err", err.Error()))
		panic(fmt.Sprintf("Open port failed with %s\n", err.Error()))
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middleware.Recovery, middleware.Logging),
	)

	pingHandler := handlers.NewPing()
	idGeneratorHandler := handlers.NewIdGenerator(service)

	pb.RegisterIdGeneratorServiceServer(server, idGeneratorHandler)
	pb.RegisterPingServiceServer(server, pingHandler)

	slog.Info("Start server on ", slog.String("port", cfg.Port), slog.String("Host", cfg.Host))
	err = server.Serve(lis)
	if err != nil {
		slog.Error("Serve server err", slog.String("error", err.Error()))
		panic("Serve server err")
	}
}
