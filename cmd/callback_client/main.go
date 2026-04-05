package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iftsoft/linker/client/callback"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	logger.Info("Start callback client")

	if err := RunClient(ctx, logger); err != nil {
		logger.Error("Callback client error", "error", err)
	}

	logger.Info("Callback client is stopped")
}

func RunClient(ctx context.Context, log *slog.Logger) error {
	// gRPC client init
	grpcCli, err := callback.NewCallbackClient(ctx, log, "127.0.0.1:9090")
	if err != nil {
		return fmt.Errorf("grpc client failed: %w", err)
	}
	defer grpcCli.Close()

	// start client loop
	period := 30 * time.Second
	timer := time.NewTimer(period)
	defer timer.Stop()
	log.Info("Starting client timer...")

	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			log.Info("Callback client context done", slog.String("error", err.Error()))
			return err

		case <-timer.C:
			err = ProcessTest(ctx, log, grpcCli)
			if err != nil {
				log.Warn("Callback client processing failed", slog.String("error", err.Error()))
			} else {
				log.Info("Callback client processing passed")
			}
			timer.Reset(period)
		}
	}
}
