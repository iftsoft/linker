package device

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

type ManagerAPI interface {
	Terminate(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error)
	SysInform(ctx context.Context, query *model.SystemQuery) (*model.SystemHealth, error)
	SysStart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error)
	SysStop(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error)
	SysRestart(ctx context.Context, query *model.SystemConfig) (*model.SystemReply, error)
}

type Manager struct {
	log *slog.Logger
	api ManagerAPI
	srv.SystemManagerServiceServer
}

func NewManager(log *slog.Logger, api ManagerAPI) *Manager {
	return &Manager{
		log: log,
		api: api,
	}
}

func (h *Manager) Register(s grpc.ServiceRegistrar) {
	srv.RegisterSystemManagerServiceServer(s, h)
}

// Terminate interrupts current operation on device
func (h *Manager) Terminate(ctx context.Context, req *srv.TerminateRequest) (*srv.TerminateResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("TerminateRequest is nil"))
	}

	query := &model.SystemQuery{
		Device: req.GetQuery().GetDevice(),
	}
	h.log.Debug("gRPC.Terminate", slog.Any("query", query))

	reply, err := h.api.Terminate(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Cancel failed", slog.Any("error", err))
	}

	resp := &srv.TerminateResponse{
		Reply: SystemReplyToProto(reply),
	}

	return resp, err
}

// SysStart turns device to initial state
func (h *Manager) SysStart(ctx context.Context, req *srv.SysStartRequest) (*srv.SysStartResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("SysStartRequest is nil"))
	}

	query := &model.SystemConfig{
		Device: req.GetConfig().GetDevice(),
	}
	h.log.Debug("gRPC.SysStart", slog.Any("query", query))

	reply, err := h.api.SysStart(ctx, query)
	if err != nil {
		h.log.Error("gRPC.SysStart failed", slog.Any("error", err))
	}

	resp := &srv.SysStartResponse{
		Reply: SystemReplyToProto(reply),
	}

	return resp, err
}

// SysStop returns status of device
func (h *Manager) SysStop(ctx context.Context, req *srv.SysStopRequest) (*srv.SysStopResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("SysStopRequest is nil"))
	}

	query := &model.SystemQuery{
		Device: req.GetQuery().GetDevice(),
	}
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

func Serialize(value any) string {
	dump, err := json.Marshal(value)
	if err != nil {
		return "{}"
	}

	return string(dump)
}

func SystemReplyToProto(data *model.SystemReply) *srv.SystemReply {
	reply := &srv.SystemReply{
		Device:  data.Device,
		Command: data.Command,
		Message: data.Message,
		Error:   uint32(data.SysError),
		State:   uint32(data.SysState),
	}
	return reply
}
