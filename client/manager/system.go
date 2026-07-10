package manager

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	system "github.com/iftsoft/linker/gen/go/linker/system/v1"
	"github.com/iftsoft/linker/model"
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
		Query: convertSystemQueryToProto(query),
	}
	resp, err := c.system.Terminate(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant Terminate for %s failed: %w", query.Device, err)
	}

	reply := convertSystemReplyToModel(resp.GetReply())
	return &reply, nil
}

// SysHealth returns health of device application
func (c *SystemManagerClient) SysHealth(ctx context.Context, query *model.SystemQuery) (*model.SystemHealth, error) {
	c.log.Debug("ManagerClient.SysHealth - grpc", slog.String("device", query.Device))

	input := &system.SysHealthRequest{
		Query: convertSystemQueryToProto(query),
	}
	resp, err := c.system.SysHealth(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant SysHealth for %s failed: %w", query.Device, err)
	}

	reply := &model.SystemHealth{
		SystemReply:   convertSystemReplyToModel(resp.GetReply()),
		DeviceMetrics: convertDeviceMetricsToModel(resp.GetMetrics()),
	}

	return reply, nil
}

// SysStart turns device driver to initial state
func (c *SystemManagerClient) SysStart(ctx context.Context, query *model.ConfigUpdate) (*model.SystemDevice, error) {
	c.log.Debug("ManagerClient.SysStart - grpc", slog.String("device", query.Device))

	input := &system.SysStartRequest{
		Config: convertConfigUpdateToProto(query),
	}
	resp, err := c.system.SysStart(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant SysStart for %s failed: %w", query.Device, err)
	}

	reply := &model.SystemDevice{
		SystemReply: convertSystemReplyToModel(resp.GetReply()),
		DeviceSetup: convertDeviceSetupToModel(resp.GetSetup()),
	}
	return reply, nil
}

// SysStop gracefully deactivates device driver
func (c *SystemManagerClient) SysStop(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	c.log.Debug("ManagerClient.SysStop - grpc", slog.String("device", query.Device))

	input := &system.SysStopRequest{
		Query: convertSystemQueryToProto(query),
	}
	resp, err := c.system.SysStop(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant SysStop for %s failed: %w", query.Device, err)
	}

	reply := convertSystemReplyToModel(resp.GetReply())
	return &reply, nil
}

// SysRestart reactivates device driver with new config
func (c *SystemManagerClient) SysRestart(ctx context.Context, query *model.ConfigUpdate) (*model.SystemDevice, error) {
	c.log.Debug("ManagerClient.SysRestart - grpc", slog.String("device", query.Device))

	input := &system.SysRestartRequest{
		Config: convertConfigUpdateToProto(query),
	}
	resp, err := c.system.SysRestart(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("commant SysRestart for %s failed: %w", query.Device, err)
	}

	reply := &model.SystemDevice{
		SystemReply: convertSystemReplyToModel(resp.GetReply()),
		DeviceSetup: convertDeviceSetupToModel(resp.GetSetup()),
	}
	return reply, nil
}

func convertSystemQueryToProto(value *model.SystemQuery) *system.SystemQuery {
	if value == nil {
		return nil
	}
	data := &system.SystemQuery{
		Device: value.Device,
	}
	return data
}

func convertConfigUpdateToProto(value *model.ConfigUpdate) *system.ConfigUpdate {
	if value == nil {
		return nil
	}
	data := &system.ConfigUpdate{
		Device:    value.Device,
		LinkType:  value.LinkType,
		PortName:  value.PortName,
		VendorId:  value.VendorID,
		ProductId: value.ProductID,
		SerialNo:  value.SerialNo,
	}
	return data
}

func convertSystemReplyToModel(value *system.SystemReply) model.SystemReply {
	if value == nil {
		return model.SystemReply{}
	}
	data := model.SystemReply{
		Device:   value.Device,
		Command:  value.Command,
		Message:  value.Message,
		SysError: model.SysError(value.SysError),
		SysState: model.SysState(value.SysState),
	}
	return data
}

func convertDeviceSetupToModel(value *system.DeviceSetup) model.DeviceSetup {
	if value == nil {
		return model.DeviceSetup{}
	}
	data := model.DeviceSetup{
		DevType:     model.DevTypeMask(value.DevType),
		Description: value.Description,
		Supported:   model.DevScopeMask(value.Supported),
		Required:    model.DevScopeMask(value.Required),
	}
	return data
}

func convertDeviceMetricsToModel(value *system.DeviceMetrics) model.DeviceMetrics {
	if value == nil {
		return model.DeviceMetrics{}
	}
	data := model.DeviceMetrics{
		Uptime:   value.Uptime,
		Moment:   value.Moment,
		DevError: model.DevError(value.DevError),
		DevState: model.DevState(value.DevState),
	}
	return data
}
