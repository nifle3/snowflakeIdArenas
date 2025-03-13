package converters

import (
	"log/slog"

	pb "github.com/nifle3/goarenas-snowflakeClone/gen/api/proto"
	"github.com/nifle3/goarenas-snowflakeClone/internal/domain/snowflakeid"
)

func FromGrpcTypeToDomain(idType pb.IdType) snowflakeid.TypeToGenerate {
	var result snowflakeid.TypeToGenerate

	switch idType {
	case pb.IdType_base64:
		result = snowflakeid.Base64
	case pb.IdType_binary:
		result = snowflakeid.Binary
	case pb.IdType_text:
		result = snowflakeid.Text
	default:
		slog.Warn("Default idtype to typeToGenerate", slog.Any("GRPC type", idType))
		result = snowflakeid.Binary
	}

	return result
}
