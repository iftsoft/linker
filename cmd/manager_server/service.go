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
	reply := &model.SystemReply{
		Device:   query.Device,
		Command:  model.CmdSystemTerminate,
		Message:  "Terminated",
		SysState: ms.SysState,
		SysError: model.SysErrSuccess,
	}
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
	reply := &model.SystemReply{
		Device:   query.Device,
		Command:  model.CmdSystemTerminate,
		Message:  "Started",
		SysState: ms.SysState,
		SysError: model.SysErrSuccess,
	}
	ms.log.Info("ManagerService.SysStart", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// SysStop gracefully deactivates device driver
func (ms *ManagerService) SysStop(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	ms.SysState = model.SysStateStopped
	reply := &model.SystemReply{
		Device:   query.Device,
		Command:  model.CmdSystemTerminate,
		Message:  "Stopped",
		SysState: ms.SysState,
		SysError: model.SysErrSuccess,
	}
	ms.log.Info("ManagerService.SysStop", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// SysRestart reactivates device driver with new config
func (ms *ManagerService) SysRestart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error) {
	ms.SysState = model.SysStateRunning
	reply := &model.SystemReply{
		Device:   query.Device,
		Command:  model.CmdSystemTerminate,
		Message:  "Restarted",
		SysState: ms.SysState,
		SysError: model.SysErrSuccess,
	}
	ms.log.Info("ManagerService.SysRestart", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// Cancel interrupts current operation on device
func (ms *ManagerService) Cancel(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	ms.DevState = model.DevStateReady
	reply := &model.DeviceReply{
		Device:  query.Device,
		Command: model.CmdDeviceCancel,
		Action:  model.DevActionBarScanning,
		State:   ms.DevState,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
	ms.log.Info("ManagerService.Cancel", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// Reset initializes device to initial state
func (ms *ManagerService) Reset(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	ms.DevState = model.DevStateReady
	reply := &model.DeviceReply{
		Device:  query.Device,
		Command: model.CmdDeviceReset,
		Action:  model.DevActionBarScanning,
		State:   ms.DevState,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
	ms.log.Info("ManagerService.Reset", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// Status returns current status of device
func (ms *ManagerService) Status(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	reply := &model.DeviceReply{
		Device:  query.Device,
		Command: model.CmdDeviceStatus,
		Action:  model.DevActionBarScanning,
		State:   ms.DevState,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
	ms.log.Info("ManagerService.Status", "query", query.String(), "reply", reply.String())
	return reply, nil
}

// Execute returns result of command execution
func (ms *ManagerService) Execute(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	reply := &model.DeviceReply{
		Device:  query.Device,
		Command: model.CmdDeviceExecute,
		Action:  model.DevActionBarScanning,
		State:   ms.DevState,
		ErrCode: model.DevErrorSuccess,
		ErrText: "Ok",
	}
	ms.log.Info("ManagerService.Execute", "query", query.String(), "reply", reply.String())
	return reply, nil
}
