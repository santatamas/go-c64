package internals

import (
	"flag"
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
	http.ServeFile(w, r, "WebUI/debug.html")
}

func StartDebugServer() (hub *Hub) {
	hub = NewHub()
	go hub.Run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})
	go http.ListenAndServe(*addr, nil)
	log.Println("[DEBUG] Debug server is listening")
	return hub
}

type DebugLog struct {
	Hub *Hub
}

func (d DebugLog) Write(p []byte) (n int, err error) {
	decodedMessage := string(p)
	d.Hub.Broadcast <- decodedMessage
	return len(p), nil
}
