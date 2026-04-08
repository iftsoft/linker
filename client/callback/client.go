package callback

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/iftsoft/linker/grpc/client"
	"github.com/iftsoft/linker/model"
)

var ErrNotInitialized = errors.New("client is not initialized")

type CallbackClient struct {
	log    *slog.Logger
	client *client.Client
	system model.SystemCallback
	device model.DeviceCallback
	mux    sync.Mutex
}

func NewCallbackClient(ctx context.Context, log *slog.Logger, address string) (*CallbackClient, error) {
	base, err := client.NewClient(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("can't create grpc client: %w", err)
	}
	return &CallbackClient{
		log:    log,
		client: base,
		system: NewSystemCallbackClient(log, base.Connection()),
		device: NewDeviceCallbackClient(log, base.Connection()),
	}, nil
}

// Close gracefully terminates grpc connection
func (c *CallbackClient) Close() error {
	return c.client.Close()
}

// SystemReply sends notification about device reply
func (c *CallbackClient) SystemReply(ctx context.Context, reply *model.SystemReply) error {
	if c.system == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.SystemReply(ctx, reply)
}

// SystemHealth sends notification about device reply
func (c *CallbackClient) SystemHealth(ctx context.Context, reply *model.SystemHealth) error {
	if c.system == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.SystemHealth(ctx, reply)
}

// DeviceReply sends notification about device reply
func (c *CallbackClient) DeviceReply(ctx context.Context, reply *model.DeviceReply) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.DeviceReply(ctx, reply)
}

// ExecuteError sends notification about execute error
func (c *CallbackClient) ExecuteError(ctx context.Context, value *model.DeviceReply) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.ExecuteError(ctx, value)
}

// StateChanged sends notification about device state changing
func (c *CallbackClient) StateChanged(ctx context.Context, value *model.DeviceState) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.StateChanged(ctx, value)
}

// ActionPrompt sends notification about action prompt for user
func (c *CallbackClient) ActionPrompt(ctx context.Context, value *model.DevicePrompt) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.ActionPrompt(ctx, value)
}

// ReaderReturn sends notification about device reading result
func (c *CallbackClient) ReaderReturn(ctx context.Context, value *model.DeviceInform) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.ReaderReturn(ctx, value)
}
