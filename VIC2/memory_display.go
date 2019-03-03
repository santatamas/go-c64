package VIC2

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/santatamas/go-c64/RAM"
	"os"
	"time"
)

type MemoryDisplay struct {
	memory        *RAM.Memory
	c64Characters func(byte) rune
	keyPress      chan *tcell.EventKey
}

const SCREEN_START_ADDR = uint16(0x400)
const COLOR_START_ADDR = uint16(0xD800)
const BACKGROUND_COLOR_ADDR = 0xD021
const FRAME_COLOR_ADDR = 0xD020

const RASTER_REGISTER_ADDR = 0xD012

const TEXT_SCREEN_WIDTH = 40
const TEXT_SCREEN_HEIGHT = 25

func NewMemoryDisplay(memory *RAM.Memory, keyPress chan *tcell.EventKey) MemoryDisplay {
	result := MemoryDisplay{memory, asciiCharacters(), keyPress}
	return result
}

func (display *MemoryDisplay) ReadCurrentState() [40][25]rune {
	result := [40][25]rune{}
	currentAdr := SCREEN_START_ADDR
	for y := 0; y < TEXT_SCREEN_HEIGHT; y++ {
		for x := 0; x < TEXT_SCREEN_WIDTH; x++ {
			result[x][y] = display.c64Characters(display.memory.ReadAbsolute(currentAdr))
			currentAdr++
		}
	}
	return result
}

func (display *MemoryDisplay) DrawState(screen tcell.Screen) {
	st := tcell.StyleDefault
	backgroundColor := display.memory.ReadAbsolute(BACKGROUND_COLOR_ADDR)
	//st = st.Background(tcell.ColorWhite)
	st = st.Background(getColor(backgroundColor))

	frameColor := display.memory.ReadAbsolute(FRAME_COLOR_ADDR)
	frameStyle := st.Background(getColor(frameColor))

	for y := 0; y < TEXT_SCREEN_HEIGHT+2; y++ {
		for x := 0; x < TEXT_SCREEN_WIDTH+2; x++ {
			screen.SetCell(x, y, frameStyle, ' ')
		}
	}

	currentAdr := SCREEN_START_ADDR
	currentColorAdr := COLOR_START_ADDR

	for y := 0; y < TEXT_SCREEN_HEIGHT; y++ {
		for x := 0; x < TEXT_SCREEN_WIDTH; x++ {

			color := display.memory.ReadAbsolute(currentColorAdr)
			char := display.c64Characters(display.memory.ReadAbsolute(currentAdr))

			st = st.Foreground(getColor(color))
			if char == ' ' {
				//st = st.Background(getColor(color))
			}
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

				display.keyPress <- ev

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
		//rasterReg := display.memory.ReadAbsolute(RASTER_REGISTER_ADDR)
		//rasterReg++
		display.memory.WriteAbsolute(RASTER_REGISTER_ADDR, 0)
		time.Sleep(100 * time.Millisecond)
		//log.Println("screen refresh")
	}

	s.Fini()
}
