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
}

func NewManagerServer(log *slog.Logger, address string, service ManagerService) *server.Server {
	grpcSrv := server.NewServer(log, address,
		system.NewManager(log, service),
		device.NewManager(log, service),
	)
	return grpcSrv
}
