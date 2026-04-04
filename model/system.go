package model

import "context"

const (
	CmdSystemReply     = "SystemReply"
	CmdSystemHealth    = "SystemHealth"
	CmdSystemTerminate = "Terminate"
	CmdSystemInform    = "SysInform"
	CmdSystemStart     = "SysStart"
	CmdSystemStop      = "SysStop"
	CmdSystemRestart   = "SysRestart"
)

// SystemManager is the client API for SystemManagerService.
type SystemManager interface {
	// Terminate gracefully terminates running device application
	Terminate(ctx context.Context, query *SystemQuery) (*SystemReply, error)
	// SysInform returns health of device application
	SysInform(ctx context.Context, query *SystemQuery) (*SystemHealth, error)
	// SysStart turns device driver to initial state
	SysStart(ctx context.Context, query *SystemConfig) (*SystemReply, error)
	// SysStop gracefully deactivates device driver
	SysStop(ctx context.Context, query *SystemQuery) (*SystemReply, error)
	// SysRestart reactivates device driver with new config
	SysRestart(ctx context.Context, query *SystemConfig) (*SystemReply, error)
}

// SystemCallback is the client API for SystemCallbackService.
type SystemCallback interface {
	// SystemReply sends notification about system reply
	SystemReply(ctx context.Context, reply *SystemReply) error
	// SystemHealth sends notification about execute error
	SystemHealth(ctx context.Context, value *SystemHealth) error
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
