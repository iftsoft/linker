package callback

import (
	"context"
	"fmt"
	"log/slog"

	system "github.com/iftsoft/linker/gen/go/linker/system/v1"
	model "github.com/iftsoft/linker/model"
)

// SystemCallbackClient is the client API for SystemCallbackService service.
type SystemCallbackClient interface {
	// SystemReply sends notification about system reply
	SystemReply(ctx context.Context, reply *model.SystemReply) error
	// SystemHealth sends notification about execute error
	SystemHealth(ctx context.Context, value *model.SystemHealth) error
}

// SystemReply sends notification about device reply
func (c *CallbackClient) SystemReply(ctx context.Context, reply *model.SystemReply) error {
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
func (c *CallbackClient) SystemHealth(ctx context.Context, reply *model.SystemHealth) error {
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
