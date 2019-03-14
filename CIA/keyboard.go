package CIA

import (
	"github.com/gdamore/tcell"
	"log"
	"strconv"
	//"time"
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
		' ': Pair{7, 4},
		';': Pair{6, 2},
		':': Pair{5, 5},

		'1': Pair{7, 0},
		'2': Pair{7, 3},
		'3': Pair{1, 0},
		'4': Pair{1, 3},
		'5': Pair{2, 0},
		'6': Pair{2, 3},
		'7': Pair{3, 0},
		'8': Pair{3, 3},
		'9': Pair{4, 0},
		'0': Pair{4, 3},

		'"': Pair{4, 3},
	}

	return func(key rune) Pair {
		log.Println("[Keyboard] Resolving rune " + string(key) + " to: " + strconv.Itoa(int(innerMap[key].row)) + "," + strconv.Itoa(int(innerMap[key].col)))

		return innerMap[key]
	}
}

type Keyboard struct {
	Cia         *CIA
	previousRow []byte
	previousCol []byte
}

func (keyboard *Keyboard) PressKey(key *tcell.EventKey) {

	for i := 0; i < 2; i++ {
		keyboard.Cia.UnsetKey(keyboard.previousRow[i], keyboard.previousCol[i])
	}

	//keyboard.Cia.SendInterrupt()

	//time.Sleep(100)
	if key.Key() == tcell.KeyEnter {
		keyboard.Cia.SetKey(0, 1, true)

		keyboard.previousRow[0] = 0
		keyboard.previousCol[0] = 1
	} else if key.Rune() == '"' {
		keyboard.Cia.SetKey(1, 7, false)
		keyboard.Cia.SetKey(7, 3, true)

		keyboard.previousRow[0] = 1
		keyboard.previousCol[0] = 7

		keyboard.previousRow[1] = 7
		keyboard.previousCol[1] = 3

	} else {
		r := key.Rune()

		log.Println("[Keyboard] PressKey called with rune:" + string(r))
		keyMask := keymap()(r)

		keyboard.Cia.SetKey(keyMask.row, keyMask.col, true)

		keyboard.previousRow[0] = keyMask.row
		keyboard.previousCol[0] = keyMask.col
	}

}

func NewKeyboard(cia *CIA) Keyboard {
	return Keyboard{cia, make([]byte, 2), make([]byte, 2)}
}
