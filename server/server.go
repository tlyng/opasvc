package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc/grpclog"

	"github.com/tlyng/opasvc/pb"
)

// Backend ...
type Backend struct {
}

var _ pb.HelloServer = (*Backend)(nil)

// New constructs a new Backend
func New() *Backend {
	return &Backend{}
}

// Say implement the HelloServer interface
func (b *Backend) Say(ctx context.Context, r *pb.Request) (*pb.Response, error) {
	message := fmt.Sprintf("%s %s", r.GetGreeting(), r.GetName())
	grpclog.Infoln("Say called, response", message)
	return &pb.Response{Message: message}, nil
}
