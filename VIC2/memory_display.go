package VIC2

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/santatamas/go-c64/RAM"
	"os"
	"time"
)

type MemoryDisplay struct {
	screenStartAddress     uint16
	colorStartAddress      uint16
	backgroundColorAddress uint16
	frameColorAddress      uint16
	memory                 *RAM.Memory
	width                  int
	height                 int
	c64Characters          func(byte) rune
}

func NewMemoryDisplay(memory *RAM.Memory) MemoryDisplay {
	result := MemoryDisplay{0x400, 0xD800, 0xD021, 0xD020, memory, 40, 25, asciiCharacters()}
	return result
}

func (display *MemoryDisplay) ReadCurrentState() [40][25]rune {
	result := [40][25]rune{}
	currentAdr := display.screenStartAddress
	for y := 0; y < display.height; y++ {
		for x := 0; x < display.width; x++ {
			result[x][y] = display.c64Characters(display.memory.ReadAbsolute(currentAdr))
			currentAdr++
		}
	}
	return result
}

func (display *MemoryDisplay) DrawState(screen tcell.Screen) {
	st := tcell.StyleDefault
	backgroundColor := display.memory.ReadAbsolute(display.backgroundColorAddress)
	st = st.Background(getColor(backgroundColor))

	frameColor := display.memory.ReadAbsolute(0xD020)
	frameStyle := st.Background(getColor(frameColor))

	for y := 0; y < display.height+2; y++ {
		for x := 0; x < display.width+2; x++ {
			screen.SetCell(x, y, frameStyle, ' ')
		}
	}

	currentAdr := display.screenStartAddress
	currentColorAdr := display.colorStartAddress

	for y := 0; y < display.height; y++ {
		for x := 0; x < display.width; x++ {

			color := display.memory.ReadAbsolute(currentColorAdr)
			char := display.c64Characters(display.memory.ReadAbsolute(currentAdr))

			st = st.Foreground(getColor(color))
			screen.SetCell(x+1, y+1, st, char)
			currentAdr++
			currentColorAdr++
		}
	}

	screen.Show()
}

func (display *MemoryDisplay) Start() {

	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	s.Clear()

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

loop:
	for {
		select {
		case <-quit:
			break loop
		case <-time.After(time.Millisecond * 50):
		}

		display.DrawState(s)
	}

	s.Fini()
}
