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

// SystemHealth implements Notification about system health
func (h *SystemCallback) SystemHealth(ctx context.Context, req *srv.SystemHealthRequest) (*srv.SystemHealthResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("SystemHealthRequest is nil"))
	}

	data := req.GetData()
	reply := model.SystemHealth{
		Device:   data.GetDevice(),
		Moment:   data.GetMoment(),
		SysError: model.SysError(data.GetSysError()),
		SysState: model.SysState(data.GetSysState()),
	}
	h.log.Debug("gRPC.SystemHealth", slog.Any("reply", reply))

	err := h.api.SystemHealth(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.SystemHealth failed", slog.Any("error", err))
	}

	resp := &srv.SystemHealthResponse{}

	return resp, err
}
