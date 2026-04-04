package system

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	srv "github.com/iftsoft/linker/gen/go/linker/system/v1"
	model "github.com/iftsoft/linker/model"
)

type Manager struct {
	log *slog.Logger
	api model.SystemManager
	srv.SystemManagerServiceServer
}

func NewManager(log *slog.Logger, api model.SystemManager) *Manager {
	return &Manager{
		log: log,
		api: api,
	}
}

func (h *Manager) Register(s grpc.ServiceRegistrar) {
	srv.RegisterSystemManagerServiceServer(s, h)
}

// Terminate interrupts system execution
func (h *Manager) Terminate(ctx context.Context, req *srv.TerminateRequest) (*srv.TerminateResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("TerminateRequest is nil"))
	}

	query := SystemQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.Terminate", slog.Any("query", query))

	reply, err := h.api.Terminate(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Terminate failed", slog.Any("error", err))
	}

	resp := &srv.TerminateResponse{
		Reply: SystemReplyToProto(reply),
	}

	return resp, err
}

// SysInform returns system health information
func (h *Manager) SysInform(ctx context.Context, req *srv.SysInformRequest) (*srv.SysInformResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("SysInformRequest is nil"))
	}

	query := SystemQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.SysInform", slog.Any("query", query))

	reply, err := h.api.SysInform(ctx, query)
	if err != nil {
		h.log.Error("gRPC.SysInform failed", slog.Any("error", err))
	}

	resp := &srv.SysInformResponse{
		Reply: SystemHealthToProto(reply),
	}

	return resp, err
}

// SysStart turns system device to initial state
func (h *Manager) SysStart(ctx context.Context, req *srv.SysStartRequest) (*srv.SysStartResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("SysStartRequest is nil"))
	}

	config := SystemConfigToModel(req.GetConfig())
	h.log.Debug("gRPC.SysStart", slog.Any("config", config))

	reply, err := h.api.SysStart(ctx, config)
	if err != nil {
		h.log.Error("gRPC.SysStart failed", slog.Any("error", err))
	}

	resp := &srv.SysStartResponse{
		Reply: SystemReplyToProto(reply),
	}

	return resp, err
}

// SysStop deactivates system device
func (h *Manager) SysStop(ctx context.Context, req *srv.SysStopRequest) (*srv.SysStopResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("SysStopRequest is nil"))
	}

	query := SystemQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.SysStop", slog.Any("query", query))

	reply, err := h.api.SysStop(ctx, query)
	if err != nil {
		h.log.Error("gRPC.SysStop failed", slog.Any("error", err))
	}

	resp := &srv.SysStopResponse{
		Reply: SystemReplyToProto(reply),
	}

	return resp, err
}

// SysRestart reactivates system device
func (h *Manager) SysRestart(ctx context.Context, req *srv.SysRestartRequest) (*srv.SysRestartResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("SysRestartRequest is nil"))
	}

	config := SystemConfigToModel(req.GetConfig())
	h.log.Debug("gRPC.SysRestart", slog.Any("config", config))

	reply, err := h.api.SysRestart(ctx, config)
	if err != nil {
		h.log.Error("gRPC.SysRestart failed", slog.Any("error", err))
	}

	resp := &srv.SysRestartResponse{
		Reply: SystemReplyToProto(reply),
	}

	return resp, err
}

func Serialize(value any) string {
	dump, err := json.Marshal(value)
	if err != nil {
		return "{}"
	}

	return string(dump)
}

func SystemQueryToModel(data *srv.SystemQuery) *model.SystemQuery {
	query := &model.SystemQuery{
		Device: data.GetDevice(),
	}
	return query
}

func SystemConfigToModel(data *srv.SystemConfig) *model.SystemConfig {
	config := &model.SystemConfig{
		Device:    data.GetDevice(),
		LinkType:  data.GetLinkType(),
		PortName:  data.GetPortName(),
		VendorID:  data.GetVendorId(),
		ProductID: data.GetProductId(),
	}
	return config
}

func SystemReplyToProto(data *model.SystemReply) *srv.SystemReply {
	reply := &srv.SystemReply{
		Device:   data.Device,
		Command:  data.Command,
		Message:  data.Message,
		SysError: uint32(data.SysError),
		SysState: uint32(data.SysState),
	}
	return reply
}

func SystemHealthToProto(data *model.SystemHealth) *srv.SystemHealth {
	reply := &srv.SystemHealth{
		Device:   data.Device,
		Moment:   data.Moment,
		SysError: uint32(data.SysError),
		SysState: uint32(data.SysState),
		Metrics:  &srv.SystemMetrics{},
	}
	return reply
}
