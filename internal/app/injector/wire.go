// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	gorm "app_module__/internal/app/model/gorm"
	gormModel "app_module__/internal/app/model/gorm/model"
	"app_module__/internal/app/service"
	"app_module__/internal/pkg/server"

	"github.com/google/wire"
)

// BuildInjector
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		// redis.InitCache,
		// mongo.InitMongo,
		// mongoModel.ModelSet,
		gorm.InitGormDB,
		gormModel.ModelSet,
		service.ServiceSet,
		server.ServerSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
