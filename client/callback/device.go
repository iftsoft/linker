package callback

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type DeviceCallbackClient struct {
	log    *slog.Logger
	device device.DeviceCallbackServiceClient
}

func NewDeviceCallbackClient(log *slog.Logger, conn *grpc.ClientConn) *DeviceCallbackClient {
	return &DeviceCallbackClient{
		log:    log,
		device: device.NewDeviceCallbackServiceClient(conn),
	}
}

// DeviceReply sends notification about device reply
func (c *DeviceCallbackClient) DeviceReply(ctx context.Context, reply *model.DeviceReply) error {
	c.log.Debug("CallbackClient.DeviceReply - grpc",
		slog.String("device", reply.Device), slog.String("command", reply.Command))

	input := &device.DeviceReplyRequest{
		Data: convertDeviceReply(reply),
	}
	_, err := c.device.DeviceReply(ctx, input)
	if err != nil {
		return fmt.Errorf("callback DeviceReply for %s.%s failed: %w", reply.Device, reply.Command, err)
	}

	return nil
}

// ExecuteError sends notification about execute error
func (c *DeviceCallbackClient) ExecuteError(ctx context.Context, value *model.DeviceReply) error {
	c.log.Debug("CallbackClient.ExecuteError - grpc",
		slog.String("device", value.Device), slog.String("command", value.Command))

	input := &device.ExecuteErrorRequest{
		Data: convertDeviceReply(value),
	}
	_, err := c.device.ExecuteError(ctx, input)
	if err != nil {
		return fmt.Errorf("callback ExecuteError for %s.%s failed: %w", value.Device, value.Command, err)
	}

	return nil
}

// StateChanged sends notification about device state changing
func (c *DeviceCallbackClient) StateChanged(ctx context.Context, value *model.DeviceState) error {
	c.log.Debug("CallbackClient.StateChanged - grpc",
		slog.String("device", value.Device), slog.String("action", value.Action.String()))

	input := &device.StateChangedRequest{
		Data: convertDeviceState(value),
	}
	_, err := c.device.StateChanged(ctx, input)
	if err != nil {
		return fmt.Errorf("callback StateChanged for %s.%s failed: %w", value.Device, value.Action.String(), err)
	}

	return nil
}

// ActionPrompt sends notification about action prompt for user
func (c *DeviceCallbackClient) ActionPrompt(ctx context.Context, value *model.DevicePrompt) error {
	c.log.Debug("CallbackClient.ActionPrompt - grpc",
		slog.String("device", value.Device), slog.String("action", value.Action.String()))

	input := &device.ActionPromptRequest{
		Data: convertDevicePrompt(value),
	}
	_, err := c.device.ActionPrompt(ctx, input)
	if err != nil {
		return fmt.Errorf("callback ActionPrompt for %s.%s failed: %w", value.Device, value.Action.String(), err)
	}

	return nil
}

// ReaderReturn sends notification about device reading result
func (c *DeviceCallbackClient) ReaderReturn(ctx context.Context, value *model.DeviceInform) error {
	c.log.Debug("CallbackClient.ReaderReturn - grpc",
		slog.String("device", value.Device), slog.String("action", value.Action.String()))

	input := &device.ReaderReturnRequest{
		Data: convertDeviceInform(value),
	}
	_, err := c.device.ReaderReturn(ctx, input)
	if err != nil {
		return fmt.Errorf("callback ReaderReturn for %s.%s failed: %w", value.Device, value.Action.String(), err)
	}

	return nil
}
