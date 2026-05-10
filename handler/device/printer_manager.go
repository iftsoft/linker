package device

import (
	"context"
	"errors"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	srv "github.com/iftsoft/linker/gen/go/linker/device/v1"
	model "github.com/iftsoft/linker/model"
)

type PrinterManager struct {
	log *slog.Logger
	api model.PrinterManager
	srv.PrinterManagerServiceServer
}

func NewPrinterManager(log *slog.Logger, api model.PrinterManager) *PrinterManager {
	return &PrinterManager{
		log: log,
		api: api,
	}
}

func (h *PrinterManager) Register(s grpc.ServiceRegistrar) {
	srv.RegisterPrinterManagerServiceServer(s, h)
}

// InitPrinter does primary initialization of printer before printing
func (h *PrinterManager) InitPrinter(ctx context.Context, req *srv.InitPrinterRequest) (*srv.InitPrinterResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("InitPrinterRequest is nil"))
	}

	query := PrinterSetupToModel(req.GetQuery())
	h.log.Debug("gRPC.InitPrinter", slog.Any("query", query))

	reply, err := h.api.InitPrinter(ctx, query)
	if err != nil {
		h.log.Error("gRPC.InitPrinter failed", slog.Any("error", err))
	}

	resp := &srv.InitPrinterResponse{
		Reply: DeviceReplyToProto(reply),
	}

	return resp, err
}

// PrintPage trys to print given text on the printer
func (h *PrinterManager) PrintPage(ctx context.Context, req *srv.PrintPageRequest) (*srv.PrintPageResponse, error) {
	if req == nil {
		return nil, MakeErrorWithDetails(codes.InvalidArgument, StrMissingRequest,
			errors.New("PrintPageRequest is nil"))
	}

	query := PrinterQueryToModel(req.GetQuery())
	h.log.Debug("gRPC.PrintPage", slog.Any("query", query))

	reply, err := h.api.PrintPage(ctx, query)
	if err != nil {
		h.log.Error("gRPC.PrintPage failed", slog.Any("error", err))
	}

	resp := &srv.PrintPageResponse{
		Reply: DeviceReplyToProto(reply),
	}

	return resp, err
}

func PrinterSetupToModel(data *srv.PrinterSetup) *model.PrinterSetup {
	query := &model.PrinterSetup{
		Device:    data.GetDevice(),
		PaperPath: data.GetPaperPath(),
		Landscape: data.GetLandscape(),
		ShowImage: data.GetShowImage(),
	}
	return query
}

func PrinterQueryToModel(data *srv.PrinterQuery) *model.PrinterQuery {
	query := &model.PrinterQuery{
		Device: data.GetDevice(),
		Text:   data.GetText(),
	}
	return query
}
