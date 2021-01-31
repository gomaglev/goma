package service

import (
	orderv1 "GOMA/internal/app/service/v1/order"
	itemv1 "GOMA/internal/app/service/v1/order/item"
	orderv2 "GOMA/internal/app/service/v2/order"

	"github.com/google/wire"
)

// ServiceSet api injection
var ServiceSet = wire.NewSet(
	RegisterSet,
	orderv1.OrderSet,
	itemv1.ItemSet,
	orderv2.OrderSet,
)
