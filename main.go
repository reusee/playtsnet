package main

import (
	"tailscale.com/tsnet"
)

func main() {
	server := &tsnet.Server{}
	ce(server.Start())
}
