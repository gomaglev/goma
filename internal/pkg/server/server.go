package server

import (
	"app_module__/internal/pkg/server/rpc"

	"github.com/google/wire"
)

// ServerSet inject grpc & gateway server
var ServerSet = wire.NewSet(
	rpc.ServerSet,
)
