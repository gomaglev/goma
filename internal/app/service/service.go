package service

import (
	gopackage_name__v1 "app_module__/internal/app/service/v1/gopackage_name__"

	"github.com/google/wire"
)

// ServiceSet api injection
var ServiceSet = wire.NewSet(
	RegisterSet,
	gopackage_name__v1.PbName__Set,
)
