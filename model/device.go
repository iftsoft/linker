package model

import (
	"context"
)

const (
	CmdDeviceReply  = "DeviceReply"
	CmdExecuteError = "ExecuteError"
	CmdStateChanged = "StateChanged"
	CmdActionPrompt = "ActionPrompt"
	CmdReaderReturn = "ReaderReturn"

	CmdDeviceCancel  = "Cancel"
	CmdDeviceReset   = "Reset"
	CmdDeviceStatus  = "Status"
	CmdDeviceExecute = "Execute"
)

// DeviceManager is the client API for DeviceManagerService.
type DeviceManager interface {
	// Cancel interrupts current operation on device
	Cancel(ctx context.Context, query *DeviceQuery) (*DeviceReply, error)
	// Reset initializes device to initial state
	Reset(ctx context.Context, query *DeviceQuery) (*DeviceReply, error)
	// Status returns current status of device
	Status(ctx context.Context, query *DeviceQuery) (*DeviceReply, error)
	// Execute returns result of command execution
	Execute(ctx context.Context, query *DeviceQuery) (*DeviceReply, error)
}

// DeviceCallback is the client API for DeviceCallbackService.
type DeviceCallback interface {
	// DeviceReply sends notification about device reply
	DeviceReply(ctx context.Context, reply *DeviceReply) error
	// ExecuteError sends notification about execute error
	ExecuteError(ctx context.Context, value *DeviceReply) error
	// StateChanged sends notification about device state changing
	StateChanged(ctx context.Context, value *DeviceState) error
	// ActionPrompt sends notification about action prompt for user
	ActionPrompt(ctx context.Context, value *DevicePrompt) error
	// ReaderReturn sends notification about device reading result
	ReaderReturn(ctx context.Context, value *DeviceInform) error
}

type DeviceQuery struct {
	Device string `json:"device"`
}

type DeviceReply struct {
	Device  string    `json:"device"`
	Command string    `json:"command"`
	Action  DevAction `json:"action"`
	State   DevState  `json:"state"`
	ErrCode DevError  `json:"err_code"`
	ErrText string    `json:"err_text"`
}

type DeviceNotify struct {
	Device string    `json:"device"`
	Action DevAction `json:"action"`
}

type StateNotify struct {
	NewState DevState `json:"new_state"`
	OldState DevState `json:"old_state"`
}

type DeviceState struct {
	DeviceNotify
	StateNotify
}

type PromptNotify struct {
	Prompt DevPrompt `json:"prompt"`
}

type DevicePrompt struct {
	DeviceNotify
	PromptNotify
}

type InformNotify struct {
	Inform string `json:"inform"`
}

type DeviceInform struct {
	DeviceNotify
	InformNotify
}
