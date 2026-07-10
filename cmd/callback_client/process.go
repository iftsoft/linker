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
	// System callback
	err := ProcessGreetingInfo(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSystemReply(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessSystemHealth(ctx, log, cli)
	if err != nil {
		return err
	}
	// Device callback
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
	// Printer callback
	err = ProcessPrinterProgress(ctx, log, cli)
	if err != nil {
		return err
	}
	// Reader callback
	err = ProcessCardPosition(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessCardDescription(ctx, log, cli)
	if err != nil {
		return err
	}
	// Printer callback
	err = ProcessNoteAccepted(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessCashIsStored(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessCashReturned(ctx, log, cli)
	if err != nil {
		return err
	}
	err = ProcessValidatorStore(ctx, log, cli)
	if err != nil {
		return err
	}

	return nil
}

func ProcessGreetingInfo(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	reply := model.GreetingInfo{
		DevName:  testDevice,
		GrpcPort: 9098,
	}
	log.Info("Processing GreetingInfo", "reply", reply)
	err := cli.GreetingInfo(ctx, &reply)
	if err != nil {
		return fmt.Errorf("greeting info error: %w", err)
	}

	return nil
}

func ProcessSystemReply(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	reply := model.SystemReply{
		Device:   testDevice,
		Command:  model.CmdSystemHealth,
		Message:  "Ok",
		SysState: model.SysStateRunning,
		SysError: model.SysErrSuccess,
	}
	log.Info("Processing SystemReply", "reply", reply)
	err := cli.SystemReply(ctx, &reply)
	if err != nil {
		return fmt.Errorf("system reply error: %w", err)
	}

	return nil
}

func ProcessSystemDevice(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	reply := model.SystemDevice{
		SystemReply: model.SystemReply{
			Device:   testDevice,
			Command:  model.CmdSystemDevice,
			Message:  "Ok",
			SysState: model.SysStateRunning,
			SysError: model.SysErrSuccess,
		},
		DeviceSetup: model.DeviceSetup{
			DevType:     model.DevTypeCustom,
			Supported:   model.ScopeFlagSystem,
			Required:    model.ScopeFlagSystem,
			Description: "Device description",
		},
	}
	log.Info("Processing SystemDevice", "reply", reply)
	err := cli.SystemDevice(ctx, &reply)
	if err != nil {
		return fmt.Errorf("system device error: %w", err)
	}

	return nil
}

func ProcessSystemHealth(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	health := model.SystemHealth{
		SystemReply: model.SystemReply{
			Device:   testDevice,
			Command:  model.CmdSystemHealth,
			Message:  "Ok",
			SysState: model.SysStateRunning,
			SysError: model.SysErrSuccess,
		},
		DeviceMetrics: model.DeviceMetrics{
			Moment:   time.Now().Unix(),
			Uptime:   1000,
			DevError: model.DevErrorSuccess,
			DevState: model.DevStateWorking,
		},
	}
	log.Info("Processing SystemHealth", "health", health)
	err := cli.SystemHealth(ctx, &health)
	if err != nil {
		return fmt.Errorf("system health error: %w", err)
	}

	return nil
}

func ProcessDeviceReply(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	reply := model.DeviceReply{
		Device:  testDevice,
		Command: model.CmdDeviceStatus,
		Action:  model.DevActionBarScanning,
		State:   model.DevStateWorking,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
	log.Info("Processing DeviceReply", "reply", reply)
	err := cli.DeviceReply(ctx, &reply)
	if err != nil {
		return fmt.Errorf("device reply error: %w", err)
	}

	return nil
}

func ProcessExecuteError(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.DeviceReply{
		Device:  testDevice,
		Command: model.CmdDeviceStatus,
		Action:  model.DevActionBarScanning,
		State:   model.DevStateWorking,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
	log.Info("Processing ExecuteError", "value", value)
	err := cli.ExecuteError(ctx, &value)
	if err != nil {
		return fmt.Errorf("execuet error error: %w", err)
	}

	return nil
}

func ProcessStateChanged(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.DeviceState{
		Device:   testDevice,
		Action:   model.DevActionBarScanning,
		OldState: model.DevStateWaiting,
		NewState: model.DevStateWorking,
	}
	log.Info("Processing StateChanged", "value", value)
	err := cli.StateChanged(ctx, &value)
	if err != nil {
		return fmt.Errorf("state changed error: %w", err)
	}

	return nil
}

func ProcessActionPrompt(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.DevicePrompt{
		Device: testDevice,
		Action: model.DevActionBarScanning,
		Prompt: model.DevPromptScanBarcode,
	}
	log.Info("Processing ActionPrompt", "value", value)
	err := cli.ActionPrompt(ctx, &value)
	if err != nil {
		return fmt.Errorf("action prompt error: %w", err)
	}

	return nil
}

func ProcessReaderReturn(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.DeviceInform{
		Device: testDevice,
		Action: model.DevActionBarScanning,
		Inform: "12345678",
	}
	log.Info("Processing ReaderReturn", "value", value)
	err := cli.ReaderReturn(ctx, &value)
	if err != nil {
		return fmt.Errorf("reader return error: %w", err)
	}

	return nil
}

func ProcessPrinterProgress(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.PrinterProgress{
		Device:   testDevice,
		DocName:  "PrinterProgress",
		PageDone: 1,
		PagesAll: 2,
	}
	log.Info("Processing PrinterProgress", "value", value)
	err := cli.PrinterProgress(ctx, &value)
	if err != nil {
		return fmt.Errorf("printer progress error: %w", err)
	}

	return nil
}

func ProcessCardPosition(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.CardPosition{
		Device:   testDevice,
		Position: 1,
	}
	log.Info("Processing CardPosition", "value", value)
	err := cli.CardPosition(ctx, &value)
	if err != nil {
		return fmt.Errorf("card position error: %w", err)
	}

	return nil
}

func ProcessCardDescription(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.CardDescription{
		Device:  testDevice,
		CardPan: "1234567890123456",
		ExpDate: "12/28",
		Holder:  "TEST_USER",
		Track1:  "fasffafafsadfasddfsdafsdsdf",
		Track2:  "012345678",
		Track3:  "",
	}
	log.Info("Processing CardDescription", "value", value)
	err := cli.CardDescription(ctx, &value)
	if err != nil {
		return fmt.Errorf("card description error: %w", err)
	}

	return nil
}

func ProcessNoteAccepted(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.ValidatorAccept{
		Device: testDevice,
		Note: model.ValidatorNote{
			Currency: model.CurrencyUSD,
			Nominal:  5,
			Count:    2,
			Amount:   10,
		},
	}
	log.Info("Processing NoteAccepted", "value", value)
	err := cli.NoteAccepted(ctx, &value)
	if err != nil {
		return fmt.Errorf("note accepted error: %w", err)
	}

	return nil
}

func ProcessCashIsStored(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.ValidatorAccept{
		Device: testDevice,
		Note: model.ValidatorNote{
			Currency: model.CurrencyUSD,
			Nominal:  5,
			Count:    2,
			Amount:   10,
		},
	}
	log.Info("Processing CashIsStored", "value", value)
	err := cli.CashIsStored(ctx, &value)
	if err != nil {
		return fmt.Errorf("note is stored error: %w", err)
	}

	return nil
}

func ProcessCashReturned(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.ValidatorAccept{
		Device: testDevice,
		Note: model.ValidatorNote{
			Currency: model.CurrencyUSD,
			Nominal:  5,
			Count:    2,
			Amount:   10,
		},
	}
	log.Info("Processing CashReturned", "value", value)
	err := cli.CashReturned(ctx, &value)
	if err != nil {
		return fmt.Errorf("note returned error: %w", err)
	}

	return nil
}

func ProcessValidatorStore(ctx context.Context, log *slog.Logger, cli *callback.CallbackClient) error {
	value := model.ValidatorBatch{
		Device:  testDevice,
		BatchId: 12,
		State:   model.StateActive,
		Details: "Cassette 1",
		Notes: []model.ValidatorNote{
			{
				Currency: model.CurrencyUSD,
				Nominal:  5,
				Count:    2,
				Amount:   10,
			},
			{
				Currency: model.CurrencyUSD,
				Nominal:  100,
				Count:    3,
				Amount:   300,
			},
		},
	}
	log.Info("Processing ValidatorStore", "value", value)
	err := cli.ValidatorStore(ctx, &value)
	if err != nil {
		return fmt.Errorf("validator store error: %w", err)
	}

	return nil
}
