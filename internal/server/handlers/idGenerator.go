package handlers

import (
	"context"

	pb "github.com/nifle3/goarenas-snowflakeClone/gen/api/proto"
	"github.com/nifle3/goarenas-snowflakeClone/internal/domain/snowflakeid"
	"github.com/nifle3/goarenas-snowflakeClone/internal/server/converters"
)

type IdGenerator struct {
	pb.UnimplementedIdGeneratorServiceServer

	service snowflakeid.Service
}

func NewIdGenerator(service snowflakeid.Service) *IdGenerator {
	return &IdGenerator{
		service: service,
	}
}

func (ig IdGenerator) Generate(ctx context.Context, res *pb.GenerateResponse) (*pb.GenerateRequest, error) {
	typeToGenerate := converters.FromGrpcTypeToDomain(res.Format)
	ig.service.Generate(ctx, typeToGenerate)
	return nil, nil
}

func (ig IdGenerator) GenerateBatch(ctx context.Context, res *pb.GenerateResponseBatch) (*pb.GenerateRequest, error) {
	typeToGenerate := converters.FromGrpcTypeToDomain(res.Format)

	ig.service.GenerateBatch(ctx, int(res.Count), typeToGenerate)

	return &pb.GenerateRequest{}, nil
}
