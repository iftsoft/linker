package model

type BatchState uint32

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
