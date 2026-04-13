package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/iftsoft/linker/handler"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	logger.Info("Start manager server")

	if err := Run(ctx, logger); err != nil {
		logger.Error("Manager server error", "error", err)
	}

	logger.Info("Manager server is stopped")
}

func Run(ctx context.Context, log *slog.Logger) error {
	// gRPC Server init
	address := "127.0.0.1:9098"
	service := NewManagerService(log)
	grpcSrv := handler.NewManagerServer(log, address, service)
	if grpcSrv == nil {
		return errors.New("grpc server is nil")
	}
	defer grpcSrv.Shutdown()

	// gRPC Server start
	go func() {
		if err := grpcSrv.Start(); err != nil {
			log.Error("GRPC server start failed", "error", err)
		}
	}()

	log.Info("Application is running now, press Ctrl+C to shutdown")
	<-ctx.Done()

	return nil
}
