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

type SystemCallback struct {
	log *slog.Logger
	api model.SystemCallback
	srv.SystemCallbackServiceServer
}

func NewSystemCallback(log *slog.Logger, api model.SystemCallback) *SystemCallback {
	return &SystemCallback{
		log: log,
		api: api,
	}
}

func (h *SystemCallback) Register(s grpc.ServiceRegistrar) {
	srv.RegisterSystemCallbackServiceServer(s, h)
}

// GreetingInfo implements Notification about system params
func (h *SystemCallback) GreetingInfo(ctx context.Context, req *srv.GreetingInfoRequest) (*srv.GreetingInfoResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("SystemGreetingInfo is nil"))
	}

	data := req.GetData()
	reply := model.GreetingInfo{
		AppName:  data.GetAppName(),
		DevName:  data.GetDevName(),
		GrpcPort: data.GetGrpcPort(),
	}
	h.log.Debug("gRPC.GreetingInfo", slog.Any("reply", reply))

	err := h.api.GreetingInfo(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.GreetingInfo failed", slog.Any("error", err))
	}

	resp := &srv.GreetingInfoResponse{}

	return resp, err
}

// SystemReply implements Notification about system reply
func (h *SystemCallback) SystemReply(ctx context.Context, req *srv.SystemReplyRequest) (*srv.SystemReplyResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("SystemReplyRequest is nil"))
	}

	data := req.GetData()
	reply := model.SystemReply{
		Device:   data.GetDevice(),
		Command:  data.GetCommand(),
		Message:  data.GetMessage(),
		SysError: model.SysError(data.GetSysError()),
		SysState: model.SysState(data.GetSysState()),
	}
	h.log.Debug("gRPC.SystemReply", slog.Any("reply", reply))

	err := h.api.SystemReply(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.SystemReply failed", slog.Any("error", err))
	}

	resp := &srv.SystemReplyResponse{}

	return resp, err
}

// SystemDevice implements Notification about system settings
func (h *SystemCallback) SystemDevice(ctx context.Context, req *srv.SystemDeviceRequest) (*srv.SystemDeviceResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("SystemDeviceRequest is nil"))
	}

	data := req.GetData()
	setup := req.GetSetup()
	reply := model.SystemDevice{
		Reply: model.SystemReply{
			Device:   data.GetDevice(),
			Command:  data.GetCommand(),
			Message:  data.GetMessage(),
			SysError: model.SysError(data.GetSysError()),
			SysState: model.SysState(data.GetSysState()),
		},
		Setup: model.SystemSetup{
			DevType:     model.DevTypeMask(setup.GetDevType()),
			Supported:   model.DevScopeMask(setup.GetSupported()),
			Required:    model.DevScopeMask(setup.GetRequired()),
			Description: setup.GetDescription(),
		},
	}
	h.log.Debug("gRPC.SystemHealth", slog.Any("reply", reply))

	err := h.api.SystemDevice(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.SystemDevice failed", slog.Any("error", err))
	}

	resp := &srv.SystemDeviceResponse{}

	return resp, err
}

// SystemHealth implements Notification about system health
func (h *SystemCallback) SystemHealth(ctx context.Context, req *srv.SystemHealthRequest) (*srv.SystemHealthResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("SystemHealthRequest is nil"))
	}

	data := req.GetData()
	metrics := req.GetMetrics()
	reply := model.SystemHealth{
		Reply: model.SystemReply{
			Device:   data.GetDevice(),
			Command:  data.GetCommand(),
			Message:  data.GetMessage(),
			SysError: model.SysError(data.GetSysError()),
			SysState: model.SysState(data.GetSysState()),
		},
		Metrics: model.SystemMetrics{
			Uptime:   metrics.GetUptime(),
			Moment:   metrics.GetMoment(),
			DevError: model.DevError(metrics.GetDevError()),
			DevState: model.DevState(metrics.GetDevState()),
		},
	}
	h.log.Debug("gRPC.SystemHealth", slog.Any("reply", reply))

	err := h.api.SystemHealth(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.SystemHealth failed", slog.Any("error", err))
	}

	resp := &srv.SystemHealthResponse{}

	return resp, err
}
