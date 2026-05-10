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

type ReaderCallback interface {
	// CardPosition sends notification about new card position
	CardPosition(ctx context.Context, value *CardPosition) error
	// CardDescription sends notification about card information
	CardDescription(ctx context.Context, value *CardDescription) error
}

type ReaderManager interface {
	// EnterCard trys to accept card in card reader device
	EnterCard(ctx context.Context, query *DeviceQuery) (*DeviceReply, error)
	// EjectCard trys to reject card from card reader device
	EjectCard(ctx context.Context, query *DeviceQuery) (*DeviceReply, error)
	// CaptureCard trys to capture card inside card reader device
	CaptureCard(ctx context.Context, query *DeviceQuery) (*DeviceReply, error)
	// ReadCard trys to read card information from card
	ReadCard(ctx context.Context, query *DeviceQuery) (*CardDescription, error)
}

type CardPosition struct {
	Device   string `json:"device"`
	Position int32  `json:"position"`
}

type CardDescription struct {
	Device  string `json:"device"`
	CardPan string `json:"card_pan"`
	ExpDate string `json:"exp_date"`
	Holder  string `json:"holder"`
	Track1  string `json:"track1"`
	Track2  string `json:"track2"`
	Track3  string `json:"track3"`
}
