package main

import (
	"context"
	"log/slog"

	"github.com/iftsoft/linker/model"
)

type CallbackService struct {
	log *slog.Logger
}

func NewCallbackService(log *slog.Logger) *CallbackService {
	return &CallbackService{
		log: log,
	}
}

// GreetingInfo sends notification about device application
func (cs *CallbackService) GreetingInfo(ctx context.Context, reply *model.GreetingInfo) error {
	cs.log.Info("CallbackService.GreetingInfo", "reply", reply)
	return nil
}

// SystemReply sends notification about device reply
func (cs *CallbackService) SystemReply(ctx context.Context, reply *model.SystemReply) error {
	cs.log.Info("CallbackService.SystemReply", "reply", reply.String())
	return nil
}

// SystemHealth sends notification about device reply
func (cs *CallbackService) SystemHealth(ctx context.Context, reply *model.SystemHealth) error {
	cs.log.Info("CallbackService.SystemHealth", "reply", reply.String())
	return nil
}

// DeviceReply sends notification about device reply
func (cs *CallbackService) DeviceReply(ctx context.Context, reply *model.DeviceReply) error {
	cs.log.Info("CallbackService.DeviceReply", "reply", reply.String())
	return nil
}

// ExecuteError sends notification about execute error
func (cs *CallbackService) ExecuteError(ctx context.Context, value *model.DeviceReply) error {
	cs.log.Info("CallbackService.ExecuteError", "value", value.String())
	return nil
}

// StateChanged sends notification about device state changing
func (cs *CallbackService) StateChanged(ctx context.Context, value *model.DeviceState) error {
	cs.log.Info("CallbackService.StateChanged", "value", value.String())
	return nil
}

// ActionPrompt sends notification about action prompt for user
func (cs *CallbackService) ActionPrompt(ctx context.Context, value *model.DevicePrompt) error {
	cs.log.Info("CallbackService.ActionPrompt", "value", value.String())
	return nil
}

// ReaderReturn sends notification about device reading result
func (cs *CallbackService) ReaderReturn(ctx context.Context, value *model.DeviceInform) error {
	cs.log.Info("CallbackService.ReaderReturn", "value", value.String())
	return nil
}

// PrinterProgress sent notification about printing progress
func (cs *CallbackService) PrinterProgress(ctx context.Context, value *model.PrinterProgress) error {
	cs.log.Info("CallbackService.PrinterProgress", "value", value)
	return nil
}

// CardPosition sends notification about new card position
func (cs *CallbackService) CardPosition(ctx context.Context, value *model.CardPosition) error {
	cs.log.Info("CallbackService.CardPosition", "value", value)
	return nil
}

// CardDescription sends notification about card information
func (cs *CallbackService) CardDescription(ctx context.Context, value *model.CardDescription) error {
	cs.log.Info("CallbackService.CardDescription", "value", value)
	return nil
}

// NoteAccepted sends notification about new note in escrow
func (cs *CallbackService) NoteAccepted(ctx context.Context, value *model.ValidatorAccept) error {
	cs.log.Info("CallbackService.NoteAccepted", "value", value.String())
	return nil
}

// CashIsStored sends notification that note is stored to cassette
func (cs *CallbackService) CashIsStored(ctx context.Context, value *model.ValidatorAccept) error {
	cs.log.Info("CallbackService.CashIsStored", "value", value.String())
	return nil
}

// CashReturned sends notification that note is returned to user
func (cs *CallbackService) CashReturned(ctx context.Context, value *model.ValidatorAccept) error {
	cs.log.Info("CallbackService.CashReturned", "value", value.String())
	return nil
}

// ValidatorStore sends notification about current cassette state
func (cs *CallbackService) ValidatorStore(ctx context.Context, value *model.ValidatorBatch) error {
	cs.log.Info("CallbackService.ValidatorStore", "value", value.String())
	return nil
}
