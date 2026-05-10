package handler

import (
	"log/slog"

	"github.com/iftsoft/linker/grpc/server"
	"github.com/iftsoft/linker/handler/device"
	"github.com/iftsoft/linker/handler/system"
	"github.com/iftsoft/linker/model"
)

type ManagerService interface {
	model.SystemManager
	model.DeviceManager
	model.PrinterManager
	model.ReaderManager
	model.ValidatorManager
}

func NewManagerServer(log *slog.Logger, address string, service ManagerService) *server.Server {
	grpcSrv := server.NewServer(log, address,
		system.NewSystemManager(log, service),
		device.NewDeviceManager(log, service),
		device.NewPrinterManager(log, service),
		device.NewReaderManager(log, service),
		device.NewValidatorManager(log, service),
	)
	return grpcSrv
}
