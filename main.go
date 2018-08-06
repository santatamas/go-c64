package main

import (
	"flag"
	"github.com/santatamas/go-c64/MOS6510"
	"github.com/santatamas/go-c64/VIC2"
	"github.com/santatamas/go-c64/internals"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "internals/debug.html")
}

func startDebugServer() (hub *internals.Hub) {
	hub = internals.NewHub()
	go hub.Run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		internals.ServeWs(hub, w, r)
	})
	go http.ListenAndServe(*addr, nil)
	return hub
}

type DebugLog struct {
	hub *internals.Hub
}

func (d DebugLog) Write(p []byte) (n int, err error) {
	decodedMessage := string(p)
	//decodedMessage = string(len(decodedMessage))
	d.hub.Broadcast <- decodedMessage
	return len(p), nil
}

func main() {

	flag.Parse()
	//file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	//log.SetOutput(file)
	hub := startDebugServer()
	debugServerLogger := DebugLog{hub: hub}
	log.SetOutput(debugServerLogger)
	time.Sleep(5 * time.Second)

	//memory, startPCH, startPCL := loadFile("./_resources/Phase2/5_addressingmodes.prg")

	memory, startPCH, startPCL := loadFile("./_resources/Phase1/Prg/4_Colors.prg")
	cpu := MOS6510.NewCPU(&memory)
	display := VIC2.NewMemoryDisplay(&memory)

	go cpu.Start(startPCL, startPCH)
	display.Start()

}
