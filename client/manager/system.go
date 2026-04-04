package manager

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	system "github.com/iftsoft/linker/gen/go/linker/system/v1"
	model "github.com/iftsoft/linker/model"
)

type SystemManagerClient struct {
	log    *slog.Logger
	system system.SystemManagerServiceClient
}

func NewSystemManagerClient(log *slog.Logger, conn *grpc.ClientConn) *SystemManagerClient {
	return &SystemManagerClient{
		log:    log,
		system: system.NewSystemManagerServiceClient(conn),
	}
}

// Terminate gracefully terminates running device application
func (c *SystemManagerClient) Terminate(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	c.log.Debug("ManagerClient.Terminate - grpc", slog.String("device", query.Device))

	input := &system.TerminateRequest{
		Query: convertSystemQuery(query),
	}
	resp, err := c.system.Terminate(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant Terminate for %s failed: %w", query.Device, err)
	}

	reply := convertSystemReply(resp.GetReply())
	return reply, nil
}

// SysInform returns health of device application
func (c *SystemManagerClient) SysInform(ctx context.Context, query *model.SystemQuery) (*model.SystemHealth, error) {
	c.log.Debug("ManagerClient.SysInform - grpc", slog.String("device", query.Device))

	input := &system.SysInformRequest{
		Query: convertSystemQuery(query),
	}
	resp, err := c.system.SysInform(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant SysInform for %s failed: %w", query.Device, err)
	}

	reply := convertSystemHealth(resp.GetReply())
	return reply, nil
}

// SysStart turns device driver to initial state
func (c *SystemManagerClient) SysStart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error) {
	c.log.Debug("ManagerClient.SysStart - grpc", slog.String("device", query.Device))

	input := &system.SysStartRequest{
		Config: convertSystemConfig(query),
	}
	resp, err := c.system.SysStart(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant SysStart for %s failed: %w", query.Device, err)
	}

	reply := convertSystemReply(resp.GetReply())
	return reply, nil
}

// SysStop gracefully deactivates device driver
func (c *SystemManagerClient) SysStop(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	c.log.Debug("ManagerClient.SysStop - grpc", slog.String("device", query.Device))

	input := &system.SysStopRequest{
		Query: convertSystemQuery(query),
	}
	resp, err := c.system.SysStop(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant SysStop for %s failed: %w", query.Device, err)
	}

	reply := convertSystemReply(resp.GetReply())
	return reply, nil
}

// SysRestart reactivates device driver with new config
func (c *SystemManagerClient) SysRestart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error) {
	c.log.Debug("ManagerClient.SysRestart - grpc", slog.String("device", query.Device))

	input := &system.SysRestartRequest{
		Config: convertSystemConfig(query),
	}
	resp, err := c.system.SysRestart(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant SysRestart for %s failed: %w", query.Device, err)
	}

	reply := convertSystemReply(resp.GetReply())
	return reply, nil
}
