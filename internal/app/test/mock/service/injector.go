package mock

import (
	orderv1 "GOMA/internal/app/service/v1/order"
	orderv2 "GOMA/internal/app/service/v2/order"
	"github.com/google/wire"
)

// MockInjectorSet inject Injector, only for test
var MockInjectorSet = wire.NewSet(wire.Struct(new(MockInjector), "*"))

type MockInjector struct {
	OrderServiceV1 *orderv1.OrderService
	OrderServiceV2 *orderv2.OrderService
}
