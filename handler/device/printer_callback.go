package device

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	srv "github.com/iftsoft/linker/gen/go/linker/device/v1"
	"github.com/iftsoft/linker/model"
)

type PrinterCallback struct {
	log *slog.Logger
	api model.PrinterCallback
	srv.PrinterCallbackServiceServer
}

func NewPrinterCallback(log *slog.Logger, api model.PrinterCallback) *PrinterCallback {
	return &PrinterCallback{
		log: log,
		api: api,
	}
}

func (h *PrinterCallback) Register(s grpc.ServiceRegistrar) {
	srv.RegisterPrinterCallbackServiceServer(s, h)
}

// PrinterProgress sent notification about printing progress
func (h *PrinterCallback) PrinterProgress(ctx context.Context, req *srv.PrinterProgressRequest) (*srv.PrinterProgressResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("PrinterProgressRequest is nil"))
	}

	value := model.PrinterProgress{
		DeviceNotify:   convertDeviceNotifyToModel(req.GetNotify()),
		ProgressNotify: convertProgressNotifyToModel(req.GetData()),
	}
	h.log.Debug("gRPC.PrinterProgress", slog.Any("value", value))

	err := h.api.PrinterProgress(ctx, &value)
	if err != nil {
		h.log.Error("gRPC.PrinterProgress failed", slog.Any("error", err))
	}

	resp := &srv.PrinterProgressResponse{}

	return resp, err
}

func convertProgressNotifyToModel(value *srv.ProgressNotify) model.ProgressNotify {
	if value == nil {
		return model.ProgressNotify{}
	}
	data := model.ProgressNotify{
		DocName:  value.GetDocName(),
		PageDone: value.GetPageDone(),
		PagesAll: value.GetPagesAll(),
	}
	return data
}
