// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package injector

import (
	gorm "GOMA/internal/app/model/gorm"
	gormModel "GOMA/internal/app/model/gorm/model"
	"GOMA/internal/app/service"
	"GOMA/internal/pkg/server"

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
