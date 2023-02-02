package app

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/grpc"
	"log"
	"net"
	"price-placements-service/internal/adapters/gs"
	"price-placements-service/internal/config"
	priceplacement "price-placements-service/internal/controllers/grpc/v1"
	"price-placements-service/internal/domain/policy"
	"price-placements-service/internal/domain/service"
	"price-placements-service/pb"
)

type App struct {
	cfg                  config.Config
	grpcServer           *grpc.Server
	productServiceServer pb.FeedServiceServer
}

func NewApp(ctx context.Context, cfg config.Config) (App, error) {
	ghSrv, err := sheets.NewService(ctx, option.WithCredentialsFile(cfg.GS.ServiceKeyPath))
	if err != nil {
		log.Fatal("не могу инициализировать google sheets")
	}

	ghPhoneRepo := gs.NewPhoneRepository(*ghSrv)
	phSrv := service.NewPhoneService(ghPhoneRepo)

	ghFeedRepo := gs.NewFeedRepository(*ghSrv)
	feedSrv := service.NewFeedService(ghFeedRepo)

	feedPolicy := policy.NewFeedPolicy(*feedSrv, *phSrv)

	srv := priceplacement.NewServer(*feedPolicy, pb.UnimplementedFeedServiceServer{})

	return App{
		cfg:                  cfg,
		productServiceServer: srv,
	}, nil
}

func (a App) Run(ctx context.Context) error {

	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.StartGRPC(ctx, a.productServiceServer)
	})
	return grp.Wait()
}

func (a App) StartGRPC(ctx context.Context, server pb.FeedServiceServer) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		log.Fatal("failed to create listener")
	}

	a.grpcServer = grpc.NewServer()

	pb.RegisterFeedServiceServer(a.grpcServer, server)

	if err := a.grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	return nil
}
