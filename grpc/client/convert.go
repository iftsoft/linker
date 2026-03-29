package client

import (
	srv "github.com/iftsoft/linker/gen/go/linker/device/v1"
	model "github.com/iftsoft/linker/model"
)

func convertDeviceReply(value *model.DeviceReply) *srv.DeviceReply {
	if value == nil {
		return nil
	}
	data := &srv.DeviceReply{
		Device:  value.Device,
		Command: value.Command,
		Action:  uint32(value.Action),
		State:   uint32(value.State),
		ErrCode: uint32(value.ErrCode),
		ErrText: value.ErrText,
	}
	return data
}

func convertDeviceState(value *model.DeviceState) *srv.DeviceState {
	if value == nil {
		return nil
	}
	data := &srv.DeviceState{
		Device:   value.Device,
		Action:   uint32(value.Action),
		NewState: uint32(value.NewState),
		OldState: uint32(value.OldState),
	}
	return data
}

func convertDevicePrompt(value *model.DevicePrompt) *srv.DevicePrompt {
	if value == nil {
		return nil
	}
	data := &srv.DevicePrompt{
		Device: value.Device,
		Action: uint32(value.Action),
		Prompt: uint32(value.Prompt),
	}
	return data
}

func convertDeviceInform(value *model.DeviceInform) *srv.DeviceInform {
	if value == nil {
		return nil
	}
	data := &srv.DeviceInform{
		Device: value.Device,
		Action: uint32(value.Action),
		Inform: value.Inform,
	}
	return data
}
