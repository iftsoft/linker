package model

type SysError uint32
type SysState uint32

// System state codes
const (
	SysErrSuccess SysError = iota
	SysErrSystemFail
	SysErrDeviceFail
)

func (e SysError) String() string {
	switch e {
	case SysErrSuccess:
		return "Success"
	case SysErrSystemFail:
		return "System fail"
	case SysErrDeviceFail:
		return "Device fail"
	default:
		return "Undefined"
	}
}

// System state codes
const (
	SysStateUndefined SysState = iota
	SysStateRunning
	SysStateStopped
	SysStateFailed
)

func (e SysState) String() string {
	switch e {
	case SysStateUndefined:
		return "Undefined"
	case SysStateRunning:
		return "Running"
	case SysStateStopped:
		return "Stopped"
	case SysStateFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}
