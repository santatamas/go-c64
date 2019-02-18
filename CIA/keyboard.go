package CIA

import (
	"github.com/gdamore/tcell"
)

type Keyboard struct {
	Cia *CIA
}

func (keyboard *Keyboard) PressKey(key *tcell.EventKey) {

	if key.Rune() == 'a' {
		keyboard.Cia.SetKey(0x1)
	}
}

func NewKeyboard(cia *CIA) Keyboard {
	return Keyboard{
		Cia: cia,
	}
}
