package v1

import (
	"price-placements-service/internal/domain/policy"
	"price-placements-service/pb"
)

type Server struct {
	policy policy.FeedPolicy
	pb.UnimplementedFeedServiceServer
}

func NewServer(feedPolicy policy.FeedPolicy, srv pb.UnimplementedFeedServiceServer) *Server {
	return &Server{
		policy:                         feedPolicy,
		UnimplementedFeedServiceServer: srv,
	}
}
