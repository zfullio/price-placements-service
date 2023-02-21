package v1

import (
	"context"
	"fmt"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
	"github.com/zfullio/price-placements-service/pb"
)

var placementFromPB = map[pb.Placement]entity.Placement{
	pb.Placement_PLACEMENT_YANDEX_REALTY: entity.Realty,
	pb.Placement_PLACEMENT_CIAN:          entity.Cian,
	pb.Placement_PLACEMENT_AVITO:         entity.Avito,
	pb.Placement_PLACEMENT_DOMCLICK:      entity.Domclick,
}

func (s Server) CheckPhonesAll(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	s.logger.Trace().Msg("CheckPhonesAll")

	res, err := s.policy.CheckPhonesAll(ctx, req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, nil
}
func (s Server) CheckPhonesRealty(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	s.logger.Trace().Msg("CheckPhonesRealty")

	res, err := s.policy.CheckPhonesRealty(ctx, req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, nil
}
func (s Server) CheckPhonesCian(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	s.logger.Trace().Msg("CheckPhonesCian")

	res, err := s.policy.CheckPhonesCian(ctx, req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, nil
}
func (s Server) CheckPhonesAvito(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	s.logger.Trace().Msg("CheckPhonesAvito")

	res, err := s.policy.CheckPhonesAvito(ctx, req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, err
}
func (s Server) CheckPhonesDomclick(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	s.logger.Trace().Msg("CheckPhonesDomclick")

	res, err := s.policy.CheckPhonesDomclick(ctx, req.SpreadsheetId)
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{
		Result: res,
	}, err
}

func (s Server) ValidateFeed(ctx context.Context, req *pb.ValidateFeedRequest) (*pb.ValidateFeedResponse, error) {
	s.logger.Trace().Msg("ValidateFeed")

	placement, ok := placementFromPB[req.Placement]
	if !ok {
		return nil, fmt.Errorf("not find placement: %s", placement)
	}

	results, err := s.policy.ValidateFeed(ctx, req.FeedUrl, placement)
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
	s.logger.Trace().Msg("ValidateFeedAll")

	res, err := s.policy.ValidateFeedAll(ctx, req.SpreadsheetId)

	return &pb.ValidateFeedAllResponse{
		Result: res,
	}, err
}
