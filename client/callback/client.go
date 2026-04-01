package callback

import (
	"context"
	"fmt"
	"log/slog"

	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	system "github.com/iftsoft/linker/gen/go/linker/system/v1"
	"github.com/iftsoft/linker/grpc/client"
	"github.com/iftsoft/linker/model"
)

type CallbackClient struct {
	log    *slog.Logger
	client *client.Client
	system system.SystemCallbackServiceClient
	device device.DeviceCallbackServiceClient
}

func NewCallbackClient(ctx context.Context, log *slog.Logger, address string) (*CallbackClient, error) {
	base, err := client.NewClient(ctx, address)
	if err != nil {
		return nil, fmt.Errorf("can't create grpc client: %w", err)
	}
	return &CallbackClient{
		log:    log,
		client: base,
		system: system.NewSystemCallbackServiceClient(base.Connection()),
		device: device.NewDeviceCallbackServiceClient(base.Connection()),
	}, nil
}
