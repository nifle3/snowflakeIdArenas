package handlers

import (
	"context"

	pb "github.com/nifle3/goarenas-snowflakeClone/gen/api/proto"
)

type Ping struct {
	pb.UnimplementedPingServiceServer
}

func NewPing() *Ping {
	return &Ping{}
}

func (Ping) Ping(context.Context, *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
