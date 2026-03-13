package device

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	srv "github.com/iftsoft/linker/gen/go/linker/device/v1"
	model "github.com/iftsoft/linker/model"
)

const (
	strMissingRequest  = "missing request"
	strConversionError = "conversion error"
	strServiceFault    = "service fault"
)

type CallbackAPI interface {
	DeviceReply(ctx context.Context, reply *model.DeviceReply) error
	ExecuteError(ctx context.Context, value *model.DeviceReply) error
	StateChanged(ctx context.Context, value *model.DeviceState) error
	ActionPrompt(ctx context.Context, value *model.DevicePrompt) error
	ReaderReturn(ctx context.Context, value *model.DeviceInform) error
}

type Callback struct {
	log *slog.Logger
	api CallbackAPI
	srv.DeviceCallbackServiceServer
}

func NewCallback(log *slog.Logger, api CallbackAPI) *Callback {
	return &Callback{
		log: log,
		api: api,
	}
}

func (h *Callback) Register(s grpc.ServiceRegistrar) {
	srv.RegisterDeviceCallbackServiceServer(s, h)
}

// DeviceReply implements Notification about device reply
func (h *Callback) DeviceReply(ctx context.Context, req *srv.DeviceReplyRequest) (*srv.DeviceReplyResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("DeviceReplyRequest is nil"))
	}

	data := req.GetData()
	reply := model.DeviceReply{
		Device:  data.GetDevice(),
		Command: data.GetCommand(),
		Action:  model.DevAction(data.GetAction()),
		State:   model.DevState(data.GetState()),
		ErrCode: model.DevError(data.GetErrCode()),
		ErrText: data.GetErrText(),
	}
	h.log.Debug("gRPC.DeviceReply", slog.Any("reply", reply))

	err := h.api.DeviceReply(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.DeviceReply failed", slog.Any("error", err))
	}

	resp := &srv.DeviceReplyResponse{}

	return resp, err
}

// ExecuteError implements Notification about execute error
func (h *Callback) ExecuteError(ctx context.Context, req *srv.ExecuteErrorRequest) (*srv.ExecuteErrorResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("ExecuteErrorRequest is nil"))
	}

	data := req.GetData()
	reply := model.DeviceReply{
		Device:  data.GetDevice(),
		Command: data.GetCommand(),
		Action:  model.DevAction(data.GetAction()),
		State:   model.DevState(data.GetState()),
		ErrCode: model.DevError(data.GetErrCode()),
		ErrText: data.GetErrText(),
	}
	h.log.Debug("gRPC.ExecuteError", slog.Any("reply", reply))

	err := h.api.ExecuteError(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.ExecuteError failed", slog.Any("error", err))
	}

	resp := &srv.ExecuteErrorResponse{}

	return resp, err
}

// StateChanged implements Notification about device state changing
func (h *Callback) StateChanged(ctx context.Context, req *srv.StateChangedRequest) (*srv.StateChangedResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("StateChangedRequest is nil"))
	}

	data := req.GetData()
	reply := model.DeviceState{
		Device:   data.GetDevice(),
		Action:   model.DevAction(data.GetAction()),
		NewState: model.DevState(data.GetNewState()),
		OldState: model.DevState(data.GetOldState()),
	}
	h.log.Debug("gRPC.StateChanged", slog.Any("reply", reply))

	err := h.api.StateChanged(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.StateChanged failed", slog.Any("error", err))
	}

	resp := &srv.StateChangedResponse{}

	return resp, err
}

// ActionPrompt implements Notification about action prompt for user
func (h *Callback) ActionPrompt(ctx context.Context, req *srv.ActionPromptRequest) (*srv.ActionPromptResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("ActionPromptRequest is nil"))
	}

	data := req.GetData()
	reply := model.DevicePrompt{
		Device: data.GetDevice(),
		Action: model.DevAction(data.GetAction()),
		Prompt: model.DevPrompt(data.GetPrompt()),
	}
	h.log.Debug("gRPC.ActionPrompt", slog.Any("reply", reply))

	err := h.api.ActionPrompt(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.ActionPrompt failed", slog.Any("error", err))
	}

	resp := &srv.ActionPromptResponse{}

	return resp, err
}

// ReaderReturn implements Notification about device reading result
func (h *Callback) ReaderReturn(ctx context.Context, req *srv.ReaderReturnRequest) (*srv.ReaderReturnResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, strMissingRequest, errors.New("ReaderReturnRequest is nil"))
	}

	data := req.GetData()
	reply := model.DeviceInform{
		Device: data.GetDevice(),
		Action: model.DevAction(data.GetAction()),
		Inform: data.GetInform(),
	}
	h.log.Debug("gRPC.ReaderReturn", slog.Any("reply", reply))

	err := h.api.ReaderReturn(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.ReaderReturn failed", slog.Any("error", err))
	}

	resp := &srv.ReaderReturnResponse{}

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
