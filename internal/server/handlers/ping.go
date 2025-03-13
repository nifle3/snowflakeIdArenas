package handlers

import pb "github.com/nifle3/goarenas-snowflakeClone/gen/api/proto"

type Ping struct {
	pb.UnimplementedPingServiceServer
}

func NewPing() *Ping {
	return &Ping{}
}
