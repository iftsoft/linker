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

	reply := convertDeviceReplyToModel(req.GetData())
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

	value := model.DeviceState{
		DeviceNotify: convertDeviceNotifyToModel(req.GetNotify()),
		StateNotify:  convertStateNotifyToModel(req.GetData()),
	}
	h.log.Debug("gRPC.StateChanged", slog.Any("value", value))

	err := h.api.StateChanged(ctx, &value)
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

	value := model.DevicePrompt{
		DeviceNotify: convertDeviceNotifyToModel(req.GetNotify()),
		PromptNotify: convertPromptNotifyToModel(req.GetData()),
	}
	h.log.Debug("gRPC.ActionPrompt", slog.Any("value", value))

	err := h.api.ActionPrompt(ctx, &value)
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

	value := model.DeviceInform{
		DeviceNotify: convertDeviceNotifyToModel(req.GetNotify()),
		InformNotify: convertInformNotifyToModel(req.GetData()),
	}
	h.log.Debug("gRPC.ReaderReturn", slog.Any("value", value))

	err := h.api.ReaderReturn(ctx, &value)
	if err != nil {
		h.log.Error("gRPC.ReaderReturn failed", slog.Any("error", err))
	}

	resp := &srv.ReaderReturnResponse{}

	return resp, err
}

func convertDeviceReplyToModel(value *srv.DeviceReply) model.DeviceReply {
	if value == nil {
		return model.DeviceReply{}
	}
	data := model.DeviceReply{
		Device:  value.GetDevice(),
		Command: value.GetCommand(),
		Action:  model.DevAction(value.GetAction()),
		State:   model.DevState(value.GetState()),
		ErrCode: model.DevError(value.GetErrCode()),
		ErrText: value.GetErrText(),
	}
	return data
}

func convertDeviceNotifyToModel(value *srv.DeviceNotify) model.DeviceNotify {
	if value == nil {
		return model.DeviceNotify{}
	}
	data := model.DeviceNotify{
		Device: value.GetDevice(),
		Action: model.DevAction(value.GetAction()),
	}
	return data
}

func convertStateNotifyToModel(value *srv.StateNotify) model.StateNotify {
	if value == nil {
		return model.StateNotify{}
	}
	data := model.StateNotify{
		NewState: model.DevState(value.GetNewState()),
		OldState: model.DevState(value.GetOldState()),
	}
	return data
}

func convertPromptNotifyToModel(value *srv.PromptNotify) model.PromptNotify {
	if value == nil {
		return model.PromptNotify{}
	}
	data := model.PromptNotify{
		Prompt: model.DevPrompt(value.GetPrompt()),
	}
	return data
}

func convertInformNotifyToModel(value *srv.InformNotify) model.InformNotify {
	if value == nil {
		return model.InformNotify{}
	}
	data := model.InformNotify{
		Inform: value.GetInform(),
	}
	return data
}
