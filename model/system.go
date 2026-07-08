package model

import (
	"context"
	"fmt"
)

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
	// GreetingInfo sends notification about device application
	GreetingInfo(ctx context.Context, query *GreetingInfo) error
	// SystemReply sends notification about system reply
	SystemReply(ctx context.Context, reply *SystemReply) error
	// SystemHealth sends notification about execute error
	SystemHealth(ctx context.Context, value *SystemHealth) error
}

type GreetingInfo struct {
	Device      string       `json:"device"`      // Name of device
	GrpcPort    uint32       `json:"grpc_port"`   // gRPC port for device management
	DevType     DevTypeMask  `json:"dev_type"`    // Implemented device types
	Supported   DevScopeMask `json:"supported"`   // Manager interfaces that driver supported
	Required    DevScopeMask `json:"required"`    // Callback interfaces that driver required
	Description string       `json:"description"` // Description of device purpose
}

type SystemQuery struct {
	Device string `json:"device"`
}

func (sys *SystemQuery) String() string {
	if sys == nil {
		return ""
	}
	str := fmt.Sprintf("Device = %s", sys.Device)
	return str
}

type SystemConfig struct {
	Device    string `json:"device"`
	LinkType  uint32 `json:"link_type"`  // 0-none, 1-COM, 2-USB
	PortName  string `json:"port_name"`  // Serial port name
	VendorID  uint32 `json:"vendor_id"`  // Device Vendor ID
	ProductID uint32 `json:"product_id"` // Device Product ID
}

func (sys *SystemConfig) String() string {
	if sys == nil {
		return ""
	}
	str := fmt.Sprintf("Device = %s, LinkType = %d, PortName = %s, VendorID = %d, ProductID = %d",
		sys.Device, sys.LinkType, sys.PortName, sys.VendorID, sys.ProductID)
	return str
}

type SystemReply struct {
	Device   string   `json:"device"`
	Command  string   `json:"command"`
	Message  string   `json:"message"`
	SysError SysError `json:"sys_error"`
	SysState SysState `json:"sys_state"`
}

func (sys *SystemReply) String() string {
	if sys == nil {
		return ""
	}
	str := fmt.Sprintf("Device = %s, Command = %s, Message = %s, SysError = %s, SysState = %s",
		sys.Device, sys.Command, sys.Message, sys.SysError.String(), sys.SysState.String())
	return str
}

type SystemMetrics struct {
	Uptime   uint64             `json:"uptime"`
	DevError DevError           `json:"dev_error"`
	DevState DevState           `json:"dev_state"`
	Counts   map[string]uint32  `json:"counts"`
	Totals   map[string]float32 `json:"totals"`
	Topics   map[string]string  `json:"topics"`
}

func (sys *SystemMetrics) String() string {
	if sys == nil {
		return ""
	}
	str := fmt.Sprintf("Uptime = %d, DevError = %s, DevState = %s",
		sys.Uptime, sys.DevError.String(), sys.DevState.String())
	return str
}

type SystemHealth struct {
	Device   string        `json:"device"`
	Moment   int64         `json:"moment"`
	SysError SysError      `json:"error"`
	SysState SysState      `json:"state"`
	Metrics  SystemMetrics `json:"metrics"`
}

func (sys *SystemHealth) String() string {
	if sys == nil {
		return ""
	}
	str := fmt.Sprintf("Device = %s, Moment = %d, SysError = %s, SysState = %s, Metrics = (%s)",
		sys.Device, sys.Moment, sys.SysError.String(), sys.SysState.String(), sys.Metrics.String())
	return str
}

func NewSystemHealth(dev string) *SystemHealth {
	sh := &SystemHealth{
		Device:   dev,
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
