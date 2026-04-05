package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/iftsoft/linker/client/callback"
	"github.com/iftsoft/linker/model"
)

const (
	testDevice = "TestDevice"
)

func ProcessTest(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	log.Info("Processing Test")
	err := ProcessSystemReply(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSystemHealth(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessDeviceReply(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessExecuteError(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessStateChanged(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessActionPrompt(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessReaderReturn(ctx, log, cli)
	if err != nil {
		return err
	}
	return nil
}

func ProcessSystemReply(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	reply := &model.SystemReply{
		Device:   testDevice,
		Command:  model.CmdSystemInform,
		Message:  "Ok",
		SysState: model.SysStateRunning,
		SysError: model.SysErrSuccess,
	}
	log.Info("Processing SystemReply", "reply", reply.String())
	err := cli.SystemReply(ctx, reply)
	if err != nil {
		return fmt.Errorf("system reply error: %w", err)
	}

	return nil
}

func ProcessSystemHealth(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	health := &model.SystemHealth{
		Device:   testDevice,
		Moment:   time.Now().Unix(),
		SysState: model.SysStateRunning,
		SysError: model.SysErrSuccess,
		Metrics: model.SystemMetrics{
			Uptime:   1000,
			DevError: model.DevErrorSuccess,
			DevState: model.DevStateWorking,
			Counts:   make(map[string]uint32),
			Totals:   make(map[string]float32),
			Topics:   make(map[string]string),
		},
	}
	log.Info("Processing SystemHealth", "health", health.String())
	err := cli.SystemHealth(ctx, health)
	if err != nil {
		return fmt.Errorf("system health error: %w", err)
	}

	return nil
}

func ProcessDeviceReply(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	reply := &model.DeviceReply{
		Device:  testDevice,
		Command: model.CmdDeviceStatus,
		Action:  model.DevActionBarScanning,
		State:   model.DevStateWorking,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
	log.Info("Processing DeviceReply", "reply", reply.String())
	err := cli.DeviceReply(ctx, reply)
	if err != nil {
		return fmt.Errorf("device reply error: %w", err)
	}

	return nil
}

func ProcessExecuteError(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := &model.DeviceReply{
		Device:  testDevice,
		Command: model.CmdDeviceStatus,
		Action:  model.DevActionBarScanning,
		State:   model.DevStateWorking,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
	log.Info("Processing ExecuteError", "value", value.String())
	err := cli.ExecuteError(ctx, value)
	if err != nil {
		return fmt.Errorf("execuet error error: %w", err)
	}

	return nil
}

func ProcessStateChanged(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := &model.DeviceState{
		Device:   testDevice,
		Action:   model.DevActionBarScanning,
		OldState: model.DevStateWaiting,
		NewState: model.DevStateWorking,
	}
	log.Info("Processing StateChanged", "value", value.String())
	err := cli.StateChanged(ctx, value)
	if err != nil {
		return fmt.Errorf("state changed error: %w", err)
	}

	return nil
}

func ProcessActionPrompt(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := &model.DevicePrompt{
		Device: testDevice,
		Action: model.DevActionBarScanning,
		Prompt: model.DevPromptScanBarcode,
	}
	log.Info("Processing ActionPrompt", "value", value.String())
	err := cli.ActionPrompt(ctx, value)
	if err != nil {
		return fmt.Errorf("action prompt error: %w", err)
	}

	return nil
}

func ProcessReaderReturn(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := &model.DeviceInform{
		Device: testDevice,
		Action: model.DevActionBarScanning,
		Inform: "12345678",
	}
	log.Info("Processing ReaderReturn", "value", value.String())
	err := cli.ReaderReturn(ctx, value)
	if err != nil {
		return fmt.Errorf("reader return error: %w", err)
	}

	return nil
}
