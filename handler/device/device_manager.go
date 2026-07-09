package device

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	srv "github.com/iftsoft/linker/gen/go/linker/device/v1"
	model "github.com/iftsoft/linker/model"
)

type DeviceManager struct {
	log *slog.Logger
	api model.DeviceManager
	srv.DeviceManagerServiceServer
}

func NewDeviceManager(log *slog.Logger, api model.DeviceManager) *DeviceManager {
	return &DeviceManager{
		log: log,
		api: api,
	}
}

func (h *DeviceManager) Register(s grpc.ServiceRegistrar) {
	srv.RegisterDeviceManagerServiceServer(s, h)
}

// Cancel interrupts current operation on device
func (h *DeviceManager) Cancel(ctx context.Context, req *srv.CancelRequest) (*srv.CancelResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("CancelRequest is nil"))
	}

	query := convertDeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.Cancel", slog.Any("query", query))

	reply, err := h.api.Cancel(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Cancel failed", slog.Any("error", err))
	}

	resp := &srv.CancelResponse{
		Reply: convertDeviceReplyToProto(reply),
	}

	return resp, err
}

// Reset turns device to initial state
func (h *DeviceManager) Reset(ctx context.Context, req *srv.ResetRequest) (*srv.ResetResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("ResetRequest is nil"))
	}

	query := convertDeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.Reset", slog.Any("query", query))

	reply, err := h.api.Reset(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Reset failed", slog.Any("error", err))
	}

	resp := &srv.ResetResponse{
		Reply: convertDeviceReplyToProto(reply),
	}

	return resp, err
}

// Status returns status of device
func (h *DeviceManager) Status(ctx context.Context, req *srv.StatusRequest) (*srv.StatusResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("StatusRequest is nil"))
	}

	query := convertDeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.Status", slog.Any("query", query))

	reply, err := h.api.Status(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Status failed", slog.Any("error", err))
	}

	resp := &srv.StatusResponse{
		Reply: convertDeviceReplyToProto(reply),
	}

	return resp, err
}

// Execute returns result of command execution
func (h *DeviceManager) Execute(ctx context.Context, req *srv.ExecuteRequest) (*srv.ExecuteResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("ExecuteRequest is nil"))
	}

	query := convertDeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.Execute", slog.Any("query", query))

	reply, err := h.api.Execute(ctx, query)
	if err != nil {
		h.log.Error("gRPC.Execute failed", slog.Any("error", err))
	}

	resp := &srv.ExecuteResponse{
		Reply: convertDeviceReplyToProto(reply),
	}

	return resp, err
}

func convertDeviceQueryToModel(data *srv.DeviceQuery) *model.DeviceQuery {
	if data == nil {
		return nil
	}
	query := &model.DeviceQuery{
		Device: data.GetDevice(),
	}
	return query
}

func convertDeviceReplyToProto(data *model.DeviceReply) *srv.DeviceReply {
	if data == nil {
		return nil
	}
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
