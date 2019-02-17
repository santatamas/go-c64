package CIA

import (
	"github.com/gdamore/tcell"
)

type Keyboard struct {
	Cia *CIA
}

func (Keyboard *Keyboard) PressKey(key tcell.Key) {

}

func NewKeyboard(cia *CIA) Keyboard {
	return Keyboard{
		Cia: cia,
	}
}
