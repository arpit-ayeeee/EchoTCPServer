package main

import (
	"flag"
	"log"

	"github.com/dicedb/dice/config"
	tcp "github.com/dicedb/dice/server"
)

// These are the commmand line flags which we'll give when running the server
func setupFlags() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the dice server") //From which ip's we'll be accepting the connection, everyone
	flag.IntVar(&config.Port, "port", 7379, "port for the dice server")         //our server will be running at this port
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("rolling the dice 🎲")
	tcp.RunSyncTCPServer() //running a synchronous tcp server
}
