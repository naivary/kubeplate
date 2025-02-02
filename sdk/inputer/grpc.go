package inputer

import (
	"context"

	"github.com/hashicorp/go-plugin"
	v1 "github.com/naivary/kubeplate/api/inputer/v1"
	"google.golang.org/grpc"
)

type GRPCPlugin struct {
	plugin.Plugin

	Impl Inputer
}

func (p *GRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	v1.RegisterInputerServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *GRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (any, error) {
	return &GRPCClient{client: v1.NewInputerClient(c)}, nil
}

type GRPCClient struct{ client v1.InputerClient }

func (m *GRPCClient) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	return m.client.Read(ctx, req)

}

type GRPCServer struct {
	v1.UnimplementedInputerServer

	Impl Inputer
}

func (m *GRPCServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	return m.Impl.Read(ctx, req)
}
