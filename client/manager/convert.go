package manager

import (
	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	system "github.com/iftsoft/linker/gen/go/linker/system/v1"
	"github.com/iftsoft/linker/model"
)

func convertSystemQuery(value *model.SystemQuery) *system.SystemQuery {
	if value == nil {
		return nil
	}
	data := &system.SystemQuery{
		Device: value.Device,
	}
	return data
}

func convertSystemConfig(value *model.SystemConfig) *system.SystemConfig {
	if value == nil {
		return nil
	}
	data := &system.SystemConfig{
		Device:    value.Device,
		LinkType:  value.LinkType,
		PortName:  value.PortName,
		VendorId:  value.VendorID,
		ProductId: value.ProductID,
	}
	return data
}

func convertSystemReply(value *system.SystemReply) *model.SystemReply {
	if value == nil {
		return nil
	}
	data := &model.SystemReply{
		Device:   value.Device,
		Command:  value.Command,
		Message:  value.Message,
		SysError: model.SysError(value.SysError),
		SysState: model.SysState(value.SysState),
	}
	return data
}

func convertSystemHealth(value *system.SystemHealth) *model.SystemHealth {
	if value == nil {
		return nil
	}
	data := &model.SystemHealth{
		Device:   value.Device,
		Moment:   value.Moment,
		SysError: model.SysError(value.SysError),
		SysState: model.SysState(value.SysState),
	}
	return data
}

func convertDeviceQuery(value *model.DeviceQuery) *device.DeviceQuery {
	if value == nil {
		return nil
	}
	data := &device.DeviceQuery{
		Device:  value.Device,
		Timeout: value.Timeout,
		Offline: value.Offline,
	}
	return data
}

func convertDeviceReply(value *device.DeviceReply) *model.DeviceReply {
	if value == nil {
		return nil
	}
	data := &model.DeviceReply{
		Device:  value.Device,
		Command: value.Command,
		Action:  model.DevAction(value.Action),
		State:   model.DevState(value.State),
		ErrCode: model.DevError(value.ErrCode),
		ErrText: value.ErrText,
	}
	return data
}
