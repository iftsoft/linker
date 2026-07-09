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

type ValidatorManager struct {
	log *slog.Logger
	api model.ValidatorManager
	srv.ValidatorManagerServiceServer
}

func NewValidatorManager(log *slog.Logger, api model.ValidatorManager) *ValidatorManager {
	return &ValidatorManager{
		log: log,
		api: api,
	}
}

func (h *ValidatorManager) Register(s grpc.ServiceRegistrar) {
	srv.RegisterValidatorManagerServiceServer(s, h)
}

// InitValidator does primary initialization of the validator
func (h *ValidatorManager) InitValidator(ctx context.Context, req *srv.InitValidatorRequest) (*srv.InitValidatorResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("InitValidatorRequest is nil"))
	}

	query := convertValidatorQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.InitValidator", slog.Any("query", query))

	reply, err := h.api.InitValidator(ctx, query)
	if err != nil {
		h.log.Error("gRPC.InitValidator failed", slog.Any("error", err))
	}

	resp := &srv.InitValidatorResponse{
		Reply: convertDeviceReplyToProto(reply),
	}

	return resp, err
}

// DoValidate starts accepting cash from user
func (h *ValidatorManager) DoValidate(ctx context.Context, req *srv.DoValidateRequest) (*srv.DoValidateResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("DoValidateRequest is nil"))
	}

	query := convertValidatorQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.DoValidate", slog.Any("query", query))

	reply, err := h.api.DoValidate(ctx, query)
	if err != nil {
		h.log.Error("gRPC.DoValidate failed", slog.Any("error", err))
	}

	resp := &srv.DoValidateResponse{
		Reply: convertDeviceReplyToProto(reply),
	}

	return resp, err
}

// AcceptNote puts the validated note to the cassette
func (h *ValidatorManager) AcceptNote(ctx context.Context, req *srv.AcceptNoteRequest) (*srv.AcceptNoteResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("AcceptNoteRequest is nil"))
	}

	query := convertValidatorQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.AcceptNote", slog.Any("query", query))

	reply, err := h.api.AcceptNote(ctx, query)
	if err != nil {
		h.log.Error("gRPC.AcceptNote failed", slog.Any("error", err))
	}

	resp := &srv.AcceptNoteResponse{
		Reply: convertDeviceReplyToProto(reply),
	}

	return resp, err
}

// ReturnNote returns the validated note to the user
func (h *ValidatorManager) ReturnNote(ctx context.Context, req *srv.ReturnNoteRequest) (*srv.ReturnNoteResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("ReturnNoteRequest is nil"))
	}

	query := convertValidatorQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.ReturnNote", slog.Any("query", query))

	reply, err := h.api.ReturnNote(ctx, query)
	if err != nil {
		h.log.Error("gRPC.ReturnNote failed", slog.Any("error", err))
	}

	resp := &srv.ReturnNoteResponse{
		Reply: convertDeviceReplyToProto(reply),
	}

	return resp, err
}

// StopValidate disables accepting new notes by validator
func (h *ValidatorManager) StopValidate(ctx context.Context, req *srv.StopValidateRequest) (*srv.StopValidateResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("StopValidateRequest is nil"))
	}

	query := convertValidatorQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.StopValidate", slog.Any("query", query))

	reply, err := h.api.StopValidate(ctx, query)
	if err != nil {
		h.log.Error("gRPC.StopValidate failed", slog.Any("error", err))
	}

	resp := &srv.StopValidateResponse{
		Reply: convertDeviceReplyToProto(reply),
	}

	return resp, err
}

// CheckValidator returns current cassette state
func (h *ValidatorManager) CheckValidator(ctx context.Context, req *srv.CheckValidatorRequest) (*srv.CheckValidatorResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("CheckValidatorRequest is nil"))
	}

	query := convertValidatorQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.CheckValidator", slog.Any("query", query))

	store, err := h.api.CheckValidator(ctx, query)
	if err != nil {
		h.log.Error("gRPC.CheckValidator failed", slog.Any("error", err))
	}

	resp := &srv.CheckValidatorResponse{
		Reply: convertDeviceReplyToProto(&store.DeviceReply),
		Batch: convertValidatorBatchToProto(&store.ValidatorBatch),
	}

	return resp, err
}

// ClearValidator returns current cassette state
func (h *ValidatorManager) ClearValidator(ctx context.Context, req *srv.ClearValidatorRequest) (*srv.ClearValidatorResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("ClearValidatorRequest is nil"))
	}

	query := convertValidatorQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.ClearValidator", slog.Any("query", query))

	store, err := h.api.ClearValidator(ctx, query)
	if err != nil {
		h.log.Error("gRPC.ClearValidator failed", slog.Any("error", err))
	}

	resp := &srv.ClearValidatorResponse{
		Reply: convertDeviceReplyToProto(&store.DeviceReply),
		Batch: convertValidatorBatchToProto(&store.ValidatorBatch),
	}

	return resp, err
}

func convertValidatorQueryToModel(data *srv.ValidatorQuery) *model.ValidatorQuery {
	if data == nil {
		return nil
	}
	query := &model.ValidatorQuery{
		Device:    data.GetDevice(),
		Currency:  model.Currency(data.GetCurrency()),
		Operation: data.GetOperation(),
	}
	return query
}

func convertValidatorNoteToProto(note model.ValidatorNote) *srv.ValidatorNote {
	reply := &srv.ValidatorNote{
		Currency: uint32(note.Currency),
		Nominal:  int64(note.Nominal),
		Count:    uint32(note.Count),
		Amount:   int64(note.Amount),
	}
	return reply
}

func convertValidatorBatchToProto(data *model.ValidatorBatch) *srv.ValidatorBatch {
	if data == nil {
		return nil
	}
	reply := &srv.ValidatorBatch{
		Device:  data.Device,
		BatchId: data.BatchId,
		State:   uint32(data.State),
		Details: data.Details,
		Notes:   make([]*srv.ValidatorNote, len(data.Notes)),
	}
	for _, note := range data.Notes {
		item := convertValidatorNoteToProto(note)
		reply.Notes = append(reply.Notes, item)
	}
	return reply
}
