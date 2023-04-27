package v1

import (
	"context"
	"fmt"
	"github.com/zfullio/price-placements-service/internal/domain/entity"
	"github.com/zfullio/price-placements-service/pb"
)

const msgErrMethod = "ошибка выполнения"
const msgMethodStarted = "Запущено"
const msgMethodFinished = "Завершено"

var placementFromPB = map[pb.Placement]entity.Placement{
	pb.Placement_PLACEMENT_UNKNOWN:       entity.UnknownPlacement,
	pb.Placement_PLACEMENT_CIAN:          entity.Cian,
	pb.Placement_PLACEMENT_YANDEX_REALTY: entity.Realty,
	pb.Placement_PLACEMENT_AVITO:         entity.Avito,
	pb.Placement_PLACEMENT_M2:            entity.M2,
	pb.Placement_PLACEMENT_DOMCLICK:      entity.DomClick,
	pb.Placement_PLACEMENT_JCAT:          entity.JCat,
	pb.Placement_PLACEMENT_CALUGA_HOUSE:  entity.CalugaHouse,
	pb.Placement_PLACEMENT_FLATOUTLET:    entity.FlatOutlet,
	pb.Placement_PLACEMENT_GDETOTDOM:     entity.GdeTotDom,
	pb.Placement_PLACEMENT_HYBRID:        entity.Hybrid,
	pb.Placement_PLACEMENT_AVAHO:         entity.Avaho,
	pb.Placement_PLACEMENT_NOVOSTROY_M:   entity.NovostroyM,
}

var PbFromPlacement = map[entity.Placement]pb.Placement{
	entity.UnknownPlacement: pb.Placement_PLACEMENT_UNKNOWN,
	entity.Cian:             pb.Placement_PLACEMENT_CIAN,
	entity.Realty:           pb.Placement_PLACEMENT_YANDEX_REALTY,
	entity.Avito:            pb.Placement_PLACEMENT_AVITO,
	entity.M2:               pb.Placement_PLACEMENT_M2,
	entity.DomClick:         pb.Placement_PLACEMENT_DOMCLICK,
	entity.JCat:             pb.Placement_PLACEMENT_JCAT,
	entity.CalugaHouse:      pb.Placement_PLACEMENT_CALUGA_HOUSE,
	entity.FlatOutlet:       pb.Placement_PLACEMENT_FLATOUTLET,
	entity.GdeTotDom:        pb.Placement_PLACEMENT_GDETOTDOM,
	entity.Hybrid:           pb.Placement_PLACEMENT_HYBRID,
	entity.Avaho:            pb.Placement_PLACEMENT_AVAHO,
	entity.NovostroyM:       pb.Placement_PLACEMENT_NOVOSTROY_M,
}

var PbFromStatusCheck = map[entity.StatusCheck]pb.Status{
	entity.Error:   pb.Status_ERROR,
	entity.Warning: pb.Status_WARNING,
	entity.OK:      pb.Status_OK,
}

func (s Server) CheckPhonesAll(ctx context.Context, req *pb.CheckPhonesAllRequest) (*pb.CheckPhonesAllResponse, error) {
	methodLogger := s.logger.With().Str("method", "CheckPhonesAll").Str("spreadsheet", req.SpreadsheetId).Logger()

	methodLogger.Info().Msg(msgMethodStarted)

	defer methodLogger.Info().Msg(msgMethodFinished)

	rawResult, err := s.policy.CheckPhonesAll(ctx, req.SpreadsheetId, req.Developer)

	if err != nil {
		methodLogger.Err(err).Msg(msgErrMethod)
		return nil, err
	}

	result := make([]*pb.CheckResult, 0, len(rawResult))
	for _, item := range rawResult {
		res, err := checkResultToPb(item)
		if err != nil {
			methodLogger.Err(err).Msg(msgErrMethod)
			return nil, err
		}
		result = append(result, res)
	}

	return &pb.CheckPhonesAllResponse{
		Result: result,
	}, nil
}
func (s Server) CheckPhonesRealty(ctx context.Context, req *pb.CheckPhonesRequest) (*pb.CheckPhonesResponse, error) {
	methodLogger := s.logger.With().Str("method", "CheckPhonesRealty").Str("spreadsheet", req.SpreadsheetId).Logger()

	methodLogger.Info().Msg(msgMethodStarted)

	defer methodLogger.Info().Msg(msgMethodFinished)

	rawResult, err := s.policy.CheckPhonesRealty(ctx, req.SpreadsheetId, req.Developer)
	if err != nil {
		methodLogger.Err(err).Msg(msgErrMethod)
		return nil, err
	}

	result := make([]*pb.CheckResult, 0, len(rawResult))
	for _, item := range rawResult {
		res, err := checkResultToPb(item)
		if err != nil {
			methodLogger.Err(err).Msg(msgErrMethod)
			return nil, err
		}
		result = append(result, res)
	}

	return &pb.CheckPhonesResponse{
		Result: result,
	}, nil
}
func (s Server) CheckPhonesCian(ctx context.Context, req *pb.CheckPhonesRequest) (*pb.CheckPhonesResponse, error) {
	methodLogger := s.logger.With().Str("method", "CheckPhonesCian").Str("spreadsheet", req.SpreadsheetId).Logger()

	methodLogger.Info().Msg(msgMethodStarted)

	defer methodLogger.Info().Msg(msgMethodFinished)

	rawResult, err := s.policy.CheckPhonesCian(ctx, req.SpreadsheetId, req.Developer)
	if err != nil {
		methodLogger.Err(err).Msg(msgErrMethod)
		return nil, err
	}

	result := make([]*pb.CheckResult, 0, len(rawResult))
	for _, item := range rawResult {
		res, err := checkResultToPb(item)
		if err != nil {
			methodLogger.Err(err).Msg(msgErrMethod)
			return nil, err
		}
		result = append(result, res)
	}

	return &pb.CheckPhonesResponse{
		Result: result,
	}, nil
}
func (s Server) CheckPhonesAvito(ctx context.Context, req *pb.CheckPhonesRequest) (*pb.CheckPhonesResponse, error) {
	methodLogger := s.logger.With().Str("method", "CheckPhonesAvito").Str("spreadsheet", req.SpreadsheetId).Logger()

	methodLogger.Info().Msg(msgMethodStarted)

	defer methodLogger.Info().Msg(msgMethodFinished)

	rawResult, err := s.policy.CheckPhonesAvito(ctx, req.SpreadsheetId, req.Developer)
	if err != nil {
		methodLogger.Err(err).Msg(msgErrMethod)
		return nil, err
	}

	result := make([]*pb.CheckResult, 0, len(rawResult))
	for _, item := range rawResult {
		res, err := checkResultToPb(item)
		if err != nil {
			methodLogger.Err(err).Msg(msgErrMethod)
			return nil, err
		}
		result = append(result, res)
	}

	return &pb.CheckPhonesResponse{
		Result: result,
	}, nil
}
func (s Server) CheckPhonesDomclick(ctx context.Context, req *pb.CheckPhonesRequest) (*pb.CheckPhonesResponse, error) {
	methodLogger := s.logger.With().Str("method", "CheckPhonesDomclick").Str("spreadsheet", req.SpreadsheetId).Logger()

	methodLogger.Info().Msg(msgMethodStarted)

	defer methodLogger.Info().Msg(msgMethodFinished)

	rawResult, err := s.policy.CheckPhonesDomclick(ctx, req.SpreadsheetId, req.Developer)
	if err != nil {
		methodLogger.Err(err).Msg(msgErrMethod)
		return nil, err
	}

	result := make([]*pb.CheckResult, 0, len(rawResult))
	for _, item := range rawResult {
		res, err := checkResultToPb(item)
		if err != nil {
			methodLogger.Err(err).Msg(msgErrMethod)
			return nil, err
		}
		result = append(result, res)
	}

	return &pb.CheckPhonesResponse{
		Result: result,
	}, nil
}

func (s Server) ValidateFeed(ctx context.Context, req *pb.ValidateFeedRequest) (*pb.ValidateFeedResponse, error) {
	methodLogger := s.logger.With().Str("method", "ValidateFeed").Str("feedUrl", req.FeedUrl).Logger()

	methodLogger.Info().Msg(msgMethodStarted)

	defer methodLogger.Info().Msg(msgMethodFinished)

	placement, ok := placementFromPB[req.Placement]
	if !ok {
		return nil, fmt.Errorf("not find placement: %s", placement)
	}

	rawResult, err := s.policy.ValidateFeed(ctx, req.FeedUrl, placement)
	if err != nil {
		return &pb.ValidateFeedResponse{
			Result: nil,
		}, err
	}

	result := make([]*pb.CheckResult, 0, len(rawResult))
	for _, item := range rawResult {
		res, err := checkResultToPb(item)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return &pb.ValidateFeedResponse{
		Result: result,
	}, nil
}
func (s Server) ValidateFeedAll(ctx context.Context, req *pb.ValidateFeedAllRequest) (*pb.ValidateFeedAllResponse, error) {
	methodLogger := s.logger.With().Str("method", "ValidateFeedAll").Str("spreadsheet", req.SpreadsheetId).Logger()

	methodLogger.Info().Msg(msgMethodStarted)

	defer methodLogger.Info().Msg(msgMethodFinished)

	rawResult, err := s.policy.ValidateFeedAll(ctx, req.SpreadsheetId, req.Developer)
	if err != nil {
		s.logger.Error().Err(err).Str("spreadsheet", req.SpreadsheetId).Msg("ValidateFeedAll")
		return nil, err
	}

	result := make([]*pb.CheckResult, 0, len(rawResult))
	for _, item := range rawResult {
		res, err := checkResultToPb(item)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}

	return &pb.ValidateFeedAllResponse{
		Result: result,
	}, err
}

func checkResultToPb(item entity.CheckResult) (*pb.CheckResult, error) {
	placement, ok := PbFromPlacement[item.Placement]
	if !ok {
		return nil, fmt.Errorf("invalid placement: %s", item.Placement)
	}

	base, ok := PbFromPlacement[item.Base]
	if !ok {
		return nil, fmt.Errorf("invalid base: %s", item.Base)
	}

	status, ok := PbFromStatusCheck[item.Status]
	if !ok {
		return nil, fmt.Errorf("invalid status")
	}

	return &pb.CheckResult{
		Developer: item.Developer,
		Placement: placement,
		Base:      base,
		Url:       item.Url,
		Message:   item.Message,
		Result:    status,
	}, nil

}
