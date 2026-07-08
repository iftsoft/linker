package manager

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

type ManagerClient struct {
	log       *slog.Logger
	client    *client.Client
	system    model.SystemManager
	device    model.DeviceManager
	printer   model.PrinterManager
	reader    model.ReaderManager
	validator model.ValidatorManager
	mux       sync.Mutex
}

func NewManagerClient(ctx context.Context, log *slog.Logger, address string) (*ManagerClient, error) {
	base, err := client.NewClient(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("can't create grpc client: %w", err)
	}
	return &ManagerClient{
		log:       log,
		client:    base,
		system:    NewSystemManagerClient(log, base.Connection()),
		device:    NewDeviceManagerClient(log, base.Connection()),
		printer:   NewPrinterManagerClient(log, base.Connection()),
		reader:    NewReaderManagerClient(log, base.Connection()),
		validator: NewValidatorManagerClient(log, base.Connection()),
	}, nil
}

// Close gracefully terminates grpc connection
func (c *ManagerClient) Close() error {
	return c.client.Close()
}

// Terminate gracefully terminates running device application
func (c *ManagerClient) Terminate(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	if c.system == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.Terminate(ctx, query)
}

// SysHealth returns health of device application
func (c *ManagerClient) SysHealth(ctx context.Context, query *model.SystemQuery) (*model.SystemHealth, error) {
	if c.system == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.SysHealth(ctx, query)
}

// SysStart turns device driver to initial state
func (c *ManagerClient) SysStart(ctx context.Context, query *model.SystemConfig) (*model.SystemDevice, error) {
	if c.system == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.SysStart(ctx, query)
}

// SysStop gracefully deactivates device driver
func (c *ManagerClient) SysStop(ctx context.Context, query *model.SystemQuery) (*model.SystemReply, error) {
	if c.system == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.SysStop(ctx, query)
}

// SysRestart reactivates device driver with new config
func (c *ManagerClient) SysRestart(ctx context.Context, query *model.SystemConfig) (*model.SystemDevice, error) {
	if c.system == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.system.SysRestart(ctx, query)
}

// Cancel interrupts current operation on device
func (c *ManagerClient) Cancel(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if c.device == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.Cancel(ctx, query)
}

// Reset initializes device to initial state
func (c *ManagerClient) Reset(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if c.device == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.Reset(ctx, query)
}

// Status returns current status of device
func (c *ManagerClient) Status(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if c.device == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.device.Status(ctx, query)
}

// InitPrinter does primary initialization of printer before printing
func (c *ManagerClient) InitPrinter(ctx context.Context, query *model.PrinterSetup) (*model.DeviceReply, error) {
	if c.printer == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.printer.InitPrinter(ctx, query)
}

// PrintPage trys to print given text on the printer
func (c *ManagerClient) PrintPage(ctx context.Context, query *model.PrinterQuery) (*model.DeviceReply, error) {
	if c.printer == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.printer.PrintPage(ctx, query)
}

// EnterCard trys to accept card in card reader device
func (c *ManagerClient) EnterCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if c.reader == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.reader.EnterCard(ctx, query)
}

// EjectCard trys to reject card from card reader device
func (c *ManagerClient) EjectCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if c.reader == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.reader.EjectCard(ctx, query)
}

// CaptureCard trys to capture card inside card reader device
func (c *ManagerClient) CaptureCard(ctx context.Context, query *model.DeviceQuery) (*model.DeviceReply, error) {
	if c.reader == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.reader.CaptureCard(ctx, query)
}

// ReadCard trys to read card information from card
func (c *ManagerClient) ReadCard(ctx context.Context, query *model.DeviceQuery) (*model.ReadCardReply, error) {
	if c.reader == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.reader.ReadCard(ctx, query)
}

// InitValidator does primary initialization of the validator
func (c *ManagerClient) InitValidator(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if c.validator == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.InitValidator(ctx, query)
}

// DoValidate starts accepting cash from user
func (c *ManagerClient) DoValidate(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if c.validator == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.DoValidate(ctx, query)
}

// AcceptNote puts the validated note to the cassette
func (c *ManagerClient) AcceptNote(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if c.validator == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.AcceptNote(ctx, query)
}

// ReturnNote returns the validated note to the user
func (c *ManagerClient) ReturnNote(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if c.validator == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.ReturnNote(ctx, query)
}

// StopValidate disables accepting new notes by validator
func (c *ManagerClient) StopValidate(ctx context.Context, query *model.ValidatorQuery) (*model.DeviceReply, error) {
	if c.validator == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.StopValidate(ctx, query)
}

// CheckValidator returns current cassette state
func (c *ManagerClient) CheckValidator(ctx context.Context, query *model.ValidatorQuery) (*model.ValidatorStore, error) {
	if c.validator == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.CheckValidator(ctx, query)
}

// ClearValidator clears all cassette data (settlement or reconciliation)
func (c *ManagerClient) ClearValidator(ctx context.Context, query *model.ValidatorQuery) (*model.ValidatorStore, error) {
	if c.validator == nil {
		return nil, ErrNotInitialized
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.validator.ClearValidator(ctx, query)
}
