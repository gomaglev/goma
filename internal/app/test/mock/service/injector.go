package mock

import (
	gopackage_name__v1 "app_module__/internal/app/service/v1/gopackage_name__"

	"github.com/google/wire"
)

// MockInjectorSet inject Injector, only for test
var MockInjectorSet = wire.NewSet(wire.Struct(new(MockInjector), "*"))

type MockInjector struct {
	OrderServiceV1 *gopackage_name__v1.PbName__Service
}
