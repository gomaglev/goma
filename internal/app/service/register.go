package service

import (
	"context"
	"fmt"

	gopackage_name__v1 "app_module__/internal/app/service/v1/gopackage_name__"
	"app_module__/internal/pkg/config"

	"app_module__/internal/pkg/server/rpc"

	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var _ rpc.IRegister = (*Register)(nil)

// RegisterSet injection
var RegisterSet = wire.NewSet(wire.Struct(new(Register), "*"), wire.Bind(new(rpc.IRegister), new(*Register)))

type Register struct {
	PbName__ServiceV1 *gopackage_name__v1.PbName__Service
}

// RegisterServiceServers implementation
func (r *Register) RegisterServiceServers(server *grpc.Server) {
	// gopackage_name__ Service
	gopackage_name__v1.RegisterPbName__ServiceServer(server, r.PbName__ServiceV1)
}

// RegisterServiceHandlerFromEndpoints implementation
func (r *Register) RegisterServiceHandlerFromEndpoints(ctx context.Context, multiplexer *runtime.ServeMux) error {
	endpoint := r.endpoint()
	dialOption := r.DialOption()
	// gopackage_name__ Service
	if err := gopackage_name__v1.RegisterPbName__ServiceHandlerFromEndpoint(
		ctx, multiplexer, endpoint, dialOption); err != nil {
		return err
	}
	return nil
}

func (r *Register) endpoint() string {
	return fmt.Sprintf("%s:%d", config.C.Gateway.Host, config.C.Gateway.Port)
}

func (r *Register) DialOption() []grpc.DialOption {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	return opts
}
