package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	v1 "github.com/naivary/kubeplate/api/input/v1"
	"github.com/naivary/kubeplate/sdk/inputer"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: inputer.Handshake,
		Plugins: map[string]plugin.Plugin{
			"inputer": &inputer.GRPCPlugin{},
		},
		Cmd: exec.Command("sh", "-c", os.Getenv("INPUTER_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		return err
	}

	raw, err := rpcClient.Dispense("inputer")
	if err != nil {
		return err
	}

	inputer := raw.(inputer.Inputer)

	ctx := context.Background()
	res, err := inputer.Read(ctx, &v1.ReadRequest{
		Path: "vars.json",
	})
	if err != nil {
		return err
	}
	fmt.Println(res.Data)
	return nil
}
