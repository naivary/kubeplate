package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	v1 "github.com/naivary/kubeplate/api/inputer/v1"
	"github.com/naivary/kubeplate/sdk/inputer"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "json_inputer",
		JSONFormat: true,
		Output:     os.Stderr,
	})
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: inputer.Handshake,
		Plugins: map[string]plugin.Plugin{
			"inputer": &inputer.GRPCPlugin{Impl: &jsonInputer{l: logger}},
		},
		Logger:     logger,
		GRPCServer: plugin.DefaultGRPCServer,
	})
	return nil
}

var _ inputer.Inputer = (*jsonInputer)(nil)

type jsonInputer struct {
	l hclog.Logger
}

func (j *jsonInputer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	j.l.Info(req.Url)
	data, err := os.ReadFile("./examples/plugin/inputer/vars.json")
	if err != nil {
		return nil, err
	}
	anymap := make(map[string]any)
	if err := json.Unmarshal(data, &anymap); err != nil {
		return nil, err
	}
	str, err := structpb.NewValue(anymap)
	if err != nil {
		return nil, err
	}
	j.l.Error("this is a error log from the plugin")
	j.l.Info("this is a info log from the plugin")
	return &v1.ReadResponse{
		Data: map[string]*structpb.Struct{
			"vars.json": str.GetStructValue(),
		},
	}, nil
}
