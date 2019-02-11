package internals

import (
	"bytes"
	//"encoding/json"
	//"log"
)

// Command represents a debugger command sent from the WebUI
type Command int

const (
	Start Command = iota + 1
	Stop
	Reset
	ExecuteNext
	GetCPUState
	GetEmulatorState
	GetMemoryContent
	SetBreakpoint
)

func (s Command) String() string {
	return toString[s]
}

var toString = map[Command]string{
	Start:            "Start",
	Stop:             "Stop",
	Reset:            "Reset",
	ExecuteNext:      "ExecuteNext",
	GetCPUState:      "GetCPUState",
	GetEmulatorState: "GetEmulatorState",
	GetMemoryContent: "GetMemoryContent",
	SetBreakpoint:    "SetBreakpoint",
}

var toID = map[string]Command{
	"Start":            Start,
	"Stop":             Stop,
	"Reset":            Reset,
	"ExecuteNext":      ExecuteNext,
	"GetCPUState":      GetCPUState,
	"GetEmulatorState": GetEmulatorState,
	"GetMemoryContent": GetMemoryContent,
	"SetBreakpoint":    SetBreakpoint,
}

// MarshalJSON marshals the enum as a quoted json string
func (s Command) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *Command) UnmarshalJSON(b string) error {
	*s = toID[b]
	return nil
}
