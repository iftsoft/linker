package manager

import (
	"context"
	"fmt"
	"log/slog"

	"google.golang.org/grpc"

	device "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type PrinterManagerClient struct {
	log    *slog.Logger
	device device.PrinterManagerServiceClient
}

func NewPrinterManagerClient(log *slog.Logger, conn *grpc.ClientConn) *PrinterManagerClient {
	return &PrinterManagerClient{
		log:    log,
		device: device.NewPrinterManagerServiceClient(conn),
	}
}

// InitPrinter does primary initialization of printer before printing
func (c *PrinterManagerClient) InitPrinter(ctx context.Context, query *model.PrinterSetup) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.InitPrinter - grpc", slog.String("device", query.Device))

	input := &device.InitPrinterRequest{
		Query: convertPrinterSetup(query),
	}
	resp, err := c.device.InitPrinter(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command InitPrinter for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReply(resp.GetReply())
	return reply, nil
}

// PrintPage trys to print given text on the printer
func (c *PrinterManagerClient) PrintPage(ctx context.Context, query *model.PrinterQuery) (*model.DeviceReply, error) {
	c.log.Debug("ManagerClient.PrintPage - grpc", slog.String("device", query.Device))

	input := &device.PrintPageRequest{
		Query: convertPrinterQuery(query),
	}
	resp, err := c.device.PrintPage(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("command PrintPage for %s failed: %w", query.Device, err)
	}

	reply := convertDeviceReply(resp.GetReply())
	return reply, nil
}

func convertPrinterSetup(value *model.PrinterSetup) *device.PrinterSetup {
	if value == nil {
		return nil
	}
	data := &device.PrinterSetup{
		Device:    value.Device,
		PaperPath: value.PaperPath,
		Landscape: value.Landscape,
		ShowImage: value.ShowImage,
	}
	return data
}

func convertPrinterQuery(value *model.PrinterQuery) *device.PrinterQuery {
	if value == nil {
		return nil
	}
	data := &device.PrinterQuery{
		Device: value.Device,
		Text:   value.Text,
	}
	return data
}
