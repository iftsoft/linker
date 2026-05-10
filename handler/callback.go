package handler

import (
	"log/slog"

	"github.com/iftsoft/linker/grpc/server"
	"github.com/iftsoft/linker/handler/device"
	"github.com/iftsoft/linker/handler/system"
	"github.com/iftsoft/linker/model"
)

type CallbackService interface {
	model.SystemCallback
	model.DeviceCallback
	model.PrinterCallback
	model.ReaderCallback
	model.ValidatorCallback
}

func NewCallbackServer(log *slog.Logger, address string, service CallbackService) *server.Server {
	grpcSrv := server.NewServer(log, address,
		system.NewSystemCallback(log, service),
		device.NewDeviceCallback(log, service),
		device.NewPrinterCallback(log, service),
		device.NewReaderCallback(log, service),
		device.NewValidatorCallback(log, service),
	)
	return grpcSrv
}
