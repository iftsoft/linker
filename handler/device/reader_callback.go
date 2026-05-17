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

	data := req.GetData()
	reply := model.CardPosition{
		Device:   data.GetDevice(),
		Position: data.GetPosition(),
	}
	h.log.Debug("gRPC.CardPosition", slog.Any("reply", reply))

	err := h.api.CardPosition(ctx, &reply)
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

	data := req.GetData()
	reply := model.CardDescription{
		Device:  data.GetDevice(),
		CardPan: model.CardPAN(data.GetCardPan()),
		ExpDate: data.GetExpDate(),
		Holder:  data.GetHolder(),
		Track1:  model.Track(data.GetTrack1()),
		Track2:  model.Track(data.GetTrack2()),
		Track3:  model.Track(data.GetTrack3()),
	}
	h.log.Debug("gRPC.CardDescription", slog.Any("reply", reply))

	err := h.api.CardDescription(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.CardDescription failed", slog.Any("error", err))
	}

	resp := &srv.CardDescriptionResponse{}

	return resp, err
}
