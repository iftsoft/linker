package model

import "context"

const (
	CmdInitPrinter     = "InitPrinter"
	CmdPrintText       = "PrintText"
	CmdPrinterProgress = "PrinterProgress"
)

type PrinterQuery struct {
	Text string `json:"text"`
}

type PrinterSetup struct {
	PaperPath int32 `json:"paper_path"`
	Landscape bool  `json:"landscape"`
	ShowImage int32 `json:"show_image"`
}

type PrinterProgress struct {
	DocName  string `json:"doc_name"`
	PageDone int32  `json:"page_done"`
	PagesAll int32  `json:"pages_all"`
}

type PrinterCallback interface {
	PrinterProgress(ctx context.Context, reply *PrinterProgress) error
}

type PrinterManager interface {
	InitPrinter(ctx context.Context, query *PrinterSetup) error
	PrintText(ctx context.Context, query *PrinterQuery) error
}
