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
func (c *ValidatorCallbackClient) NoteAccepted(ctx context.Context, value *model.ValidatorAccept) error {
	c.log.Debug("CallbackClient.NoteAccepted - grpc",
		slog.String("device", value.Device), slog.String("note", value.Note.String()))

	input := &device.NoteAcceptedRequest{
		Notify: convertDeviceNotifyToProto(&value.DeviceNotify),
		Data:   convertAcceptNotifyToProto(&value.AcceptNotify),
	}
	_, err := c.device.NoteAccepted(ctx, input)
	if err != nil {
		return fmt.Errorf("callback NoteAccepted for %s.%s failed: %w", value.Device, value.Note.String(), err)
	}

	return nil
}

// CashIsStored sends notification that note is stored to cassette
func (c *ValidatorCallbackClient) CashIsStored(ctx context.Context, value *model.ValidatorAccept) error {
	c.log.Debug("CallbackClient.CashIsStored - grpc",
		slog.String("device", value.Device), slog.String("note", value.Note.String()))

	input := &device.CashIsStoredRequest{
		Notify: convertDeviceNotifyToProto(&value.DeviceNotify),
		Data:   convertAcceptNotifyToProto(&value.AcceptNotify),
	}
	_, err := c.device.CashIsStored(ctx, input)
	if err != nil {
		return fmt.Errorf("callback CashIsStored for %s.%s failed: %w", value.Device, value.Note.String(), err)
	}

	return nil
}

// CashReturned sends notification that note is returned to user
func (c *ValidatorCallbackClient) CashReturned(ctx context.Context, value *model.ValidatorAccept) error {
	c.log.Debug("CallbackClient.CashReturned - grpc",
		slog.String("device", value.Device), slog.String("note", value.Note.String()))

	input := &device.CashReturnedRequest{
		Notify: convertDeviceNotifyToProto(&value.DeviceNotify),
		Data:   convertAcceptNotifyToProto(&value.AcceptNotify),
	}
	_, err := c.device.CashReturned(ctx, input)
	if err != nil {
		return fmt.Errorf("callback CashReturned for %s.%s failed: %w", value.Device, value.Note.String(), err)
	}

	return nil
}

// ValidatorStore sends notification about current cassette state
func (c *ValidatorCallbackClient) ValidatorStore(ctx context.Context, value *model.ValidatorBatch) error {
	c.log.Debug("CallbackClient.ValidatorStore - grpc",
		slog.String("device", value.Device), slog.String("state", value.BatchState.String()))

	input := &device.ValidatorStoreRequest{
		Notify: convertDeviceNotifyToProto(&value.DeviceNotify),
		Data:   convertBatchContentToProto(&value.BatchContent),
	}
	_, err := c.device.ValidatorStore(ctx, input)
	if err != nil {
		return fmt.Errorf("callback ValidatorStore for %s.%s failed: %w", value.Device, value.BatchState.String(), err)
	}

	return nil
}

func convertAcceptNotifyToProto(value *model.AcceptNotify) *device.AcceptNotify {
	if value == nil {
		return nil
	}
	data := &device.AcceptNotify{
		Note: &device.ValidatorNote{
			Currency: uint32(value.Note.Currency),
			Nominal:  int64(value.Note.Nominal),
			Count:    uint32(value.Note.Count),
			Amount:   int64(value.Note.Amount),
		},
	}
	return data
}

func convertBatchContentToProto(value *model.BatchContent) *device.BatchContent {
	if value == nil {
		return nil
	}
	data := &device.BatchContent{
		BatchId:    value.BatchId,
		BatchState: uint32(value.BatchState),
		Details:    value.Details,
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
