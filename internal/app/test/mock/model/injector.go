package mock

import (
	"app_module__/internal/app/model"

	"github.com/google/wire"
)

// TestInjectorSet inject Injector, only for test
var ModelInjectorSet = wire.NewSet(wire.Struct(new(ModelInjector), "*"))

type ModelInjector struct {
	MessageTypeName__Model model.IMessageTypeName__
}
