package model

type SysError uint32
type SysState uint32

const (
	CmdSystemReply     = "SystemReply"
	CmdSystemHealth    = "SystemHealth"
	CmdSystemTerminate = "Terminate"
	CmdSystemInform    = "SysInform"
	CmdSystemStart     = "SysStart"
	CmdSystemStop      = "SysStop"
	CmdSystemRestart   = "SysRestart"
)

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

type SystemQuery struct {
	Device string `json:"device"`
}

type SystemConfig struct {
	Device    string `json:"device"`
	LinkType  uint32 `json:"link_type"`  // 0-none, 1-COM, 2-USB
	PortName  string `json:"port_name"`  // Serial port name
	VendorID  uint32 `json:"vendor_id"`  // Device Vendor ID
	ProductID uint32 `json:"product_id"` // Device Product ID
}

type SystemReply struct {
	Device   string   `json:"device"`
	Command  string   `json:"command"`
	Message  string   `json:"message"`
	SysError SysError `json:"sys_error"`
	SysState SysState `json:"sys_state"`
}

type SystemMetrics struct {
	Uptime   uint64             `json:"uptime"`
	DevError DevError           `json:"dev_error"`
	DevState DevState           `json:"dev_state"`
	Counts   map[string]uint32  `json:"counts"`
	Totals   map[string]float32 `json:"totals"`
	Topics   map[string]string  `json:"topics"`
}

type SystemHealth struct {
	Device   string        `json:"device"`
	Moment   int64         `json:"moment"`
	SysError SysError      `json:"error"`
	SysState SysState      `json:"state"`
	Metrics  SystemMetrics `json:"metrics"`
}

func NewSystemHealth() *SystemHealth {
	sh := &SystemHealth{
		Moment:   0,
		SysError: 0,
		SysState: 0,
		Metrics: SystemMetrics{
			Uptime:   0,
			DevError: 0,
			DevState: 0,
			Counts:   make(map[string]uint32),
			Totals:   make(map[string]float32),
			Topics:   make(map[string]string),
		},
	}
	return sh
}

type SystemCallback interface {
	SystemReply(name string, reply *SystemReply) error
	SystemHealth(name string, reply *SystemHealth) error
}

type SystemManager interface {
	Terminate(name string, query *SystemQuery) error
	SysInform(name string, query *SystemQuery) error
	SysStart(name string, query *SystemConfig) error
	SysStop(name string, query *SystemQuery) error
	SysRestart(name string, query *SystemConfig) error
}
