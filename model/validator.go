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
	CmdNoteAccept     = "NoteAccept"
	CmdNoteReturn     = "NoteReturn"
	CmdStopValidate   = "StopValidate"
	CmdCheckValidator = "CheckValidator"
	CmdClearValidator = "ClearValidator"
)

type ValidNoteList []*ValidatorNote

type ValidatorNote struct {
	Device   string   `json:"device"`
	Currency Currency `json:"currency"`
	Count    Counter  `json:"count"`
	Nominal  Amount   `json:"nominal"`
	Amount   Amount   `json:"amount"`
}

type BatchState int16

const (
	StateEmpty BatchState = iota
	StateActive
	StateCorrect
	StateMismatch
)

func (e BatchState) String() string {
	switch e {
	case StateEmpty:
		return "Empty"
	case StateActive:
		return "Active"
	case StateCorrect:
		return "Correct"
	case StateMismatch:
		return "Mismatch"
	default:
		return "Unknown"
	}
}

type ValidatorBatch struct {
	Notes   ValidNoteList `json:"notes"`
	BatchId int64         `json:"batch_id"`
	State   BatchState    `json:"state"`
	Detail  string        `json:"detail"`
}

func (dev *ValidatorBatch) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Batch Id=%d, State=%s, Detail=%s, %s",
		dev.BatchId, dev.State.String(), dev.Detail, dev.Notes.String())
	return str
}

func (vn *ValidatorNote) String() string {
	if vn == nil {
		return ""
	}
	str := fmt.Sprintf("%s Note %7.2f * %3d = %9.2f of %3d (%s) - %s",
		vn.Device, vn.Nominal, vn.Count, vn.Amount, vn.Currency, vn.Currency.IsoCode(), vn.Currency.String())
	return str
}

func (vl ValidNoteList) String() string {
	str := "Validator Note List:"
	for i, note := range vl {
		if note != nil {
			str += fmt.Sprintf("\n    Line:%2d - %s", i, note.String())
		}
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
	Currency Currency `json:"currency"`
	Nominal  Amount   `json:"nominal"`
	Count    Counter  `json:"count"`
	Amount   Amount   `json:"amount"`
}

func (dev *ValidatorAccept) String() string {
	if dev == nil {
		return ""
	}
	str := fmt.Sprintf("Nominal: %7.2f, Count: %d, Amount: %7.2f, Currency: %d (%s) %s",
		dev.Nominal, dev.Count, dev.Amount, dev.Currency, dev.Currency.IsoCode(), dev.Currency.String())
	return str
}

type ValidatorQuery struct {
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

//type ValidatorReply struct {
//	Currency DevCurrency
//}

type ValidatorCallback interface {
	NoteAccepted(ctx context.Context, value *ValidatorAccept) error
	CashIsStored(ctx context.Context, value *ValidatorAccept) error
	CashReturned(ctx context.Context, value *ValidatorAccept) error
	ValidatorStore(ctx context.Context, reply *ValidatorStore) error
}

type ValidatorManager interface {
	InitValidator(ctx context.Context, query *ValidatorQuery) error
	DoValidate(ctx context.Context, query *ValidatorQuery) error
	NoteAccept(ctx context.Context, query *ValidatorQuery) error
	NoteReturn(ctx context.Context, query *ValidatorQuery) error
	StopValidate(ctx context.Context, query *ValidatorQuery) error
	CheckValidator(ctx context.Context, query *ValidatorQuery) error
	ClearValidator(ctx context.Context, query *ValidatorQuery) error
}

type ValidatorBooker interface {
	InitNoteList(ctx context.Context, list ValidNoteList) error
	ReadNoteList(ctx context.Context, data *ValidatorBatch) error
	DepositNote(ctx context.Context, extraId int64, value *ValidatorAccept) error
	CloseBatch(ctx context.Context, data *ValidatorBatch) error
}
