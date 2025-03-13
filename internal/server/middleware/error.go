package middleware

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errMissingPeer     = status.Errorf(codes.InvalidArgument, "missing peer")
)
