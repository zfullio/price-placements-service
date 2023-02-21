package v1

import (
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/domain/policy"
	"github.com/zfullio/price-placements-service/pb"
)

type Server struct {
	policy policy.FeedPolicy
	logger *zerolog.Logger
	pb.UnimplementedFeedServiceServer
}

func NewServer(feedPolicy policy.FeedPolicy, logger *zerolog.Logger, srv pb.UnimplementedFeedServiceServer) *Server {
	apiLogger := logger.With().Str("api", "grpc").Logger()
	return &Server{
		policy:                         feedPolicy,
		logger:                         &apiLogger,
		UnimplementedFeedServiceServer: srv,
	}
}
