package model

import "strings"

type SysError uint32
type SysState uint32
type DevTypeMask uint64
type DevScopeMask uint64

// System state codes
const (
	SysErrSuccess SysError = iota
	SysErrSystemFail
	SysErrDeviceFail
)

func (e SysError) String() string {
	switch e {
	case SysErrSuccess:
		return "Success"
	case SysErrSystemFail:
		return "System fail"
	case SysErrDeviceFail:
		return "Device fail"
	default:
		return "Undefined"
	}
}

// System state codes
const (
	SysStateUndefined SysState = iota
	SysStateRunning
	SysStateStopped
	SysStateFailed
)

func (e SysState) String() string {
	switch e {
	case SysStateUndefined:
		return "Undefined"
	case SysStateRunning:
		return "Running"
	case SysStateStopped:
		return "Stopped"
	case SysStateFailed:
		return "Failed"
	default:
		return "Unknown"
	}
}

// Scope Flags
const (
	ScopeFlagSystem DevScopeMask = 1 << iota
	ScopeFlagDevice
	ScopeFlagPrinter
	ScopeFlagReader
	ScopeFlagPinPad
	ScopeFlagValidator
	ScopeFlagDispenser
	ScopeFlagVending
	ScopeFlagCustom
	ScopeFlagUnknown = 0
)

// Device types
const (
	DevTypePrinter DevTypeMask = 1 << iota
	DevTypeCardReader
	DevTypeBarScanner
	DevTypeCashValidator
	DevTypeCoinValidator
	DevTypeCashDispenser
	DevTypeCoinDispenser
	DevTypeVending
	DevTypePINEntry
	DevTypeCustom
	DevTypeUndefined = 0
	DevTypeCommon    = 0x0FF
)

func (e DevTypeMask) ToString() string {
	if e == DevTypeUndefined {
		return "Undefined"
	}
	list := make([]string, 0)
	if (e & DevTypePrinter) == DevTypePrinter {
		list = append(list, "Printer")
	}
	if (e & DevTypeCardReader) == DevTypeCardReader {
		list = append(list, "CardReader")
	}
	if (e & DevTypeBarScanner) == DevTypeBarScanner {
		list = append(list, "BarScanner")
	}
	if (e & DevTypeCashValidator) == DevTypeCashValidator {
		list = append(list, "CashValidator")
	}
	if (e & DevTypeCoinValidator) == DevTypeCoinValidator {
		list = append(list, "CoinValidator")
	}
	if (e & DevTypeCashDispenser) == DevTypeCashDispenser {
		list = append(list, "CashDispenser")
	}
	if (e & DevTypeCoinDispenser) == DevTypeCoinDispenser {
		list = append(list, "CoinDispenser")
	}
	if (e & DevTypeVending) == DevTypeVending {
		list = append(list, "Vending")
	}
	if (e & DevTypePINEntry) == DevTypePINEntry {
		list = append(list, "PINEntry")
	}
	if (e & DevTypeCustom) == DevTypeCustom {
		list = append(list, "Custom")
	}
	return strings.Join(list, ",")
}
