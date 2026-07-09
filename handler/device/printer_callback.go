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

	reply := convertPrinterProgressToModel(req.GetData())
	h.log.Debug("gRPC.PrinterProgress", slog.Any("reply", reply))

	err := h.api.PrinterProgress(ctx, &reply)
	if err != nil {
		h.log.Error("gRPC.PrinterProgress failed", slog.Any("error", err))
	}

	resp := &srv.PrinterProgressResponse{}

	return resp, err
}

func convertPrinterProgressToModel(value *srv.PrinterProgress) model.PrinterProgress {
	if value == nil {
		return model.PrinterProgress{}
	}
	data := model.PrinterProgress{
		Device:   value.GetDevice(),
		DocName:  value.GetDocName(),
		PageDone: value.GetPageDone(),
		PagesAll: value.GetPagesAll(),
	}
	return data
}
