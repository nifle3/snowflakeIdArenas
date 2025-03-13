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
	model, err := ig.service.Generate(ctx, typeToGenerate)
	if err != nil {
		return nil, err
	}
	req := &pb.GenerateRequest{
		Format: res.Format,
		Value: &pb.GenerateRequest_IdInt64{
			IdInt64: model.Base,
		},
	}

	return req, nil
}

func (ig IdGenerator) GenerateBatch(ctx context.Context, res *pb.GenerateResponseBatch) (*pb.GenerateRequestBatch, error) {
	typeToGenerate := converters.FromGrpcTypeToDomain(res.Format)

	ig.service.GenerateBatch(ctx, int(res.Count), typeToGenerate)

	return &pb.GenerateRequestBatch{}, nil
}
