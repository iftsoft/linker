package manager

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/iftsoft/linker/grpc/client"
	"github.com/iftsoft/linker/model"
)

var ErrNotInitialized = errors.New("client is not initialized")

type ManagerClient struct {
	log    *slog.Logger
	client *client.Client
	system model.SystemManager
	device model.DeviceManager
}

func NewManagerClient(ctx context.Context, log *slog.Logger, address string) (*ManagerClient, error) {
	base, err := client.NewClient(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("can't create grpc client: %w", err)
	}
	return &ManagerClient{
		log:    log,
		client: base,
		system: NewSystemManagerClient(log, base.Connection()),
		device: NewDeviceManagerClient(log, base.Connection()),
	}, nil
}

// Close gracefully terminates grpc connection
func (c *ManagerClient) Close() error {
	return c.client.Close()
}

// Terminate gracefully terminates running device application
func (c *ManagerClient) Terminate(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	if c.system != nil {
		return c.system.Terminate(ctx, query)
	}
	return nil, ErrNotInitialized
}

// SysInform returns health of device application
func (c *ManagerClient) SysInform(ctx context.Context, query *model.SystemQuery) (*model.SystemHealth, error) {
	if c.system != nil {
		return c.system.SysInform(ctx, query)
	}
	return nil, ErrNotInitialized
}

// SysStart turns device driver to initial state
func (c *ManagerClient) SysStart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error) {
	if c.system != nil {
		return c.system.SysStart(ctx, query)
	}
	return nil, ErrNotInitialized
}

// SysStop gracefully deactivates device driver
func (c *ManagerClient) SysStop(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	if c.system != nil {
		return c.system.SysStop(ctx, query)
	}
	return nil, ErrNotInitialized
}

// SysRestart reactivates device driver with new config
func (c *ManagerClient) SysRestart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error) {
	if c.system != nil {
		return c.system.SysRestart(ctx, query)
	}
	return nil, ErrNotInitialized
}

// Cancel interrupts current operation on device
func (c *ManagerClient) Cancel(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if c.system != nil {
		return c.device.Cancel(ctx, query)
	}
	return nil, ErrNotInitialized
}

// Reset initializes device to initial state
func (c *ManagerClient) Reset(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if c.system != nil {
		return c.device.Reset(ctx, query)
	}
	return nil, ErrNotInitialized
}

// Status returns current status of device
func (c *ManagerClient) Status(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if c.system != nil {
		return c.device.Status(ctx, query)
	}
	return nil, ErrNotInitialized
}
