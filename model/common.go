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
	Device  string        `json:"device"`
	Command string        `json:"command"`
	Action  EnumDevAction `json:"action"`
	State   EnumDevState  `json:"state"`
	ErrCode EnumDevError  `json:"err_code"`
	ErrText string        `json:"err_text"`
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
	Action   EnumDevAction `json:"action"`
	NewState EnumDevState  `json:"new_state"`
	OldState EnumDevState  `json:"old_state"`
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
	Action EnumDevAction `json:"action"`
	Prompt EnumDevPrompt `json:"prompt"`
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
	Action EnumDevAction `json:"action"`
	Inform string        `json:"inform"`
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
	DeviceReply(name string, reply *DeviceReply) error
	ExecuteError(name string, value *DeviceReply) error
	StateChanged(name string, value *DeviceState) error
	ActionPrompt(name string, value *DevicePrompt) error
	ReaderReturn(name string, value *DeviceInform) error
}

type DeviceManager interface {
	Cancel(name string, query *DeviceQuery) error
	Reset(name string, query *DeviceQuery) error
	Status(name string, query *DeviceQuery) error
	RunAction(name string, query *DeviceQuery) error
	StopAction(name string, query *DeviceQuery) error
}
