package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/iftsoft/linker/client/manager"
	"github.com/iftsoft/linker/model"
)

const (
	testDevice = "TestDevice"
)

func ProcessTest(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	log.Info("Processing Test")

	err := ProcessSysStart(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessReset(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSysStop(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSysInform(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSysStart(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessCancel(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSysRestart(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessStatus(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessTerminate(ctx, log, cli)
	if err != nil {
		return err
	}

	return nil
}

func ProcessTerminate(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := &model.SystemQuery{
		Device: testDevice,
	}
	reply, err := cli.Terminate(ctx, query)
	if err != nil {
		return fmt.Errorf("terminate app error: %w", err)
	}
	log.Info("Processing Terminate", "query", query.String(), "reply", reply.String())

	return nil
}

func ProcessSysInform(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := &model.SystemQuery{
		Device: testDevice,
	}
	reply, err := cli.SysInform(ctx, query)
	if err != nil {
		return fmt.Errorf("system inform error: %w", err)
	}
	log.Info("Processing SysInform", "query", query.String(), "reply", reply.String())

	return nil
}

func ProcessSysStart(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := &model.SystemConfig{
		Device:   testDevice,
		LinkType: 0,
		PortName: "usb",
	}
	reply, err := cli.SysStart(ctx, query)
	if err != nil {
		return fmt.Errorf("system start error: %w", err)
	}
	log.Info("Processing SysStart", "query", query.String(), "reply", reply.String())

	return nil
}

func ProcessSysStop(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := &model.SystemQuery{
		Device: testDevice,
	}
	reply, err := cli.SysStop(ctx, query)
	if err != nil {
		return fmt.Errorf("system stop error: %w", err)
	}
	log.Info("Processing SysStop", "query", query.String(), "reply", reply.String())

	return nil
}

func ProcessSysRestart(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := &model.SystemConfig{
		Device:   testDevice,
		LinkType: 0,
		PortName: "usb",
	}
	reply, err := cli.SysRestart(ctx, query)
	if err != nil {
		return fmt.Errorf("system restart error: %w", err)
	}
	log.Info("Processing SysRestart", "query", query.String(), "reply", reply.String())

	return nil
}

func ProcessCancel(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := &model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.Cancel(ctx, query)
	if err != nil {
		return fmt.Errorf("device cancel error: %w", err)
	}
	log.Info("Processing Cancel", "query", query.String(), "reply", reply.String())

	return nil
}

func ProcessReset(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := &model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.Reset(ctx, query)
	if err != nil {
		return fmt.Errorf("device reset error: %w", err)
	}
	log.Info("Processing Reset", "query", query.String(), "reply", reply.String())

	return nil
}

func ProcessStatus(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := &model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.Status(ctx, query)
	if err != nil {
		return fmt.Errorf("device status error: %w", err)
	}
	log.Info("Processing Status", "query", query.String(), "reply", reply.String())

	return nil
}
