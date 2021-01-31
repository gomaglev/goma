package mock

import (
	"GOMA/internal/app/model"
	"github.com/google/wire"
)

// TestInjectorSet inject Injector, only for test
var ModelInjectorSet = wire.NewSet(wire.Struct(new(ModelInjector), "*"))

type ModelInjector struct {
	OrderModel model.IOrder
}
