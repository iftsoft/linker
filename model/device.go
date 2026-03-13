package common

import "fmt"

const (
	CmdDeviceReply  = "DeviceReply"
	CmdExecuteError = "ExecuteError"
	CmdStateChanged = "StateChanged"
	CmdActionPrompt = "ActionPrompt"
	CmdReaderReturn = "ReaderReturn"

	CmdDeviceCancel = "Cancel"
	CmdDeviceReset  = "Reset"
	CmdDeviceStatus = "Status"
	CmdRunAction    = "RunAction"
	CmdStopAction   = "StopAction"
)

type DeviceQuery struct {
	Device  string `json:"device"`
	Timeout int64  `json:"timeout"`
	Offline bool   `json:"offline"`
}

func (dev *DeviceQuery) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Timeout = %d",
		dev.Timeout)
	return str
}

type DeviceReply struct {
	Device  string    `json:"device"`
	Command string    `json:"command"`
	Action  DevAction `json:"action"`
	State   DevState  `json:"state"`
	ErrCode DevError  `json:"err_code"`
	ErrText string    `json:"err_text"`
}

func (dev *DeviceReply) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Action = %s, State = %s, ErrCode = %s, ErrText = %s",
		dev.Action, dev.State, dev.ErrCode, dev.ErrText)
	return str
}

type DeviceState struct {
	Device   string    `json:"device"`
	Action   DevAction `json:"action"`
	NewState DevState  `json:"new_state"`
	OldState DevState  `json:"old_state"`
}

func (dev *DeviceState) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Action = %s, NewState = %s, OldState = %s",
		dev.Action, dev.NewState, dev.OldState)
	return str
}

type DevicePrompt struct {
	Device string    `json:"device"`
	Action DevAction `json:"action"`
	Prompt DevPrompt `json:"prompt"`
}

func (dev *DevicePrompt) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Action = %s, Prompt = %s",
		dev.Action, dev.Prompt)
	return str
}

type DeviceInform struct {
	Device string    `json:"device"`
	Action DevAction `json:"action"`
	Inform string    `json:"inform"`
}

func (dev *DeviceInform) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Action = %s, Inform = %s",
		dev.Action, dev.Inform)
	return str
}

type DeviceCallback interface {
	DeviceReply(reply *DeviceReply) error
	ExecuteError(value *DeviceReply) error
	StateChanged(value *DeviceState) error
	ActionPrompt(value *DevicePrompt) error
	ReaderReturn(value *DeviceInform) error
}

type DeviceManager interface {
	Cancel(name string, query *DeviceQuery) error
	Reset(name string, query *DeviceQuery) error
	Status(name string, query *DeviceQuery) error
	RunAction(name string, query *DeviceQuery) error
	StopAction(name string, query *DeviceQuery) error
}
