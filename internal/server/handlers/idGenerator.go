package handlers

import pb "github.com/nifle3/goarenas-snowflakeClone/gen/api/proto"

type IdGenerator struct {
	pb.UnimplementedIdGeneratorServiceServer
}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{}
}
