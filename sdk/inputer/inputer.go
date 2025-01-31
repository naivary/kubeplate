package inputer

import (
	"context"

	v1 "github.com/naivary/kubeplate/api/input/v1"
)

type Inputer interface {
	Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error)
}

