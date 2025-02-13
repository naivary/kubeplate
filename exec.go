package kubeplate

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hashicorp/go-plugin"
	v1 "github.com/naivary/kubeplate/api/inputer/v1"
	"github.com/naivary/kubeplate/sdk/inputer"
	"github.com/naivary/kubeplate/sdk/outputer"
)

const kubeplate = ".kubeplate"

func LoadVars(inputerPath, varsURL string) (*v1.ReadResponse, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	binary := filepath.Join(home, kubeplate, inputerPath)
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: inputer.Handshake,
		Plugins: map[string]plugin.Plugin{
			"inputer": &inputer.GRPCPlugin{},
		},
		Cmd: exec.Command("sh", "-c", binary),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC,
		},
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}
	raw, err := rpcClient.Dispense("inputer")
	if err != nil {
		return nil, err
	}

	inputer := raw.(inputer.Inputer)
	ctx := context.Background()
	res, err := inputer.Read(ctx, &v1.ReadRequest{
		Url: varsURL,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(res.Data)
	return res, nil
}

func OutputTo(out outputer.Outputer, data io.Reader) error {
	return nil
}
