package main

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/iftsoft/linker/model"
)

var errEmptyQuery = errors.New("query cannot be nil")

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
	if query == nil {
		return nil, errEmptyQuery
	}
	ms.SysState = model.SysStateStopped
	reply := ms.getSystemReply(query.Device)
	ms.log.Info("ManagerService.Terminate", "query", *query, "reply", reply)
	return &reply, nil
}

// SysHealth returns health of device application
func (ms *ManagerService) SysHealth(ctx context.Context, query *model.SystemQuery) (*model.SystemHealth, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := model.SystemHealth{
		SystemReply: ms.getSystemReply(query.Device),
		DeviceMetrics: model.DeviceMetrics{
			Moment:   time.Now().Unix(),
			Uptime:   1000,
			DevError: model.DevErrorSuccess,
			DevState: ms.DevState,
		},
	}
	ms.log.Info("ManagerService.SysHealth", "query", *query, "reply", reply)
	return &reply, nil
}

// SysStart turns device driver to initial state
func (ms *ManagerService) SysStart(ctx context.Context, query *model.ConfigUpdate) (*model.SystemDevice, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	ms.SysState = model.SysStateRunning
	reply := model.SystemDevice{
		SystemReply: ms.getSystemReply(query.Device),
		DeviceSetup: model.DeviceSetup{
			DevType:   model.DevTypePrinter,
			Supported: model.ScopeFlagSystem,
			Required:  model.ScopeFlagSystem,
		},
	}
	ms.log.Info("ManagerService.SysStart", "query", *query, "reply", reply)
	return &reply, nil
}

// SysStop gracefully deactivates device driver
func (ms *ManagerService) SysStop(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	ms.SysState = model.SysStateStopped
	reply := ms.getSystemReply(query.Device)
	ms.log.Info("ManagerService.SysStop", "query", *query, "reply", reply)
	return &reply, nil
}

// SysRestart reactivates device driver with new config
func (ms *ManagerService) SysRestart(ctx context.Context, query *model.ConfigUpdate) (*model.SystemDevice, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	ms.SysState = model.SysStateRunning
	reply := model.SystemDevice{
		SystemReply: ms.getSystemReply(query.Device),
		DeviceSetup: model.DeviceSetup{
			DevType:   model.DevTypePrinter,
			Supported: model.ScopeFlagSystem,
			Required:  model.ScopeFlagSystem,
		},
	}
	ms.log.Info("ManagerService.SysRestart", "query", *query, "reply", reply)
	return &reply, nil
}

// Cancel interrupts current operation on device
func (ms *ManagerService) Cancel(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	ms.DevState = model.DevStateReady
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.Cancel", "query", *query, "reply", reply)
	return &reply, nil
}

// Reset initializes device to initial state
func (ms *ManagerService) Reset(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	ms.DevState = model.DevStateReady
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.Reset", "query", *query, "reply", reply)
	return &reply, nil
}

// Status returns current status of device
func (ms *ManagerService) Status(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.Status", "query", *query, "reply", reply)
	return &reply, nil
}

// Execute returns result of command execution
func (ms *ManagerService) Execute(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.Execute", "query", *query, "reply", reply)
	return &reply, nil
}

// InitPrinter does primary initialization of printer before printing
func (ms *ManagerService) InitPrinter(ctx context.Context, query *model.PrinterSetup) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.InitPrinter", "query", *query, "reply", reply)
	return &reply, nil
}

// PrintPage trys to print given text on the printer
func (ms *ManagerService) PrintPage(ctx context.Context, query *model.PrinterQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.PrintPage", "query", *query, "reply", reply)
	return &reply, nil
}

// EnterCard trys to accept card in card reader device
func (ms *ManagerService) EnterCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.EnterCard", "query", *query, "reply", reply)
	return &reply, nil
}

// EjectCard trys to reject card from card reader device
func (ms *ManagerService) EjectCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.EjectCard", "query", *query, "reply", reply)
	return &reply, nil
}

// CaptureCard trys to capture card inside card reader device
func (ms *ManagerService) CaptureCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.CaptureCard", "query", *query, "reply", reply)
	return &reply, nil
}

// ReadCard trys to read card information from card
func (ms *ManagerService) ReadCard(ctx context.Context, query *model.DeviceQuery) (*model.ReadCardReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := model.ReadCardReply{
		DeviceReply: ms.getDeviceReply(query.Device),
		CardDescription: model.CardDescription{
			Device: query.Device,
		},
	}
	ms.log.Info("ManagerService.ReadCard", "query", *query, "reply", reply)
	return &reply, nil
}

// InitValidator does primary initialization of the validator
func (ms *ManagerService) InitValidator(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.InitValidator", "query", *query, "reply", reply)
	return &reply, nil
}

// DoValidate starts accepting cash from user
func (ms *ManagerService) DoValidate(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.DoValidate", "query", *query, "reply", reply)
	return &reply, nil
}

// AcceptNote puts the validated note to the cassette
func (ms *ManagerService) AcceptNote(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.AcceptNote", "query", *query, "reply", reply)
	return &reply, nil
}

// ReturnNote returns the validated note to the user
func (ms *ManagerService) ReturnNote(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.ReturnNote", "query", *query, "reply", reply)
	return &reply, nil
}

// StopValidate disables accepting new notes by validator
func (ms *ManagerService) StopValidate(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := ms.getDeviceReply(query.Device)
	ms.log.Info("ManagerService.StopValidate", "query", *query, "reply", reply)
	return &reply, nil
}

// CheckValidator returns current cassette state
func (ms *ManagerService) CheckValidator(ctx context.Context, query *model.ValidatorQuery) (*model.ValidatorStore, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := model.ValidatorStore{
		DeviceReply: ms.getDeviceReply(query.Device),
		ValidatorBatch: model.ValidatorBatch{
			Device:  query.Device,
			BatchId: 123,
			State:   model.StateCorrect,
			Details: "Cassette 32",
			Notes: model.ValidatorBox{
				model.ValidatorNote{
					Currency: model.CurrencyUSD,
					Nominal:  5,
					Count:    2,
					Amount:   10,
				},
				model.ValidatorNote{
					Currency: model.CurrencyUSD,
					Nominal:  100,
					Count:    2,
					Amount:   200,
				},
			},
		},
	}
	ms.log.Info("ManagerService.CheckValidator", "query", *query, "reply", reply)
	return &reply, nil
}

// ClearValidator clears all cassette data (settlement or reconciliation)
func (ms *ManagerService) ClearValidator(ctx context.Context, query *model.ValidatorQuery) (*model.ValidatorStore, error) {
	if query == nil {
		return nil, errEmptyQuery
	}
	reply := model.ValidatorStore{
		DeviceReply: ms.getDeviceReply(query.Device),
		ValidatorBatch: model.ValidatorBatch{
			Device:  query.Device,
			BatchId: 123,
			State:   model.StateCorrect,
			Details: "Cassette 32",
			Notes: model.ValidatorBox{
				model.ValidatorNote{
					Currency: model.CurrencyUSD,
					Nominal:  5,
					Count:    2,
					Amount:   10,
				},
				model.ValidatorNote{
					Currency: model.CurrencyUSD,
					Nominal:  100,
					Count:    2,
					Amount:   200,
				},
			},
		},
	}
	ms.log.Info("ManagerService.ClearValidator", "query", *query, "reply", reply)
	return &reply, nil
}

func (ms *ManagerService) getSystemReply(device string) model.SystemReply {
	return model.SystemReply{
		Device:   device,
		Command:  model.CmdSystemTerminate,
		Message:  "Restarted",
		SysState: ms.SysState,
		SysError: model.SysErrSuccess,
	}
}

func (ms *ManagerService) getDeviceReply(device string) model.DeviceReply {
	return model.DeviceReply{
		Device:  device,
		Command: model.CmdDeviceExecute,
		Action:  model.DevActionBarScanning,
		State:   ms.DevState,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
}
