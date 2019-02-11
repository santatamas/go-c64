package internals

import (
	"bytes"
	"encoding/json"
)

// Command represents a debugger command sent from the WebUI
type Command int

const (
	Start Command = iota + 1
	Stop
	ExecuteNext
	GetCPUState
	GetEmulatorState
	GetMemoryContent
)

func (s Command) String() string {
	return toString[s]
}

var toString = map[Command]string{
	Start:            "Start",
	Stop:             "Stop",
	ExecuteNext:      "ExecuteNext",
	GetCPUState:      "GetCPUState",
	GetEmulatorState: "GetEmulatorState",
	GetMemoryContent: "GetMemoryContent",
}

var toID = map[string]Command{
	"Start":            Start,
	"Stop":             Stop,
	"ExecuteNext":      ExecuteNext,
	"GetCPUState":      GetCPUState,
	"GetEmulatorState": GetEmulatorState,
	"GetMemoryContent": GetMemoryContent,
}

// MarshalJSON marshals the enum as a quoted json string
func (s Command) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *Command) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = toID[j]
	return nil
}
