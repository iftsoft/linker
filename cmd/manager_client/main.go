package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iftsoft/linker/client/manager"
	"github.com/iftsoft/linker/model"
)

const (
	testDevice = "TestDevice"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)
	logger.Info("Start manager client")

	if err := RunClient(ctx, logger); err != nil {
		logger.Error("Manager client error", "error", err)
	}

	logger.Info("Manager client is stopped")
}

func RunClient(ctx context.Context, log *slog.Logger) error {
	// gRPC client init
	grpcCli, err := manager.NewManagerClient(ctx, log, "127.0.0.1:9098")
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
			log.Info("Manager client context done", slog.String("error", err.Error()))
			return err

		case <-timer.C:
			err = ProcessTest(ctx, grpcCli)
			if err != nil {
				log.Warn("Manager client processing failed", slog.String("error", err.Error()))
			} else {
				log.Info("Manager client processing passed")
			}
			timer.Reset(period)
		}
	}
}

func ProcessTest(ctx context.Context, cli *manager.ManagerClient) error {
	sysRep1 := &model.SystemQuery{
		Device: testDevice,
	}
	_, err := cli.Terminate(ctx, sysRep1)
	if err != nil {
		return fmt.Errorf("system query error: %w", err)
	}

	return nil
}
