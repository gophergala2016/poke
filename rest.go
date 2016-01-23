package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)
type Response struct {
	Status string `json:"status"`
}


func errorResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func restHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{"OK"}
	b, err := json.Marshal(resp)

	if err != nil {
		errorResponse(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func StartRestInterface(host string, port int) {

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Println("starting rest interface..", addr)

	mux := http.NewServeMux()
	mux.HandleFunc("/",restHandler)
	http.ListenAndServe(addr, mux)
}
