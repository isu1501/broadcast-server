package main

import (
	"fmt"
	"os"
	"broadcast-server/server"
	"broadcast-server/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("broadcast-server <start|connect>")
		return
	}

	switch os.Args[1] {
	case "start":
		server.StartServer()
	case "connect":
		client.StartClient()
	default:
		fmt.Println("Invalid command. Use 'start' or 'connect'")
	}
}
