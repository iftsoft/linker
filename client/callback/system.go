package callback

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	system "github.com/iftsoft/linker/gen/go/linker/system/v1"
	"github.com/iftsoft/linker/model"
)

type SystemCallbackClient struct {
	log    *slog.Logger
	system system.SystemCallbackServiceClient
}

func NewSystemCallbackClient(log *slog.Logger, conn *grpc.ClientConn) *SystemCallbackClient {
	return &SystemCallbackClient{
		log:    log,
		system: system.NewSystemCallbackServiceClient(conn),
	}
}

// GreetingInfo sends notification about device application
func (c *SystemCallbackClient) GreetingInfo(ctx context.Context, reply *model.GreetingInfo) error {
	c.log.Debug("CallbackClient.GreetingInfo - grpc",
		slog.String("device", reply.Device), slog.Uint64("GrpcPort", uint64(reply.GrpcPort)))

	input := &system.GreetingInfoRequest{
		Data: convertGreetingInfo(reply),
	}
	_, err := c.system.GreetingInfo(ctx, input)
	if err != nil {
		return fmt.Errorf("callback GreetingInfo for %s.%d failed: %w", reply.Device, reply.GrpcPort, err)
	}

	return nil
}

// SystemReply sends notification about device reply
func (c *SystemCallbackClient) SystemReply(ctx context.Context, reply *model.SystemReply) error {
	c.log.Debug("CallbackClient.SystemReply - grpc",
		slog.String("device", reply.Device), slog.String("command", reply.Command))

	input := &system.SystemReplyRequest{
		Data: convertSystemReply(reply),
	}
	_, err := c.system.SystemReply(ctx, input)
	if err != nil {
		return fmt.Errorf("callback SystemReply for %s.%s failed: %w", reply.Device, reply.Command, err)
	}

	return nil
}

// SystemHealth sends notification about device reply
func (c *SystemCallbackClient) SystemHealth(ctx context.Context, reply *model.SystemHealth) error {
	c.log.Debug("CallbackClient.SystemHealth - grpc",
		slog.String("device", reply.Device))

	input := &system.SystemHealthRequest{
		Data: convertSystemHealth(reply),
	}
	_, err := c.system.SystemHealth(ctx, input)
	if err != nil {
		return fmt.Errorf("callback SystemHealth for %s failed: %w", reply.Device, err)
	}

	return nil
}

func convertGreetingInfo(value *model.GreetingInfo) *system.GreetingInfo {
	if value == nil {
		return nil
	}
	data := &system.GreetingInfo{
		Device:      value.Device,
		GrpcPort:    value.GrpcPort,
		DevType:     uint64(value.DevType),
		Supported:   uint64(value.Supported),
		Required:    uint64(value.Required),
		Description: value.Description,
	}
	return data
}

func convertSystemReply(value *model.SystemReply) *system.SystemReply {
	if value == nil {
		return nil
	}
	data := &system.SystemReply{
		Device:   value.Device,
		Command:  value.Command,
		Message:  value.Message,
		SysError: uint32(value.SysError),
		SysState: uint32(value.SysState),
	}
	return data
}

func convertSystemHealth(value *model.SystemHealth) *system.SystemHealth {
	if value == nil {
		return nil
	}
	data := &system.SystemHealth{
		Device:   value.Device,
		Moment:   value.Moment,
		SysError: uint32(value.SysError),
		SysState: uint32(value.SysState),
	}
	return data
}
