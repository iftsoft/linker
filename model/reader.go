package model

import (
	"context"
	"fmt"
	"strings"
)

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
	ReadCard(ctx context.Context, query *DeviceQuery) (*ReadCardReply, error)
}

type CardPosition struct {
	Device   string `json:"device"`
	Position int32  `json:"position"`
}

type CardPAN string
type Track string

type CardDescription struct {
	Device  string  `json:"device"`
	CardPan CardPAN `json:"card_pan"`
	ExpDate string  `json:"exp_date"`
	Holder  string  `json:"holder"`
	Track1  Track   `json:"track1"`
	Track2  Track   `json:"track2"`
	Track3  Track   `json:"track3"`
}

type ReadCardReply struct {
	DeviceReply
	CardDescription
}

func (pan CardPAN) String() string {
	size := len(pan)
	if size < 10 {
		return string(pan)
	}
	beg := pan[0:4]
	mid := strings.Repeat("*", size-8)
	end := pan[size-4 : size]
	str := fmt.Sprintf("%s%s%s", beg, mid, end)
	return str
}

func (tr Track) String() string {
	size := len(tr)
	if size < 10 {
		return string(tr)
	}
	beg := tr[0:4]
	end := tr[size-4 : size]
	str := fmt.Sprintf("%s-%d-%s", beg, size-8, end)
	return str
}
