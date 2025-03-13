package middleware

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errMissingPeer     = status.Errorf(codes.InvalidArgument, "missing peer")
)

func Logging(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errMissingPeer
	}

	slog.Info("gRPC received",
		slog.String("method", info.FullMethod),
		slog.String("client", p.Addr.String()),
		slog.Any("metadata", md),
	)

	resp, err := handler(ctx, req)

	code := status.Code(err)
	slog.Info("gRPC complete",
		slog.String("method", info.FullMethod),
		slog.String("client", p.Addr.String()),
		slog.String("status", code.String()),
		slog.Duration("duration", time.Since(start)),
		slog.Any("error", err),
	)

	return resp, err
}
