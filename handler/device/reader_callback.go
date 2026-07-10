package device

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	srv "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type ReaderCallback struct {
	log *slog.Logger
	api model.ReaderCallback
	srv.ReaderCallbackServiceServer
}

func NewReaderCallback(log *slog.Logger, api model.ReaderCallback) *ReaderCallback {
	return &ReaderCallback{
		log: log,
		api: api,
	}
}

func (h *ReaderCallback) Register(s grpc.ServiceRegistrar) {
	srv.RegisterReaderCallbackServiceServer(s, h)
}

// CardPosition sends notification about new card position
func (h *ReaderCallback) CardPosition(ctx context.Context, req *srv.CardPositionRequest) (*srv.CardPositionResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("CardPositionRequest is nil"))
	}

	value := model.CardPosition{
		DeviceNotify:   convertDeviceNotifyToModel(req.GetNotify()),
		PositionNotify: convertPositionNotifyToModel(req.GetData()),
	}
	h.log.Debug("gRPC.CardPosition", slog.Any("value", value))

	err := h.api.CardPosition(ctx, &value)
	if err != nil {
		h.log.Error("gRPC.CardPosition failed", slog.Any("error", err))
	}

	resp := &srv.CardPositionResponse{}

	return resp, err
}

// CardDescription sends notification about card information
func (h *ReaderCallback) CardDescription(ctx context.Context, req *srv.CardDescriptionRequest) (*srv.CardDescriptionResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("CardDescriptionRequest is nil"))
	}

	value := model.CardDescription{
		DeviceNotify: convertDeviceNotifyToModel(req.GetNotify()),
		CardContent:  convertCardContentToModel(req.GetData()),
	}
	h.log.Debug("gRPC.CardDescription", slog.Any("value", value))

	err := h.api.CardDescription(ctx, &value)
	if err != nil {
		h.log.Error("gRPC.CardDescription failed", slog.Any("error", err))
	}

	resp := &srv.CardDescriptionResponse{}

	return resp, err
}

func convertPositionNotifyToModel(value *srv.PositionNotify) model.PositionNotify {
	if value == nil {
		return model.PositionNotify{}
	}
	data := model.PositionNotify{
		Position: value.GetPosition(),
	}
	return data
}

func convertCardContentToModel(value *srv.CardContent) model.CardContent {
	if value == nil {
		return model.CardContent{}
	}
	data := model.CardContent{
		CardPan: model.CardPAN(value.GetCardPan()),
		ExpDate: value.GetExpDate(),
		Holder:  value.GetHolder(),
		Track1:  model.Track(value.GetTrack1()),
		Track2:  model.Track(value.GetTrack2()),
		Track3:  model.Track(value.GetTrack3()),
	}
	return data
}
