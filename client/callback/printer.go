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
func (c *PrinterCallbackClient) PrinterProgress(ctx context.Context, reply *model.PrinterProgress) error {
	c.log.Debug("CallbackClient.PrinterProgress - grpc",
		slog.String("device", reply.Device), slog.String("document", reply.DocName))

	input := &device.PrinterProgressRequest{
		Data: convertPrinterProgress(reply),
	}
	_, err := c.device.PrinterProgress(ctx, input)
	if err != nil {
		return fmt.Errorf("callback PrinterProgress for %s.%s failed: %w", reply.Device, reply.DocName, err)
	}

	return nil
}

func convertPrinterProgress(value *model.PrinterProgress) *device.PrinterProgress {
	if value == nil {
		return nil
	}
	data := &device.PrinterProgress{
		Device:   value.Device,
		DocName:  value.DocName,
		PageDone: uint32(value.PageDone),
		PagesAll: uint32(value.PagesAll),
	}
	return data
}
