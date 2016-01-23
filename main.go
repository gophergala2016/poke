package main

import (
	"flag"
)

var (
	restHost = flag.String("b", "0.0.0.0", "start REST interface on HOST")
	restPort = flag.Int("p", 8080, "use PORT for REST interface")
	wsHost   = flag.String("s", "0.0.0.0", "start websocket  on HOST")
	wsPort   = flag.Int("r", 9090, "use PORT for websocket interface")
)



func main() {

	flag.Parse()

	go StartRestInterface(*restHost, *restPort)

	StartWebsocketInterface(*wsHost, *wsPort)
}
