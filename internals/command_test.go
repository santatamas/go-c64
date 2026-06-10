package internals

import (
	"testing"
)

func TestCommand_String_Start(t *testing.T) {
	cmd := Start
	if cmd.String() != "Start" {
		t.Errorf("Expected 'Start', got '%s'", cmd.String())
	}
}

func TestCommand_String_Stop(t *testing.T) {
	cmd := Stop
	if cmd.String() != "Stop" {
		t.Errorf("Expected 'Stop', got '%s'", cmd.String())
	}
}

func TestCommand_String_Reset(t *testing.T) {
	cmd := Reset
	if cmd.String() != "Reset" {
		t.Errorf("Expected 'Reset', got '%s'", cmd.String())
	}
}

func TestCommand_String_ExecuteNext(t *testing.T) {
	cmd := ExecuteNext
	if cmd.String() != "ExecuteNext" {
		t.Errorf("Expected 'ExecuteNext', got '%s'", cmd.String())
	}
}

func TestCommand_String_GetCPUState(t *testing.T) {
	cmd := GetCPUState
	if cmd.String() != "GetCPUState" {
		t.Errorf("Expected 'GetCPUState', got '%s'", cmd.String())
	}
}

func TestCommand_String_GetEmulatorState(t *testing.T) {
	cmd := GetEmulatorState
	if cmd.String() != "GetEmulatorState" {
		t.Errorf("Expected 'GetEmulatorState', got '%s'", cmd.String())
	}
}

func TestCommand_String_GetCIAState(t *testing.T) {
	cmd := GetCIAState
	if cmd.String() != "GetCIAState" {
		t.Errorf("Expected 'GetCIAState', got '%s'", cmd.String())
	}
}

func TestCommand_String_GetMemoryContent(t *testing.T) {
	cmd := GetMemoryContent
	if cmd.String() != "GetMemoryContent" {
		t.Errorf("Expected 'GetMemoryContent', got '%s'", cmd.String())
	}
}

func TestCommand_String_SetBreakpoint(t *testing.T) {
	cmd := SetBreakpoint
	if cmd.String() != "SetBreakpoint" {
		t.Errorf("Expected 'SetBreakpoint', got '%s'", cmd.String())
	}
}

func TestCommand_MarshalJSON_Start(t *testing.T) {
	cmd := Start
	data, err := cmd.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON returned error: %v", err)
	}
	expected := `"Start"`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestCommand_MarshalJSON_Stop(t *testing.T) {
	cmd := Stop
	data, err := cmd.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON returned error: %v", err)
	}
	expected := `"Stop"`
	if string(data) != expected {
		t.Errorf("Expected %s, got %s", expected, string(data))
	}
}

func TestCommand_MarshalJSON_AllCommands(t *testing.T) {
	tests := []struct {
		cmd      Command
		expected string
	}{
		{Start, `"Start"`},
		{Stop, `"Stop"`},
		{Reset, `"Reset"`},
		{ExecuteNext, `"ExecuteNext"`},
		{GetCPUState, `"GetCPUState"`},
		{GetEmulatorState, `"GetEmulatorState"`},
		{GetCIAState, `"GetCIAState"`},
		{GetMemoryContent, `"GetMemoryContent"`},
		{SetBreakpoint, `"SetBreakpoint"`},
	}

	for _, tt := range tests {
		data, err := tt.cmd.MarshalJSON()
		if err != nil {
			t.Fatalf("MarshalJSON for %v returned error: %v", tt.cmd, err)
		}
		if string(data) != tt.expected {
			t.Errorf("For command %v: expected %s, got %s", tt.cmd, tt.expected, string(data))
		}
	}
}

func TestCommand_UnmarshalJSON_Start(t *testing.T) {
	var cmd Command
	err := cmd.UnmarshalJSON("Start")
	if err != nil {
		t.Fatalf("UnmarshalJSON returned error: %v", err)
	}
	if cmd != Start {
		t.Errorf("Expected Start (%d), got %d", Start, cmd)
	}
}

func TestCommand_UnmarshalJSON_Stop(t *testing.T) {
	var cmd Command
	err := cmd.UnmarshalJSON("Stop")
	if err != nil {
		t.Fatalf("UnmarshalJSON returned error: %v", err)
	}
	if cmd != Stop {
		t.Errorf("Expected Stop (%d), got %d", Stop, cmd)
	}
}

func TestCommand_UnmarshalJSON_AllCommands(t *testing.T) {
	tests := []struct {
		input    string
		expected Command
	}{
		{"Start", Start},
		{"Stop", Stop},
		{"Reset", Reset},
		{"ExecuteNext", ExecuteNext},
		{"GetCPUState", GetCPUState},
		{"GetEmulatorState", GetEmulatorState},
		{"GetCIAState", GetCIAState},
		{"GetMemoryContent", GetMemoryContent},
		{"SetBreakpoint", SetBreakpoint},
	}

	for _, tt := range tests {
		var cmd Command
		err := cmd.UnmarshalJSON(tt.input)
		if err != nil {
			t.Fatalf("UnmarshalJSON for '%s' returned error: %v", tt.input, err)
		}
		if cmd != tt.expected {
			t.Errorf("For input '%s': expected %d, got %d", tt.input, tt.expected, cmd)
		}
	}
}

func TestCommand_UnmarshalJSON_UnknownCommand(t *testing.T) {
	var cmd Command
	err := cmd.UnmarshalJSON("UnknownCommand")
	if err != nil {
		t.Fatalf("UnmarshalJSON returned error: %v", err)
	}
	// Unknown command should result in zero value
	if cmd != 0 {
		t.Errorf("Expected 0 for unknown command, got %d", cmd)
	}
}

func TestCommand_String_Roundtrip(t *testing.T) {
	commands := []Command{Start, Stop, Reset, ExecuteNext, GetCPUState, GetEmulatorState, GetCIAState, GetMemoryContent, SetBreakpoint}

	for _, cmd := range commands {
		str := cmd.String()
		var parsed Command
		parsed.UnmarshalJSON(str)
		if parsed != cmd {
			t.Errorf("Roundtrip failed for %v: String()='%s', parsed back to %v", cmd, str, parsed)
		}
	}
}
