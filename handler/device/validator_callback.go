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

type ValidatorCallback struct {
	log *slog.Logger
	api model.ValidatorCallback
	srv.ValidatorCallbackServiceServer
}

func NewValidatorCallback(log *slog.Logger, api model.ValidatorCallback) *ValidatorCallback {
	return &ValidatorCallback{
		log: log,
		api: api,
	}
}

func (h *ValidatorCallback) Register(s grpc.ServiceRegistrar) {
	srv.RegisterValidatorCallbackServiceServer(s, h)
}

// NoteAccepted sends notification about new note in escrow
func (h *ValidatorCallback) NoteAccepted(ctx context.Context, req *srv.NoteAcceptedRequest) (*srv.NoteAcceptedResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("NoteAcceptedRequest is nil"))
	}

	data := req.GetData()
	reply := model.ValidatorAccept{
		Device: data.GetDevice(),
		Note:   convertValidatorNoteToModel(data.GetNote()),
	}
	h.log.Debug("gRPC.NoteAccepted", slog.Any("reply", reply))

	err := h.api.NoteAccepted(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.NoteAccepted failed", slog.Any("error", err))
	}

	resp := &srv.NoteAcceptedResponse{}

	return resp, err
}

// CashIsStored sends notification that note is stored to cassette
func (h *ValidatorCallback) CashIsStored(ctx context.Context, req *srv.CashIsStoredRequest) (*srv.CashIsStoredResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("CashIsStoredRequest is nil"))
	}

	data := req.GetData()
	reply := model.ValidatorAccept{
		Device: data.GetDevice(),
		Note:   convertValidatorNoteToModel(data.GetNote()),
	}
	h.log.Debug("gRPC.CashIsStored", slog.Any("reply", reply))

	err := h.api.CashIsStored(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.CashIsStored failed", slog.Any("error", err))
	}

	resp := &srv.CashIsStoredResponse{}

	return resp, err
}

// CashReturned sends notification that note is returned to user
func (h *ValidatorCallback) CashReturned(ctx context.Context, req *srv.CashReturnedRequest) (*srv.CashReturnedResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("CashReturnedRequest is nil"))
	}

	data := req.GetData()
	reply := model.ValidatorAccept{
		Device: data.GetDevice(),
		Note:   convertValidatorNoteToModel(data.GetNote()),
	}
	h.log.Debug("gRPC.CashReturned", slog.Any("reply", reply))

	err := h.api.CashReturned(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.CashReturned failed", slog.Any("error", err))
	}

	resp := &srv.CashReturnedResponse{}

	return resp, err
}

// ValidatorStore sends notification about current cassette state
func (h *ValidatorCallback) ValidatorStore(ctx context.Context, req *srv.ValidatorStoreRequest) (*srv.ValidatorStoreResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("ValidatorStoreRequest is nil"))
	}

	reply := convertValidatorBatchToModel(req.GetData())
	h.log.Debug("gRPC.ValidatorStore", slog.Any("reply", reply))

	err := h.api.ValidatorStore(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.ValidatorStore failed", slog.Any("error", err))
	}

	resp := &srv.ValidatorStoreResponse{}

	return resp, err
}

func convertValidatorNoteToModel(note *srv.ValidatorNote) model.ValidatorNote {
	return model.ValidatorNote{
		Currency: model.Currency(note.GetCurrency()),
		Nominal:  model.Amount(note.GetNominal()),
		Count:    model.Counter(note.GetCount()),
		Amount:   model.Amount(note.GetAmount()),
	}
}

func convertValidatorBatchToModel(data *srv.ValidatorBatch) model.ValidatorBatch {
	batch := model.ValidatorBatch{
		Device:  data.GetDevice(),
		BatchId: data.GetBatchId(),
		State:   model.BatchState(data.GetState()),
		Details: data.GetDetails(),
	}
	for _, note := range data.GetNotes() {
		batch.Notes = append(batch.Notes, convertValidatorNoteToModel(note))
	}
	return batch
}
