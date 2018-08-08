package main

import (
	"flag"
	"github.com/santatamas/go-c64/internals"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "Internal debug server address.")

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
	d.hub.Broadcast <- decodedMessage
	return len(p), nil
}
