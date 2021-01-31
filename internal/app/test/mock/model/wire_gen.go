// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package mock

import (
	"GOMA/internal/app/model/gorm"
	"GOMA/internal/app/model/gorm/model"
)

// Injectors from wire.go:

func BuildModelInjector() (*ModelInjector, func(), error) {
	db, cleanup, err := gorm.InitGormDB()
	if err != nil {
		return nil, nil, err
	}
	order := &model.Order{
		DB: db,
	}
	modelInjector := &ModelInjector{
		OrderModel: order,
	}
	return modelInjector, func() {
		cleanup()
	}, nil
}
