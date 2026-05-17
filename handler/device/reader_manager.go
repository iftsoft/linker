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

type ReaderManager struct {
	log *slog.Logger
	api model.ReaderManager
	srv.ReaderManagerServiceServer
}

func NewReaderManager(log *slog.Logger, api model.ReaderManager) *ReaderManager {
	return &ReaderManager{
		log: log,
		api: api,
	}
}

func (h *ReaderManager) Register(s grpc.ServiceRegistrar) {
	srv.RegisterReaderManagerServiceServer(s, h)
}

// EnterCard trys to accept card in card reader device
func (h *ReaderManager) EnterCard(ctx context.Context, req *srv.EnterCardRequest) (*srv.EnterCardResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("EnterCardRequest is nil"))
	}

	query := DeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.EnterCard", slog.Any("query", query))

	reply, err := h.api.EnterCard(ctx, query)
	if err != nil {
		h.log.Error("gRPC.EnterCard failed", slog.Any("error", err))
	}

	resp := &srv.EnterCardResponse{
		Reply: DeviceReplyToProto(reply),
	}

	return resp, err
}

// EjectCard trys to reject card from card reader device
func (h *ReaderManager) EjectCard(ctx context.Context, req *srv.EjectCardRequest) (*srv.EjectCardResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("EjectCardRequest is nil"))
	}

	query := DeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.EjectCard", slog.Any("query", query))

	reply, err := h.api.EjectCard(ctx, query)
	if err != nil {
		h.log.Error("gRPC.EjectCard failed", slog.Any("error", err))
	}

	resp := &srv.EjectCardResponse{
		Reply: DeviceReplyToProto(reply),
	}

	return resp, err
}

// CaptureCard trys to capture card inside card reader device
func (h *ReaderManager) CaptureCard(ctx context.Context, req *srv.CaptureCardRequest) (*srv.CaptureCardResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("CaptureCardRequest is nil"))
	}

	query := DeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.CaptureCard", slog.Any("query", query))

	reply, err := h.api.CaptureCard(ctx, query)
	if err != nil {
		h.log.Error("gRPC.CaptureCard failed", slog.Any("error", err))
	}

	resp := &srv.CaptureCardResponse{
		Reply: DeviceReplyToProto(reply),
	}

	return resp, err
}

// ReadCard trys to read card information from card
func (h *ReaderManager) ReadCard(ctx context.Context, req *srv.ReadCardRequest) (*srv.ReadCardResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("ReadCardRequest is nil"))
	}

	query := DeviceQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.ReadCard", slog.Any("query", query))

	back, err := h.api.ReadCard(ctx, query)
	if err != nil {
		h.log.Error("gRPC.ReadCard failed", slog.Any("error", err))
	}

	resp := &srv.ReadCardResponse{
		Reply: DeviceReplyToProto(back.Reply),
		Card:  CardDescriptionToProto(back.Card),
	}

	return resp, err
}

func CardDescriptionToProto(data *model.CardDescription) *srv.CardDescription {
	reply := &srv.CardDescription{
		Device:  data.Device,
		CardPan: string(data.CardPan),
		ExpDate: data.ExpDate,
		Holder:  data.Holder,
		Track1:  string(data.Track1),
		Track2:  string(data.Track2),
		Track3:  string(data.Track3),
	}
	return reply
}
