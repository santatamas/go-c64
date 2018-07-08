package MOS6510

import (
	"testing"

	"github.com/santatamas/go-c64/RAM"
)

func TestCPU_stackPush(t *testing.T) {
	type fields struct {
		memory  *RAM.Memory
		A       byte
		Y       byte
		X       byte
		S       byte
		P       byte
		PC      uint16
		SP      byte
		SP_LOW  uint16
		SP_HIGH uint16
	}
	type args struct {
		value byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CPU{
				memory:  tt.fields.memory,
				A:       tt.fields.A,
				Y:       tt.fields.Y,
				X:       tt.fields.X,
				S:       tt.fields.S,
				P:       tt.fields.P,
				PC:      tt.fields.PC,
				SP:      tt.fields.SP,
				SP_LOW:  tt.fields.SP_LOW,
				SP_HIGH: tt.fields.SP_HIGH,
			}
			c.stackPush(tt.args.value)
		})
	}
}
