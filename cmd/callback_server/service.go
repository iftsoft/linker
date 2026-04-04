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

// SystemReply sends notification about device reply
func (c *CallbackService) SystemReply(ctx context.Context, reply *model.SystemReply) error {
	c.log.Info("CallbackService.SystemReply", "reply", reply)
	return nil
}

// SystemHealth sends notification about device reply
func (c *CallbackService) SystemHealth(ctx context.Context, reply *model.SystemHealth) error {
	c.log.Info("CallbackService.SystemHealth", "reply", reply)
	return nil
}

// DeviceReply sends notification about device reply
func (c *CallbackService) DeviceReply(ctx context.Context, reply *model.DeviceReply) error {
	c.log.Info("CallbackService.DeviceReply", "reply", reply)
	return nil
}

// ExecuteError sends notification about execute error
func (c *CallbackService) ExecuteError(ctx context.Context, value *model.DeviceReply) error {
	c.log.Info("CallbackService.ExecuteError", "value", value)
	return nil
}

// StateChanged sends notification about device state changing
func (c *CallbackService) StateChanged(ctx context.Context, value *model.DeviceState) error {
	c.log.Info("CallbackService.StateChanged", "value", value)
	return nil
}

// ActionPrompt sends notification about action prompt for user
func (c *CallbackService) ActionPrompt(ctx context.Context, value *model.DevicePrompt) error {
	c.log.Info("CallbackService.ActionPrompt", "value", value)
	return nil
}

// ReaderReturn sends notification about device reading result
func (c *CallbackService) ReaderReturn(ctx context.Context, value *model.DeviceInform) error {
	c.log.Info("CallbackService.ReaderReturn", "value", value)
	return nil
}
