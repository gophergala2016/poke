package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	host = flag.String("b", "0.0.0.0", "listen on HOST")
	port = flag.Int("p", 8080, "use PORT")
)

type Response struct {
	Status string `json:"status"`
}


func ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func RestHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{"OK"}
	b, err := json.Marshal(resp)

	if err != nil {
		ErrorResponse(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {

	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	log.Println("starting server on", addr)

	http.HandleFunc("/", RestHandler)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Println("ERROR:", err)
	}

}
