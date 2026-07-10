package callback

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/iftsoft/linker/grpc/client"
	"github.com/iftsoft/linker/model"
)

var ErrNotInitialized = errors.New("client is not initialized")

type CallbackClient struct {
	log       *slog.Logger
	client    *client.Client
	system    model.SystemCallback
	device    model.DeviceCallback
	printer   model.PrinterCallback
	reader    model.ReaderCallback
	validator model.ValidatorCallback
	mux       sync.Mutex
}

func NewCallbackClient(ctx context.Context, log *slog.Logger, address string) (*CallbackClient, error) {
	base, err := client.NewClient(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("can't create grpc client: %w", err)
	}
	return &CallbackClient{
		log:       log,
		client:    base,
		system:    NewSystemCallbackClient(log, base.Connection()),
		device:    NewDeviceCallbackClient(log, base.Connection()),
		printer:   NewPrinterCallbackClient(log, base.Connection()),
		reader:    NewReaderCallbackClient(log, base.Connection()),
		validator: NewValidatorCallbackClient(log, base.Connection()),
	}, nil
}

// Close gracefully terminates grpc connection
func (c *CallbackClient) Close() {
	err := c.client.Close()
	if err != nil {
		c.log.Error("can't close client", "error", err)
	}
}

// GreetingInfo sends notification about device application
func (c *CallbackClient) GreetingInfo(ctx context.Context, reply *model.GreetingInfo) error {
	if c.system == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.GreetingInfo(ctx, reply)
}

// SystemReply sends notification about system reply
func (c *CallbackClient) SystemReply(ctx context.Context, reply *model.SystemReply) error {
	if c.system == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.SystemReply(ctx, reply)
}

// SystemDevice sends notification about device settings
func (c *CallbackClient) SystemDevice(ctx context.Context, reply *model.SystemDevice) error {
	if c.system == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.SystemDevice(ctx, reply)
}

// SystemHealth sends notification about device health
func (c *CallbackClient) SystemHealth(ctx context.Context, reply *model.SystemHealth) error {
	if c.system == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.SystemHealth(ctx, reply)
}

// DeviceReply sends notification about device reply
func (c *CallbackClient) DeviceReply(ctx context.Context, reply *model.DeviceReply) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.DeviceReply(ctx, reply)
}

// ExecuteError sends notification about execute error
func (c *CallbackClient) ExecuteError(ctx context.Context, value *model.DeviceReply) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.ExecuteError(ctx, value)
}

// StateChanged sends notification about device state changing
func (c *CallbackClient) StateChanged(ctx context.Context, value *model.DeviceState) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.StateChanged(ctx, value)
}

// ActionPrompt sends notification about action prompt for user
func (c *CallbackClient) ActionPrompt(ctx context.Context, value *model.DevicePrompt) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.ActionPrompt(ctx, value)
}

// ReaderReturn sends notification about device reading result
func (c *CallbackClient) ReaderReturn(ctx context.Context, value *model.DeviceInform) error {
	if c.device == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.ReaderReturn(ctx, value)
}

// PrinterProgress sent notification about printing progress
func (c *CallbackClient) PrinterProgress(ctx context.Context, value *model.PrinterProgress) error {
	if c.printer == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.printer.PrinterProgress(ctx, value)
}

// CardPosition sends notification about new card position
func (c *CallbackClient) CardPosition(ctx context.Context, value *model.CardPosition) error {
	if c.reader == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.reader.CardPosition(ctx, value)
}

// CardDescription sends notification about card information
func (c *CallbackClient) CardDescription(ctx context.Context, value *model.CardDescription) error {
	if c.reader == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.reader.CardDescription(ctx, value)
}

// NoteAccepted sends notification about new note in escrow
func (c *CallbackClient) NoteAccepted(ctx context.Context, value *model.ValidatorAccept) error {
	if c.validator == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.NoteAccepted(ctx, value)
}

// CashIsStored sends notification that note is stored to cassette
func (c *CallbackClient) CashIsStored(ctx context.Context, value *model.ValidatorAccept) error {
	if c.validator == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.CashIsStored(ctx, value)
}

// CashReturned sends notification that note is returned to user
func (c *CallbackClient) CashReturned(ctx context.Context, value *model.ValidatorAccept) error {
	if c.validator == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.CashReturned(ctx, value)
}

// ValidatorStore sends notification about current cassette state
func (c *CallbackClient) ValidatorStore(ctx context.Context, value *model.ValidatorBatch) error {
	if c.validator == nil {
		return ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.ValidatorStore(ctx, value)
}
