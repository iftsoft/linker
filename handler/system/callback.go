package system

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	srv "github.com/iftsoft/linker/gen/go/linker/system/v1"
	model "github.com/iftsoft/linker/model"
)

const (
	strMissingRequest  = "missing request"
	strConversionError = "conversion error"
	strServiceFault    = "service fault"
)

type Callback struct {
	log *slog.Logger
	api model.SystemCallback
	srv.SystemCallbackServiceServer
}

func NewCallback(log *slog.Logger, api model.SystemCallback) *Callback {
	return &Callback{
		log: log,
		api: api,
	}
}

func (h *Callback) Register(s grpc.ServiceRegistrar) {
	srv.RegisterSystemCallbackServiceServer(s, h)
}

// SystemReply implements Notification about system reply
func (h *Callback) SystemReply(ctx context.Context, req *srv.SystemReplyRequest) (*srv.SystemReplyResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("SystemReplyRequest is nil"))
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
func (h *Callback) SystemHealth(ctx context.Context, req *srv.SystemHealthRequest) (*srv.SystemHealthResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("SystemHealthRequest is nil"))
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

func MakeErrorWithDetails(code codes.Code, msg string, e error) error {
	details := &errdetails.ErrorInfo{
		Reason: e.Error(),
		Domain: "apis.base-cms",
	}

	sts, err := status.New(code, msg).WithDetails(details)
	if err != nil {
		sts = status.New(codes.Internal, err.Error())
	}

	return sts.Err()
}
