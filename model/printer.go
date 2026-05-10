package model

import "context"

const (
	CmdInitPrinter     = "InitPrinter"
	CmdPrintPage       = "PrintPage"
	CmdPrinterProgress = "PrinterProgress"
)

type PrinterCallback interface {
	// PrinterProgress sent notification about printing progress
	PrinterProgress(ctx context.Context, reply *PrinterProgress) error
}

type PrinterManager interface {
	// InitPrinter does primary initialization of printer before printing
	InitPrinter(ctx context.Context, query *PrinterSetup) (*DeviceReply, error)
	// PrintPage trys to print given text on the printer
	PrintPage(ctx context.Context, query *PrinterQuery) (*DeviceReply, error)
}

type PrinterQuery struct {
	Device string `json:"device"`
	Text   string `json:"text"`
}

type PrinterSetup struct {
	Device    string `json:"device"`
	PaperPath uint32 `json:"paper_path"`
	Landscape bool   `json:"landscape"`
	ShowImage uint32 `json:"show_image"`
}

type PrinterProgress struct {
	Device   string `json:"device"`
	DocName  string `json:"doc_name"`
	PageDone int32  `json:"page_done"`
	PagesAll int32  `json:"pages_all"`
}
