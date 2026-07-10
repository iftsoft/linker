package model

import (
	"context"
)

const (
	CmdSystemReply     = "SystemReply"
	CmdSystemHealth    = "SystemHealth"
	CmdSystemDevice    = "SystemDevice"
	CmdSystemTerminate = "Terminate"
	CmdSystemStart     = "SysStart"
	CmdSystemStop      = "SysStop"
	CmdSystemRestart   = "SysRestart"
)

// SystemManager is the client API for SystemManagerService.
type SystemManager interface {
	// Terminate gracefully terminates running device application
	Terminate(ctx context.Context, query *SystemQuery) (*SystemReply, error)
	// SysHealth returns health of device application
	SysHealth(ctx context.Context, query *SystemQuery) (*SystemHealth, error)
	// SysStart turns device driver to initial state
	SysStart(ctx context.Context, query *ConfigUpdate) (*SystemDevice, error)
	// SysStop gracefully deactivates device driver
	SysStop(ctx context.Context, query *SystemQuery) (*SystemReply, error)
	// SysRestart reactivates device driver with new config
	SysRestart(ctx context.Context, query *ConfigUpdate) (*SystemDevice, error)
}

// SystemCallback is the client API for SystemCallbackService.
type SystemCallback interface {
	// GreetingInfo sends notification about device application
	GreetingInfo(ctx context.Context, query *GreetingInfo) error
	// SystemReply sends notification about system reply
	SystemReply(ctx context.Context, reply *SystemReply) error
	// SystemDevice sends notification about system device
	SystemDevice(ctx context.Context, reply *SystemDevice) error
	// SystemHealth sends notification about execute error
	SystemHealth(ctx context.Context, value *SystemHealth) error
}

type GreetingInfo struct {
	AppName  string `json:"app_name"`  // Name of application
	DevName  string `json:"dev_name"`  // Name of device
	GrpcPort int64  `json:"grpc_port"` // gRPC port for device management
}

type SystemQuery struct {
	Device string `json:"device"`
}

type ConfigUpdate struct {
	Device    string `json:"device"`
	LinkType  uint32 `json:"link_type"`  // 0-none, 1-COM, 2-USB
	PortName  string `json:"port_name"`  // Serial port name
	VendorID  uint32 `json:"vendor_id"`  // Device Vendor ID
	ProductID uint32 `json:"product_id"` // Device Product ID
	SerialNo  string `json:"serial_no"`  // Serial Number
}

type SystemReply struct {
	Device   string   `json:"device"`
	Command  string   `json:"command"`
	Message  string   `json:"message"`
	SysError SysError `json:"sys_error"`
	SysState SysState `json:"sys_state"`
}

type DeviceSetup struct {
	DevType     DevTypeMask  `json:"dev_type"`    // Implemented device types
	Supported   DevScopeMask `json:"supported"`   // Manager interfaces that driver supported
	Required    DevScopeMask `json:"required"`    // Callback interfaces that driver required
	Description string       `json:"description"` // Description of device purpose
}

type SystemDevice struct {
	SystemReply
	DeviceSetup
}

type DeviceMetrics struct {
	Moment   int64          `json:"moment"`
	Uptime   uint64         `json:"uptime"`
	DevError DevError       `json:"dev_error"`
	DevState DevState       `json:"dev_state"`
	Content  map[string]any `json:"content"`
}

type SystemHealth struct {
	SystemReply
	DeviceMetrics
}
