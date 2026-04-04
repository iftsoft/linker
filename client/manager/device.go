package manager

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type DeviceManagerClient struct {
	log    *slog.Logger
	device device.DeviceManagerServiceClient
}

func NewDeviceManagerClient(log *slog.Logger, conn *grpc.ClientConn) *DeviceManagerClient {
	return &DeviceManagerClient{
		log:    log,
		device: device.NewDeviceManagerServiceClient(conn),
	}
}

// Cancel interrupts current operation on device
func (c *DeviceManagerClient) Cancel(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.Cancel - grpc", slog.String("device", query.Device))

	input := &device.CancelRequest{
		Query: convertDeviceQuery(query),
	}
	resp, err := c.device.Cancel(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command Cancel for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReply(resp.GetReply())
	return reply, nil
}

// Reset initializes device to initial state
func (c *DeviceManagerClient) Reset(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.Reset - grpc", slog.String("device", query.Device))

	input := &device.ResetRequest{
		Query: convertDeviceQuery(query),
	}
	resp, err := c.device.Reset(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command Reset for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReply(resp.GetReply())
	return reply, nil
}

// Status returns current status of device
func (c *DeviceManagerClient) Status(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.Status - grpc", slog.String("device", query.Device))

	input := &device.StatusRequest{
		Query: convertDeviceQuery(query),
	}
	resp, err := c.device.Status(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command Status for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReply(resp.GetReply())
	return reply, nil
}
