package app

import (
	"context"
	"fmt"
	"github.com/nikoksr/notify"
	"github.com/rs/zerolog"
	"github.com/zfullio/price-placements-service/internal/adapters/gs"
	"github.com/zfullio/price-placements-service/internal/config"
	priceplacement "github.com/zfullio/price-placements-service/internal/controllers/grpc/v1"
	"github.com/zfullio/price-placements-service/internal/domain/policy"
	"github.com/zfullio/price-placements-service/internal/domain/service"
	"github.com/zfullio/price-placements-service/pb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	cfg                  config.Config
	grpcServer           *grpc.Server
	productServiceServer pb.FeedServiceServer
	Logger               *zerolog.Logger
	Notify               notify.Notifier
}

func NewApp(ctx context.Context, logger *zerolog.Logger, cfg config.Config, notify notify.Notifier) App {
	ghSrv, err := sheets.NewService(ctx, option.WithCredentialsFile(cfg.GS.ServiceKeyPath))
	if err != nil {
		logger.Fatal().Err(err).Msg("can't init app")
	}

	ghPhoneRepo := gs.NewPhoneRepository(*ghSrv, logger)
	phSrv := service.NewPhoneService(ghPhoneRepo, logger)

	ghFeedRepo := gs.NewFeedRepository(*ghSrv, logger)
	feedSrv := service.NewFeedService(ghFeedRepo, logger)

	feedPolicy := policy.NewFeedPolicy(*feedSrv, *phSrv, logger)

	srv := priceplacement.NewServer(*feedPolicy, logger, pb.UnimplementedFeedServiceServer{})

	return App{
		cfg:                  cfg,
		productServiceServer: srv,
		Logger:               logger,
		Notify:               notify,
	}
}

func (a App) Run(ctx context.Context) error {

	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.StartGRPC(a.productServiceServer)
	})
	return grp.Wait()
}

func (a App) StartGRPC(server pb.FeedServiceServer) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to create listener")
	}

	a.grpcServer = grpc.NewServer()
	pb.RegisterFeedServiceServer(a.grpcServer, server)

	a.Logger.Info().Msg(fmt.Sprintf(" GRPC ?????????????? ???? %s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	err = a.Notify.Send(context.Background(), "Price-placements Service", fmt.Sprintf("gRPC ?????????????? ???? %v:%v", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("???????????? ???????????????? ??????????????????????")
	}

	if err := a.grpcServer.Serve(lis); err != nil {
		a.Logger.Fatal().Err(err).Msg("can't start gRPC server")
	}
	return nil
}
