package main

import (
	"context"
	"log/slog"

	"github.com/iftsoft/linker/model"
)

type ManagerService struct {
	log *slog.Logger
}

func NewManagerService(log *slog.Logger) *ManagerService {
	return &ManagerService{
		log: log,
	}
}

// SystemReply sends notification about device reply
func (c *ManagerService) SystemReply(ctx context.Context, reply *model.SystemReply) error {
	c.log.Info("ManagerService.SystemReply", "reply", reply)
	return nil
}

// Terminate gracefully terminates running device application
func (c *ManagerService) Terminate(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	c.log.Info("ManagerService.Terminate", "query", query)
	reply := &model.SystemReply{
		Device: query.Device,
	}
	return reply, nil
}

// SysInform returns health of device application
func (c *ManagerService) SysInform(ctx context.Context, query *model.SystemQuery) (*model.SystemHealth, error) {
	c.log.Info("ManagerService.SysInform", "query", query)
	reply := &model.SystemHealth{
		Device: query.Device,
	}
	return reply, nil
}

// SysStart turns device driver to initial state
func (c *ManagerService) SysStart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error) {
	c.log.Info("ManagerService.SysStart", "query", query)
	reply := &model.SystemReply{
		Device: query.Device,
	}
	return reply, nil
}

// SysStop gracefully deactivates device driver
func (c *ManagerService) SysStop(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	c.log.Info("ManagerService.SysStop", "query", query)
	reply := &model.SystemReply{
		Device: query.Device,
	}
	return reply, nil
}

// SysRestart reactivates device driver with new config
func (c *ManagerService) SysRestart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error) {
	c.log.Info("ManagerService.SysRestart", "query", query)
	reply := &model.SystemReply{
		Device: query.Device,
	}
	return reply, nil
}

// Cancel interrupts current operation on device
func (c *ManagerService) Cancel(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	c.log.Info("ManagerService.Cancel", "query", query)
	reply := &model.DeviceReply{
		Device: query.Device,
	}
	return reply, nil
}

// Reset initializes device to initial state
func (c *ManagerService) Reset(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	c.log.Info("ManagerService.Reset", "query", query)
	reply := &model.DeviceReply{
		Device: query.Device,
	}
	return reply, nil
}

// Status returns current status of device
func (c *ManagerService) Status(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	c.log.Info("ManagerService.Status", "query", query)
	reply := &model.DeviceReply{
		Device: query.Device,
	}
	return reply, nil
}
