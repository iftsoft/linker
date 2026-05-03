package model

import "context"

const (
	CmdCardPosition    = "CardPosition"
	CmdCardDescription = "CardDescription"
	CmdChipResponse    = "ChipResponse"
	CmdEnterCard       = "EnterCard"
	CmdEjectCard       = "EjectCard"
	CmdCaptureCard     = "CaptureCard"
	CmdReadCard        = "ReadCard"
	CmdChipGetATR      = "ChipGetATR"
	CmdChipPowerOff    = "ChipPowerOff"
	CmdChipCommand     = "ChipCommand"
)

type ReaderCardPos struct {
	Position int16 `json:"position"`
}

type ReaderCardInfo struct {
	Track1  string `json:"track1"`
	Track2  string `json:"track2"`
	Track3  string `json:"track3"`
	RawData string `json:"raw_data"`
	CardPan string `json:"card_pan"`
	ExpDate string `json:"exp_date"`
	Holder  string `json:"holder"`
}

//type ReaderChipQuery struct {
//	Protocol int16  `json:"protocol"`
//	Query    []byte `json:"query"`
//}

//type ReaderChipReply struct {
//	DeviceReply
//	Protocol int16  `json:"protocol"`
//	Reply    []byte `json:"reply"`
//}

type ReaderCallback interface {
	CardPosition(ctx context.Context, value *ReaderCardPos) error
	CardDescription(ctx context.Context, value *ReaderCardInfo) error
	//ChipResponse(ctx context.Context, reply *ReaderChipReply) error
}

type ReaderManager interface {
	EnterCard(ctx context.Context, query *DeviceQuery) error
	EjectCard(ctx context.Context, query *DeviceQuery) error
	CaptureCard(ctx context.Context, query *DeviceQuery) error
	ReadCard(ctx context.Context, query *DeviceQuery) error
	//ChipGetATR(ctx context.Context, query *DeviceQuery) error
	//ChipPowerOff(ctx context.Context, query *DeviceQuery) error
	//ChipCommand(ctx context.Context, query *ReaderChipQuery) error
}
