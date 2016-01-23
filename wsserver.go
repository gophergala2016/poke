package main

import (
	"net/http"
	"golang.org/x/net/websocket"
	"log"
	"fmt"
)

func wsHandler(ws *websocket.Conn) {
	log.Println("new connection")
}

func StartWebsocketInterface(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("starting websocket interfce...", addr)
	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(wsHandler))
	http.ListenAndServe(addr, mux)
}
