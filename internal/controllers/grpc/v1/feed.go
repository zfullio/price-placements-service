package v1

import (
	"context"
	"fmt"
	"price-placements-service/internal/domain/entity"
	"price-placements-service/pb"
)

func (s Server) CheckPhonesAll(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	res, err := s.policy.CheckPhonesAll(req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, nil
}
func (s Server) CheckPhonesRealty(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	res, err := s.policy.CheckPhonesRealty(req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, nil
}
func (s Server) CheckPhonesCian(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	res, err := s.policy.CheckPhonesCian(req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, nil
}
func (s Server) CheckPhonesAvito(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	res, err := s.policy.CheckPhonesAvito(req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, err
}
func (s Server) CheckPhonesDomclick(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	res, err := s.policy.CheckPhonesDomclick(req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, err
}

func (s Server) ValidateFeed(ctx context.Context, req *pb.ValidateFeedRequest) (*pb.ValidateFeedResponse, error) {
	var placement entity.Placement
	switch req.Placement {
	case pb.Placement_PLACEMENT_YANDEX_REALTY:
		placement = entity.Realty
	case pb.Placement_PLACEMENT_CIAN:
		placement = entity.Cian
	case pb.Placement_PLACEMENT_AVITO:
		placement = entity.Avito
	case pb.Placement_PLACEMENT_DOMCLICK:
		placement = entity.Domclick

	}
	results, err := s.policy.ValidateFeed(req.FeedUrl, placement)
	if err != nil {
		return &pb.ValidateFeedResponse{
			Result: results,
		}, err
	}

	if len(results) == 0 {
		results = append(results, fmt.Sprintf("feed: %s. placement: %s. Все в порядке", req.FeedUrl, placement))
	}
	return &pb.ValidateFeedResponse{
		Result: results,
	}, nil
}
func (s Server) ValidateFeedAll(ctx context.Context, req *pb.ValidateFeedAllRequest) (*pb.ValidateFeedAllResponse, error) {
	res, err := s.policy.ValidateFeedAll(req.SpreadsheetId)
	return &pb.ValidateFeedAllResponse{
		Result: res,
	}, err
}
