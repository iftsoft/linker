package callback

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type PrinterCallbackClient struct {
	log    *slog.Logger
	device device.PrinterCallbackServiceClient
}

func NewPrinterCallbackClient(log *slog.Logger, conn *grpc.ClientConn) *PrinterCallbackClient {
	return &PrinterCallbackClient{
		log:    log,
		device: device.NewPrinterCallbackServiceClient(conn),
	}
}

// PrinterProgress sent notification about printing progress
func (c *PrinterCallbackClient) PrinterProgress(ctx context.Context, value *model.PrinterProgress) error {
	c.log.Debug("CallbackClient.PrinterProgress - grpc",
		slog.String("device", value.Device), slog.String("document", value.DocName))

	input := &device.PrinterProgressRequest{
		Notify: convertDeviceNotifyToProto(&value.DeviceNotify),
		Data:   convertProgressNotifyToProto(&value.ProgressNotify),
	}
	_, err := c.device.PrinterProgress(ctx, input)
	if err != nil {
		return fmt.Errorf("callback PrinterProgress for %s.%s failed: %w", value.Device, value.DocName, err)
	}

	return nil
}

func convertProgressNotifyToProto(value *model.ProgressNotify) *device.ProgressNotify {
	if value == nil {
		return nil
	}
	data := &device.ProgressNotify{
		DocName:  value.DocName,
		PageDone: value.PageDone,
		PagesAll: value.PagesAll,
	}
	return data
}
