package callback

import (
	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	system "github.com/iftsoft/linker/gen/go/linker/system/v1"
	"github.com/iftsoft/linker/model"
)

func convertDeviceReply(value *model.DeviceReply) *device.DeviceReply {
	if value == nil {
		return nil
	}
	data := &device.DeviceReply{
		Device:  value.Device,
		Command: value.Command,
		Action:  uint32(value.Action),
		State:   uint32(value.State),
		ErrCode: uint32(value.ErrCode),
		ErrText: value.ErrText,
	}
	return data
}

func convertDeviceState(value *model.DeviceState) *device.DeviceState {
	if value == nil {
		return nil
	}
	data := &device.DeviceState{
		Device:   value.Device,
		Action:   uint32(value.Action),
		NewState: uint32(value.NewState),
		OldState: uint32(value.OldState),
	}
	return data
}

func convertDevicePrompt(value *model.DevicePrompt) *device.DevicePrompt {
	if value == nil {
		return nil
	}
	data := &device.DevicePrompt{
		Device: value.Device,
		Action: uint32(value.Action),
		Prompt: uint32(value.Prompt),
	}
	return data
}

func convertDeviceInform(value *model.DeviceInform) *device.DeviceInform {
	if value == nil {
		return nil
	}
	data := &device.DeviceInform{
		Device: value.Device,
		Action: uint32(value.Action),
		Inform: value.Inform,
	}
	return data
}

func convertSystemReply(value *model.SystemReply) *system.SystemReply {
	if value == nil {
		return nil
	}
	data := &system.SystemReply{
		Device:   value.Device,
		Command:  value.Command,
		Message:  value.Message,
		SysError: uint32(value.SysError),
		SysState: uint32(value.SysState),
	}
	return data
}

func convertSystemHealth(value *model.SystemHealth) *system.SystemHealth {
	if value == nil {
		return nil
	}
	data := &system.SystemHealth{
		Device:   value.Device,
		Moment:   value.Moment,
		SysError: uint32(value.SysError),
		SysState: uint32(value.SysState),
	}
	return data
}
