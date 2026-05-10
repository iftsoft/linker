package model

import (
	"context"
	"fmt"
)

const (
	CmdNoteAccepted   = "NoteAccepted"
	CmdCashIsStored   = "CashIsStored"
	CmdCashReturned   = "CashReturned"
	CmdValidatorStore = "ValidatorStore"
	CmdInitValidator  = "InitValidator"
	CmdDoValidate     = "DoValidate"
	CmdAcceptNote     = "AcceptNote"
	CmdReturnNote     = "ReturnNote"
	CmdStopValidate   = "StopValidate"
	CmdCheckValidator = "CheckValidator"
	CmdClearValidator = "ClearValidator"
)

type ValidatorCallback interface {
	// NoteAccepted sends notification about new note in escrow
	NoteAccepted(ctx context.Context, value *ValidatorAccept) error
	// CashIsStored sends notification that note is stored to cassette
	CashIsStored(ctx context.Context, value *ValidatorAccept) error
	// CashReturned sends notification that note is returned to user
	CashReturned(ctx context.Context, value *ValidatorAccept) error
	// ValidatorStore sends notification about current cassette state
	ValidatorStore(ctx context.Context, reply *ValidatorBatch) error
}

type ValidatorManager interface {
	// InitValidator does primary initialization of the validator
	InitValidator(ctx context.Context, query *ValidatorQuery) (*DeviceReply, error)
	// DoValidate starts accepting cash from user
	DoValidate(ctx context.Context, query *ValidatorQuery) (*DeviceReply, error)
	// AcceptNote puts the validated note to the cassette
	AcceptNote(ctx context.Context, query *ValidatorQuery) (*DeviceReply, error)
	// ReturnNote returns the validated note to the user
	ReturnNote(ctx context.Context, query *ValidatorQuery) (*DeviceReply, error)
	// StopValidate disables accepting new notes by validator
	StopValidate(ctx context.Context, query *ValidatorQuery) (*DeviceReply, error)
	// CheckValidator returns current cassette state
	CheckValidator(ctx context.Context, query *ValidatorQuery) (*ValidatorStore, error)
	// ClearValidator clears all cassette data (settlement or reconciliation)
	ClearValidator(ctx context.Context, query *ValidatorQuery) (*ValidatorStore, error)
}

type ValidatorNote struct {
	Currency Currency `json:"currency"`
	Nominal  Amount   `json:"nominal"`
	Count    Counter  `json:"count"`
	Amount   Amount   `json:"amount"`
}

func (vn *ValidatorNote) String() string {
	if vn == nil {
		return ""
	}
	str := fmt.Sprintf("%9s * %3d = %12s %s - %s",
		AmountText(vn.Nominal, vn.Currency), vn.Count, AmountText(vn.Amount, vn.Currency),
		vn.Currency.IsoCode(), vn.Currency.PlainText())
	return str
}

type ValidatorBatch struct {
	Device  string          `json:"device"`
	BatchId int64           `json:"batch_id"`
	State   BatchState      `json:"state"`
	Details string          `json:"details"`
	Notes   []ValidatorNote `json:"notes"`
}

func (dev *ValidatorBatch) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Device=%s, Batch Id=%d, State=%s, Details=%s",
		dev.Device, dev.BatchId, dev.State.String(), dev.Details)
	for i, note := range dev.Notes {
		str += fmt.Sprintf("\n    Line:%2d - Note: %s", i, note.String())
	}
	return str
}

type ValidatorStore struct {
	DeviceReply
	ValidatorBatch
}

func (dev *ValidatorStore) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("%s, %s",
		dev.DeviceReply.String(), dev.ValidatorBatch.String())
	return str
}

type ValidatorAccept struct {
	Device string        `json:"device"`
	Note   ValidatorNote `json:"note"`
}

func (dev *ValidatorAccept) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Device: %7, Note: %s", dev.Device, dev.Note.String())
	return str
}

type ValidatorQuery struct {
	Device    string   `json:"device"`
	Currency  Currency `json:"currency"`
	Operation int64    `json:"operation"`
}

func (dev *ValidatorQuery) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Currency = %s, Operation = %d",
		dev.Currency, dev.Operation)
	return str
}

//type ValidatorBooker interface {
//	InitNoteList(ctx context.Context, list ValidatorNoteList) error
//	ReadNoteList(ctx context.Context, data *ValidatorBatch) error
//	DepositNote(ctx context.Context, extraId int64, value *ValidatorAccept) error
//	CloseBatch(ctx context.Context, data *ValidatorBatch) error
//}
