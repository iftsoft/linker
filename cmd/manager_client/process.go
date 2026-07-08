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

	// system manager
	err := ProcessSysStart(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSysHealth(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSysStop(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSysRestart(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessTerminate(ctx, log, cli)
	if err != nil {
		return err
	}
	// device manager
	err = ProcessCancel(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessReset(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessStatus(ctx, log, cli)
	if err != nil {
		return err
	}
	// printer manager
	err = ProcessInitPrinter(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessPrintPage(ctx, log, cli)
	if err != nil {
		return err
	}
	// reader manager
	err = ProcessEnterCard(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessEjectCard(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessCaptureCard(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessReadCard(ctx, log, cli)
	if err != nil {
		return err
	}
	// validator manager
	err = ProcessInitValidator(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessDoValidate(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessAcceptNote(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessReturnNote(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessStopValidate(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessCheckValidator(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessClearValidator(ctx, log, cli)
	if err != nil {
		return err
	}

	return nil
}

func ProcessTerminate(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.SystemQuery{
		Device: testDevice,
	}
	reply, err := cli.Terminate(ctx, &query)
	if err != nil {
		return fmt.Errorf("terminate app error: %w", err)
	}
	log.Info("Processing Terminate", "query", query, "reply", reply)

	return nil
}

func ProcessSysHealth(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.SystemQuery{
		Device: testDevice,
	}
	reply, err := cli.SysHealth(ctx, &query)
	if err != nil {
		return fmt.Errorf("system inform error: %w", err)
	}
	log.Info("Processing SysHealth", "query", query, "reply", reply)

	return nil
}

func ProcessSysStart(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.SystemConfig{
		Device:   testDevice,
		LinkType: 0,
		PortName: "usb",
	}
	reply, err := cli.SysStart(ctx, &query)
	if err != nil {
		return fmt.Errorf("system start error: %w", err)
	}
	log.Info("Processing SysStart", "query", query, "reply", reply)

	return nil
}

func ProcessSysStop(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.SystemQuery{
		Device: testDevice,
	}
	reply, err := cli.SysStop(ctx, &query)
	if err != nil {
		return fmt.Errorf("system stop error: %w", err)
	}
	log.Info("Processing SysStop", "query", query, "reply", reply)

	return nil
}

func ProcessSysRestart(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.SystemConfig{
		Device:   testDevice,
		LinkType: 0,
		PortName: "usb",
	}
	reply, err := cli.SysRestart(ctx, &query)
	if err != nil {
		return fmt.Errorf("system restart error: %w", err)
	}
	log.Info("Processing SysRestart", "query", query, "reply", reply)

	return nil
}

func ProcessCancel(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.Cancel(ctx, &query)
	if err != nil {
		return fmt.Errorf("device cancel error: %w", err)
	}
	log.Info("Processing Cancel", "query", query, "reply", reply)

	return nil
}

func ProcessReset(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.Reset(ctx, &query)
	if err != nil {
		return fmt.Errorf("device reset error: %w", err)
	}
	log.Info("Processing Reset", "query", query, "reply", reply)

	return nil
}

func ProcessStatus(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.Status(ctx, &query)
	if err != nil {
		return fmt.Errorf("device status error: %w", err)
	}
	log.Info("Processing Status", "query", query, "reply", reply)

	return nil
}

func ProcessInitPrinter(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.PrinterSetup{
		Device:    testDevice,
		PaperPath: 1,
		Landscape: false,
		ShowImage: 0,
	}
	reply, err := cli.InitPrinter(ctx, &query)
	if err != nil {
		return fmt.Errorf("init printer error: %w", err)
	}
	log.Info("Processing InitPrinter", "query", query, "reply", reply)

	return nil
}

func ProcessPrintPage(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.PrinterQuery{
		Device: testDevice,
		Text:   "Hello World!",
	}
	reply, err := cli.PrintPage(ctx, &query)
	if err != nil {
		return fmt.Errorf("print page error: %w", err)
	}
	log.Info("Processing PrintPage", "query", query, "reply", reply)

	return nil
}

func ProcessEnterCard(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.EnterCard(ctx, &query)
	if err != nil {
		return fmt.Errorf("enter card error: %w", err)
	}
	log.Info("Processing EnterCard", "query", query, "reply", reply)

	return nil
}

func ProcessEjectCard(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.EjectCard(ctx, &query)
	if err != nil {
		return fmt.Errorf("eject card error: %w", err)
	}
	log.Info("Processing EjectCard", "query", query, "reply", reply)

	return nil
}

func ProcessCaptureCard(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.CaptureCard(ctx, &query)
	if err != nil {
		return fmt.Errorf("capture card error: %w", err)
	}
	log.Info("Processing CaptureCard", "query", query, "reply", reply)

	return nil
}

func ProcessReadCard(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.DeviceQuery{
		Device: testDevice,
	}
	reply, err := cli.ReadCard(ctx, &query)
	if err != nil {
		return fmt.Errorf("read card error: %w", err)
	}
	log.Info("Processing ReadCard", "query", query, "reply", reply)

	return nil
}

func ProcessInitValidator(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.ValidatorQuery{
		Device:    testDevice,
		Currency:  model.CurrencyUSD,
		Operation: 123,
	}
	reply, err := cli.InitValidator(ctx, &query)
	if err != nil {
		return fmt.Errorf("init validator error: %w", err)
	}
	log.Info("Processing InitValidator", "query", query, "reply", reply)

	return nil
}

func ProcessDoValidate(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.ValidatorQuery{
		Device:    testDevice,
		Currency:  model.CurrencyUSD,
		Operation: 123,
	}
	reply, err := cli.DoValidate(ctx, &query)
	if err != nil {
		return fmt.Errorf("start validate error: %w", err)
	}
	log.Info("Processing DoValidate", "query", query, "reply", reply)

	return nil
}

func ProcessAcceptNote(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.ValidatorQuery{
		Device:    testDevice,
		Currency:  model.CurrencyUSD,
		Operation: 123,
	}
	reply, err := cli.AcceptNote(ctx, &query)
	if err != nil {
		return fmt.Errorf("accept note error: %w", err)
	}
	log.Info("Processing AcceptNote", "query", query, "reply", reply)

	return nil
}

func ProcessReturnNote(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.ValidatorQuery{
		Device:    testDevice,
		Currency:  model.CurrencyUSD,
		Operation: 123,
	}
	reply, err := cli.ReturnNote(ctx, &query)
	if err != nil {
		return fmt.Errorf("return note error: %w", err)
	}
	log.Info("Processing ReturnNote", "query", query, "reply", reply)

	return nil
}

func ProcessStopValidate(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.ValidatorQuery{
		Device:    testDevice,
		Currency:  model.CurrencyUSD,
		Operation: 123,
	}
	reply, err := cli.StopValidate(ctx, &query)
	if err != nil {
		return fmt.Errorf("stop validate error: %w", err)
	}
	log.Info("Processing StopValidate", "query", query, "reply", reply)

	return nil
}

func ProcessCheckValidator(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.ValidatorQuery{
		Device:    testDevice,
		Currency:  model.CurrencyUSD,
		Operation: 123,
	}
	reply, err := cli.CheckValidator(ctx, &query)
	if err != nil {
		return fmt.Errorf("check validator error: %w", err)
	}
	log.Info("Processing CheckValidator", "query", query, "reply", reply)

	return nil
}

func ProcessClearValidator(ctx context.Context, log *slog.Logger, cli *manager.ManagerClient) error {
	query := model.ValidatorQuery{
		Device:    testDevice,
		Currency:  model.CurrencyUSD,
		Operation: 123,
	}
	reply, err := cli.ClearValidator(ctx, &query)
	if err != nil {
		return fmt.Errorf("clear validator error: %w", err)
	}
	log.Info("Processing ClearValidator", "query", query, "reply", reply)

	return nil
}
