package main

import (
	"context"

	"tailscale.com/client/tailscale"
	"tailscale.com/tsnet"
)

func init() {
	tailscale.I_Acknowledge_This_API_Is_Unstable = true
}

func main() {
	server := &tsnet.Server{}
	ce(server.Start())

	ctx := context.Background()

	conn, err := server.Dial(ctx, "tcp", "100.79.235.80:10000")
	ce(err)
	defer conn.Close()

	select {}

}
