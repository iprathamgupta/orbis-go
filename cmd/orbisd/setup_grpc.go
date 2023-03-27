package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sourcenetwork/orbis-go/adapter/grpcserver"
	"github.com/sourcenetwork/orbis-go/app"
	"github.com/sourcenetwork/orbis-go/config"
	"github.com/sourcenetwork/orbis-go/pkg/cleaner"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/reflection"
)

func setupGRPCServer(cfg config.GRPC, errGrp *errgroup.Group, clnr *cleaner.Cleaner, app *app.App) error {

	// Create a gRPC server object
	s := grpcserver.NewGRPCServer(app)

	reflection.Register(s)

	// Create a listener on TCP port
	lis, err := net.Listen("tcp", cfg.GRPCURL)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	// Start the gRPC Server.
	errGrp.Go(func() error {
		app.Logger().Infof("Serving gRPC on %s", cfg.GRPCURL)
		return s.Serve(lis)
	})

	// Serve REST with gRPC-Gateway.
	gwServer, err := grpcserver.NewGRPCGatewayServer(cfg)
	if err != nil {
		return fmt.Errorf("create gRPC gateway server: %w", err)
	}
	errGrp.Go(func() error {
		app.Logger().Infof("Serving gRPC-Gateway on http://%s", cfg.RESTURL)
		return gwServer.ListenAndServe()
	})

	clnr.Regster(func() {

		app.Logger().Infof("Shutting down gRPC-Gateway server")
		err := gwServer.Shutdown(context.Background())
		if err != nil {
			app.Logger().Errorf("Shutting down gRPC-Gateway, server error: %s", err)
		}

		app.Logger().Infof("Shutting down gRPC server")
		s.GracefulStop()
	})

	return nil
}
