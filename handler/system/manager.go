package system

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	srv "github.com/iftsoft/linker/gen/go/linker/system/v1"
	model "github.com/iftsoft/linker/model"
)

type SystemManager struct {
	log *slog.Logger
	api model.SystemManager
	srv.SystemManagerServiceServer
}

func NewSystemManager(log *slog.Logger, api model.SystemManager) *SystemManager {
	return &SystemManager{
		log: log,
		api: api,
	}
}

func (h *SystemManager) Register(s grpc.ServiceRegistrar) {
	srv.RegisterSystemManagerServiceServer(s, h)
}

// Terminate interrupts system execution
func (h *SystemManager) Terminate(ctx context.Context, req *srv.TerminateRequest) (*srv.TerminateResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("TerminateRequest is nil"))
	}

	query := convertSystemQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.Terminate", slog.Any("query", query))

	reply, err := h.api.Terminate(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Terminate failed", slog.Any("error", err))
	}

	resp := &srv.TerminateResponse{
		Reply: convertSystemReplyToProto(reply),
	}

	return resp, err
}

// SysHealth returns system health information
func (h *SystemManager) SysHealth(ctx context.Context, req *srv.SysHealthRequest) (*srv.SysHealthResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("SysHealthRequest is nil"))
	}

	query := convertSystemQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.SysHealth", slog.Any("query", query))

	reply, err := h.api.SysHealth(ctx, query)
	if err != nil {
		h.log.Error("gRPC.SysHealth failed", slog.Any("error", err))
	}

	resp := &srv.SysHealthResponse{
		Reply:   convertSystemReplyToProto(&reply.SystemReply),
		Metrics: convertDeviceMetricsToProto(&reply.DeviceMetrics),
	}

	return resp, err
}

// SysStart turns system device to initial state
func (h *SystemManager) SysStart(ctx context.Context, req *srv.SysStartRequest) (*srv.SysStartResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("SysStartRequest is nil"))
	}

	config := convertConfigUpdateToModel(req.GetConfig())
	h.log.Debug("gRPC.SysStart", slog.Any("config", config))

	reply, err := h.api.SysStart(ctx, config)
	if err != nil {
		h.log.Error("gRPC.SysStart failed", slog.Any("error", err))
	}

	resp := &srv.SysStartResponse{
		Reply: convertSystemReplyToProto(&reply.SystemReply),
		Setup: convertDeviceSetupToProto(&reply.DeviceSetup),
	}

	return resp, err
}

// SysStop deactivates system device
func (h *SystemManager) SysStop(ctx context.Context, req *srv.SysStopRequest) (*srv.SysStopResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("SysStopRequest is nil"))
	}

	query := convertSystemQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.SysStop", slog.Any("query", query))

	reply, err := h.api.SysStop(ctx, query)
	if err != nil {
		h.log.Error("gRPC.SysStop failed", slog.Any("error", err))
	}

	resp := &srv.SysStopResponse{
		Reply: convertSystemReplyToProto(reply),
	}

	return resp, err
}

// SysRestart reactivates system device
func (h *SystemManager) SysRestart(ctx context.Context, req *srv.SysRestartRequest) (*srv.SysRestartResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("SysRestartRequest is nil"))
	}

	config := convertConfigUpdateToModel(req.GetConfig())
	h.log.Debug("gRPC.SysRestart", slog.Any("config", config))

	reply, err := h.api.SysRestart(ctx, config)
	if err != nil {
		h.log.Error("gRPC.SysRestart failed", slog.Any("error", err))
	}

	resp := &srv.SysRestartResponse{
		Reply: convertSystemReplyToProto(&reply.SystemReply),
		Setup: convertDeviceSetupToProto(&reply.DeviceSetup),
	}

	return resp, err
}

func convertSystemQueryToModel(data *srv.SystemQuery) *model.SystemQuery {
	query := &model.SystemQuery{
		Device: data.GetDevice(),
	}
	return query
}

func convertConfigUpdateToModel(data *srv.ConfigUpdate) *model.ConfigUpdate {
	config := &model.ConfigUpdate{
		Device:    data.GetDevice(),
		LinkType:  data.GetLinkType(),
		PortName:  data.GetPortName(),
		VendorID:  data.GetVendorId(),
		ProductID: data.GetProductId(),
		SerialNo:  data.GetSerialNo(),
	}
	return config
}

func convertSystemReplyToProto(data *model.SystemReply) *srv.SystemReply {
	reply := &srv.SystemReply{
		Device:   data.Device,
		Command:  data.Command,
		Message:  data.Message,
		SysError: uint32(data.SysError),
		SysState: uint32(data.SysState),
	}
	return reply
}

func convertDeviceSetupToProto(data *model.DeviceSetup) *srv.DeviceSetup {
	reply := &srv.DeviceSetup{
		DevType:     uint64(data.DevType),
		Supported:   uint64(data.Supported),
		Required:    uint64(data.Required),
		Description: data.Description,
	}
	return reply
}

func convertDeviceMetricsToProto(data *model.DeviceMetrics) *srv.DeviceMetrics {
	reply := &srv.DeviceMetrics{
		Uptime:   data.Uptime,
		Moment:   data.Moment,
		DevError: uint32(data.DevError),
		DevState: uint32(data.DevState),
	}
	return reply
}
