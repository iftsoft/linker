package device

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	srv "github.com/iftsoft/linker/gen/go/linker/device/v1"
	model "github.com/iftsoft/linker/model"
)

type ManagerAPI interface {
	Cancel(ctx context.Context, reply *model.DeviceQuery) (*model.DeviceReply, error)
	Reset(ctx context.Context, reply *model.DeviceQuery) (*model.DeviceReply, error)
	Status(ctx context.Context, reply *model.DeviceQuery) (*model.DeviceReply, error)
}

type Manager struct {
	log *slog.Logger
	api ManagerAPI
	srv.DeviceManagerServiceServer
}

func NewManager(log *slog.Logger, api ManagerAPI) *Manager {
	return &Manager{
		log: log,
		api: api,
	}
}

func (h *Manager) Register(s grpc.ServiceRegistrar) {
	srv.RegisterDeviceManagerServiceServer(s, h)
}

// Cancel interrupts current operation on device
func (h *Manager) Cancel(ctx context.Context, req *srv.CancelRequest) (*srv.CancelResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("CancelRequest is nil"))
	}

	query := DeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.Cancel", slog.Any("query", query))

	reply, err := h.api.Cancel(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Cancel failed", slog.Any("error", err))
	}

	resp := &srv.CancelResponse{
		Reply: DeviceReplyToProto(reply),
	}

	return resp, err
}

// Reset turns device to initial state
func (h *Manager) Reset(ctx context.Context, req *srv.ResetRequest) (*srv.ResetResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("ResetRequest is nil"))
	}

	query := DeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.Reset", slog.Any("query", query))

	reply, err := h.api.Reset(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Reset failed", slog.Any("error", err))
	}

	resp := &srv.ResetResponse{
		Reply: DeviceReplyToProto(reply),
	}

	return resp, err
}

// Status returns status of device
func (h *Manager) Status(ctx context.Context, req *srv.StatusRequest) (*srv.StatusResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("StatusRequest is nil"))
	}

	query := DeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.Status", slog.Any("query", query))

	reply, err := h.api.Status(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Status failed", slog.Any("error", err))
	}

	resp := &srv.StatusResponse{
		Reply: DeviceReplyToProto(reply),
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

func DeviceQueryToModel(data *srv.DeviceQuery) *model.DeviceQuery {
	query := &model.DeviceQuery{
		Device:  data.GetDevice(),
		Timeout: data.GetTimeout(),
		Offline: data.GetOffline(),
	}
	return query
}

func DeviceReplyToProto(data *model.DeviceReply) *srv.DeviceReply {
	reply := &srv.DeviceReply{
		Device:  data.Device,
		Command: data.Command,
		Action:  uint32(data.Action),
		State:   uint32(data.State),
		ErrCode: uint32(data.ErrCode),
		ErrText: data.ErrText,
	}
	return reply
}
