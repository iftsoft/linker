package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/iftsoft/linker/model"
)

type ManagerService struct {
	log *slog.Logger

	SysState model.SysState
	DevState model.DevState
}

func NewManagerService(log *slog.Logger) *ManagerService {
	return &ManagerService{
		log: log,

		SysState: model.SysStateUndefined,
		DevState: model.DevStateUndefined,
	}
}

// Terminate gracefully terminates running device application
func (ms *ManagerService) Terminate(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	ms.SysState = model.SysStateStopped
	reply := ms.getSystemReply(query.Device)
	ms.log.Info("ManagerService.Terminate", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// SysInform returns health of device application
func (ms *ManagerService) SysInform(ctx context.Context, query *model.SystemQuery) (*model.SystemHealth, error) {
	reply := &model.SystemHealth{
		Device:   query.Device,
		Moment:   time.Now().Unix(),
		SysState: ms.SysState,
		SysError: model.SysErrSuccess,
		Metrics: model.SystemMetrics{
			Uptime:   1000,
			DevError: model.DevErrorSuccess,
			DevState: ms.DevState,
		},
	}
	ms.log.Info("ManagerService.SysInform", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// SysStart turns device driver to initial state
func (ms *ManagerService) SysStart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error) {
	ms.SysState = model.SysStateRunning
	reply := ms.getSystemReply(query.Device)
	ms.log.Info("ManagerService.SysStart", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// SysStop gracefully deactivates device driver
func (ms *ManagerService) SysStop(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	ms.SysState = model.SysStateStopped
	reply := ms.getSystemReply(query.Device)
	ms.log.Info("ManagerService.SysStop", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// SysRestart reactivates device driver with new config
func (ms *ManagerService) SysRestart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error) {
	ms.SysState = model.SysStateRunning
	reply := ms.getSystemReply(query.Device)
	ms.log.Info("ManagerService.SysRestart", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// Cancel interrupts current operation on device
func (ms *ManagerService) Cancel(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	ms.DevState = model.DevStateReady
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.Cancel", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// Reset initializes device to initial state
func (ms *ManagerService) Reset(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	ms.DevState = model.DevStateReady
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.Reset", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// Status returns current status of device
func (ms *ManagerService) Status(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.Status", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// Execute returns result of command execution
func (ms *ManagerService) Execute(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.Execute", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// InitPrinter does primary initialization of printer before printing
func (ms *ManagerService) InitPrinter(ctx context.Context, query *model.PrinterSetup) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.InitPrinter", "query", query, "reply", reply)
	return reply, nil
}

// PrintPage trys to print given text on the printer
func (ms *ManagerService) PrintPage(ctx context.Context, query *model.PrinterQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.PrintPage", "query", query, "reply", reply)
	return reply, nil
}

// EnterCard trys to accept card in card reader device
func (ms *ManagerService) EnterCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.EnterCard", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// EjectCard trys to reject card from card reader device
func (ms *ManagerService) EjectCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.EjectCard", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// CaptureCard trys to capture card inside card reader device
func (ms *ManagerService) CaptureCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.CaptureCard", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// ReadCard trys to read card information from card
func (ms *ManagerService) ReadCard(ctx context.Context, query *model.DeviceQuery) (*model.ReadCardReply, error) {
	reply := &model.ReadCardReply{
		Reply: ms.getDeviceReply(query.Device),
		Card: &model.CardDescription{
			Device: query.Device,
		},
	}
	ms.log.Info("ManagerService.ReadCard", "query", query.String(), "reply", reply)
	return reply, nil
}

// InitValidator does primary initialization of the validator
func (ms *ManagerService) InitValidator(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.InitValidator", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// DoValidate starts accepting cash from user
func (ms *ManagerService) DoValidate(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.DoValidate", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// AcceptNote puts the validated note to the cassette
func (ms *ManagerService) AcceptNote(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.AcceptNote", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// ReturnNote returns the validated note to the user
func (ms *ManagerService) ReturnNote(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.ReturnNote", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// StopValidate disables accepting new notes by validator
func (ms *ManagerService) StopValidate(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.StopValidate", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// CheckValidator returns current cassette state
func (ms *ManagerService) CheckValidator(ctx context.Context, query *model.ValidatorQuery) (*model.ValidatorStore, error) {
	reply := &model.ValidatorStore{
		Reply: ms.getDeviceReply(query.Device),
		Batch: &model.ValidatorBatch{
			Device: query.Device,
		},
	}
	ms.log.Info("ManagerService.CheckValidator", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// ClearValidator clears all cassette data (settlement or reconciliation)
func (ms *ManagerService) ClearValidator(ctx context.Context, query *model.ValidatorQuery) (*model.ValidatorStore, error) {
	reply := &model.ValidatorStore{
		Reply: ms.getDeviceReply(query.Device),
		Batch: &model.ValidatorBatch{
			Device: query.Device,
		},
	}
	ms.log.Info("ManagerService.ClearValidator", "query", query.String(), "reply", reply.String())
	return reply, nil
}

func (ms *ManagerService) getSystemReply(device string) *model.SystemReply {
	return &model.SystemReply{
		Device:   device,
		Command:  model.CmdSystemTerminate,
		Message:  "Restarted",
		SysState: ms.SysState,
		SysError: model.SysErrSuccess,
	}
}

func (ms *ManagerService) getDeviceReply(device string) *model.DeviceReply {
	return &model.DeviceReply{
		Device:  device,
		Command: model.CmdDeviceExecute,
		Action:  model.DevActionBarScanning,
		State:   ms.DevState,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
}
