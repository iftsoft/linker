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
}

func NewCallbackServer(log *slog.Logger, address string, service CallbackService) *server.Server {
	grpcSrv := server.NewServer(log, address,
		system.NewCallback(log, service),
		device.NewCallback(log, service),
	)
	return grpcSrv
}
