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

type DeviceCallback struct {
	log *slog.Logger
	api model.DeviceCallback
	srv.DeviceCallbackServiceServer
}

func NewDeviceCallback(log *slog.Logger, api model.DeviceCallback) *DeviceCallback {
	return &DeviceCallback{
		log: log,
		api: api,
	}
}

func (h *DeviceCallback) Register(s grpc.ServiceRegistrar) {
	srv.RegisterDeviceCallbackServiceServer(s, h)
}

// DeviceReply implements Notification about device reply
func (h *DeviceCallback) DeviceReply(ctx context.Context, req *srv.DeviceReplyRequest) (*srv.DeviceReplyResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("DeviceReplyRequest is nil"))
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
func (h *DeviceCallback) ExecuteError(ctx context.Context, req *srv.ExecuteErrorRequest) (*srv.ExecuteErrorResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("ExecuteErrorRequest is nil"))
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
func (h *DeviceCallback) StateChanged(ctx context.Context, req *srv.StateChangedRequest) (*srv.StateChangedResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("StateChangedRequest is nil"))
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
func (h *DeviceCallback) ActionPrompt(ctx context.Context, req *srv.ActionPromptRequest) (*srv.ActionPromptResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("ActionPromptRequest is nil"))
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
func (h *DeviceCallback) ReaderReturn(ctx context.Context, req *srv.ReaderReturnRequest) (*srv.ReaderReturnResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("ReaderReturnRequest is nil"))
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
