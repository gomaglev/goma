// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package mock

import (
	"GOMA/internal/app/model/gorm"
	"GOMA/internal/app/model/gorm/model"
	"GOMA/internal/app/service"
	"github.com/google/wire"
)

// BuildMockInjector for tests in api
func BuildMockInjector() (*MockInjector, func(), error) {
	wire.Build(
		gorm.InitGormDB,
		model.ModelSet,
		service.ServiceSet,
		MockInjectorSet,
	)
	return new(MockInjector), nil, nil
}
