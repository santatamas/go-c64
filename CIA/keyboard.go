package CIA

import (
	"github.com/gdamore/tcell"
	"log"
	"strconv"
)

type Pair struct {
	row, col byte
}

/*
WRITE TO PORT A               READ PORT B (56321, $DC01)
56320/$DC00
         Bit 7   Bit 6   Bit 5   Bit 4   Bit 3   Bit 2   Bit 1   Bit 0
Bit 7    STOP    Q       C=      SPACE   2       CTRL    <-      1
Bit 6    /       ^       =       RSHIFT  HOME    ;       *       LIRA
Bit 5    ,       @       :       .       -       L       P       +
Bit 4    N       O       K       M       0       J       I       9
Bit 3    V       U       H       B       8       G       Y       7
Bit 2    X       T       F       C       6       D       R       5
Bit 1    LSHIFT  E       S       Z       4       A       W       3
Bit 0    CRSR DN F5      F3      F1      F7      CRSR RT RETURN  DELETE
*/

func keymap() func(rune) Pair {

	innerMap := map[rune]Pair{
		'a': Pair{1, 2},
		'b': Pair{3, 4},
		'c': Pair{2, 4},
		'd': Pair{2, 2},
		'e': Pair{1, 6},
		'f': Pair{2, 5},
		'g': Pair{3, 2},
		'h': Pair{3, 5},
		'i': Pair{4, 1},
		'j': Pair{4, 2},
		'k': Pair{4, 5},
		'l': Pair{5, 2},
		'm': Pair{4, 4},
		'n': Pair{4, 7},
		'o': Pair{4, 6},
		'p': Pair{5, 1},
		'q': Pair{7, 6},
		'r': Pair{2, 1},
		's': Pair{1, 5},
		't': Pair{2, 6},
		'u': Pair{3, 6},
		'v': Pair{3, 7},
		'w': Pair{1, 1},
		'x': Pair{2, 7},
		'y': Pair{3, 1},
		'z': Pair{1, 4},
	}

	return func(key rune) Pair {
		log.Println("[Keyboard] Resolving rune " + string(key) + " to: " + strconv.Itoa(int(innerMap[key].row)) + "," + strconv.Itoa(int(innerMap[key].col)))
		return innerMap[key]
	}
}

type Keyboard struct {
	Cia *CIA
}

func (keyboard *Keyboard) PressKey(key *tcell.EventKey) {

	r := key.Rune()
	log.Println("[Keyboard] PressKey called with rune:" + string(r))
	keyMask := keymap()(r)

	keyboard.Cia.SetKey(keyMask.row, keyMask.col)
}

func NewKeyboard(cia *CIA) Keyboard {
	return Keyboard{cia}
}
