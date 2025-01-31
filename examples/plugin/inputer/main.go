package main

import (
	"context"
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	v1 "github.com/naivary/kubeplate/api/input/v1"
	"github.com/naivary/kubeplate/sdk/inputer"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: inputer.Handshake,
		Plugins: map[string]plugin.Plugin{
			"inputer": &inputer.GRPCPlugin{Impl: &jsonInputer{}},
		},
		Logger: logger,
		GRPCServer: plugin.DefaultGRPCServer,
	})
	return nil
}

var _ inputer.Inputer = (*jsonInputer)(nil)

type jsonInputer struct{}

func (j *jsonInputer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	res := &v1.ReadResponse{
		Data: map[string]*anypb.Any{
			"name": {Value: []byte("kubeplate")},
		},
	}
	return res, nil
}
