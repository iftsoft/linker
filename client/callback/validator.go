package callback

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type ValidatorCallbackClient struct {
	log    *slog.Logger
	device device.ValidatorCallbackServiceClient
}

func NewValidatorCallbackClient(log *slog.Logger, conn *grpc.ClientConn) *ValidatorCallbackClient {
	return &ValidatorCallbackClient{
		log:    log,
		device: device.NewValidatorCallbackServiceClient(conn),
	}
}

// NoteAccepted sends notification about new note in escrow
func (c *ValidatorCallbackClient) NoteAccepted(ctx context.Context, reply *model.ValidatorAccept) error {
	c.log.Debug("CallbackClient.NoteAccepted - grpc",
		slog.String("device", reply.Device), slog.String("note", reply.Note.String()))

	input := &device.NoteAcceptedRequest{
		Data: convertValidatorAccept(reply),
	}
	_, err := c.device.NoteAccepted(ctx, input)
	if err != nil {
		return fmt.Errorf("callback NoteAccepted for %s.%s failed: %w", reply.Device, reply.Note.String(), err)
	}

	return nil
}

// CashIsStored sends notification that note is stored to cassette
func (c *ValidatorCallbackClient) CashIsStored(ctx context.Context, reply *model.ValidatorAccept) error {
	c.log.Debug("CallbackClient.CashIsStored - grpc",
		slog.String("device", reply.Device), slog.String("note", reply.Note.String()))

	input := &device.CashIsStoredRequest{
		Data: convertValidatorAccept(reply),
	}
	_, err := c.device.CashIsStored(ctx, input)
	if err != nil {
		return fmt.Errorf("callback CashIsStored for %s.%s failed: %w", reply.Device, reply.Note.String(), err)
	}

	return nil
}

// CashReturned sends notification that note is returned to user
func (c *ValidatorCallbackClient) CashReturned(ctx context.Context, reply *model.ValidatorAccept) error {
	c.log.Debug("CallbackClient.CashReturned - grpc",
		slog.String("device", reply.Device), slog.String("note", reply.Note.String()))

	input := &device.CashReturnedRequest{
		Data: convertValidatorAccept(reply),
	}
	_, err := c.device.CashReturned(ctx, input)
	if err != nil {
		return fmt.Errorf("callback CashReturned for %s.%s failed: %w", reply.Device, reply.Note.String(), err)
	}

	return nil
}

// ValidatorStore sends notification about current cassette state
func (c *ValidatorCallbackClient) ValidatorStore(ctx context.Context, reply *model.ValidatorBatch) error {
	c.log.Debug("CallbackClient.ValidatorStore - grpc",
		slog.String("device", reply.Device), slog.String("state", reply.State.String()))

	input := &device.ValidatorStoreRequest{
		Data: convertValidatorBatch(reply),
	}
	_, err := c.device.ValidatorStore(ctx, input)
	if err != nil {
		return fmt.Errorf("callback ValidatorStore for %s.%s failed: %w", reply.Device, reply.State.String(), err)
	}

	return nil
}

func convertValidatorAccept(value *model.ValidatorAccept) *device.ValidatorAccept {
	if value == nil {
		return nil
	}
	data := &device.ValidatorAccept{
		Device: value.Device,
		Note: &device.ValidatorNote{
			Currency: uint32(value.Note.Currency),
			Nominal:  int64(value.Note.Nominal),
			Count:    uint32(value.Note.Count),
			Amount:   int64(value.Note.Amount),
		},
	}
	return data
}

func convertValidatorBatch(value *model.ValidatorBatch) *device.ValidatorBatch {
	if value == nil {
		return nil
	}
	data := &device.ValidatorBatch{
		Device:  value.Device,
		BatchId: value.BatchId,
		State:   uint32(value.State),
		Details: value.Details,
	}
	for _, note := range value.Notes {
		devNote := &device.ValidatorNote{
			Currency: uint32(note.Currency),
			Nominal:  int64(note.Nominal),
			Count:    uint32(note.Count),
			Amount:   int64(note.Amount),
		}
		data.Notes = append(data.Notes, devNote)
	}
	return data
}
