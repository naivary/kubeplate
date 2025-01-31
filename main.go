package main

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
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
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}
	raw, err := rpcClient.Dispense("inputer")
	if err != nil {
		return err
	}
	impl := raw.(inputer.Inputer)
	req := &v1.ReadRequest{
		Path: "vars.json",
	}
	res, err := impl.Read(context.Background(), req)
	if err != nil {
		return err
	}
	logger := hclog.New(hclog.DefaultOptions)
	logger.Info("got the res", "res", res.Data)
	client.Kill()
	return nil
}
