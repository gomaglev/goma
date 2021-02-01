package iutil

import (
	"context"

	"app_module__/internal/pkg/config"
	"app_module__/pkg/logger"
	"app_module__/pkg/trace"
	"app_module__/pkg/unique"

	"github.com/teris-io/shortid"
)

const (
	configFile = "../../../../../configs/config.toml"
)

var idFunc = func() string {
	return unique.NewSnowflakeID().String()
}
var shortIDFunc = func() string {
	return ""
}

// InitID ...
func InitID() {
	switch config.C.UniqueID.Type {
	case "uuid":
		idFunc = func() string {
			return unique.MustUUID().String()
		}
	case "object":
		idFunc = func() string {
			return unique.NewObjectID().Hex()
		}
	default:
		// Initialize snowflake node
		err := unique.SetSnowflakeNode(config.C.UniqueID.Snowflake.Node, config.C.UniqueID.Snowflake.Epoch)
		if err != nil {
			panic(err)
		}

		logger.SetTraceIDFunc(func() string {
			return unique.NewSnowflakeID().String()
		})

		trace.SetIDFunc(func() string {
			return unique.NewSnowflakeID().String()
		})

		idFunc = func() string {
			return unique.NewSnowflakeID().String()
		}
	}
}

// InitShortID ...
func InitShortID() {
	// 注入Shortid
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err)
	}
	shortIDFunc = func() string {
		return sid.MustGenerate()
	}
}

// NewID Create unique id
func NewID() string {
	return idFunc()
}

// NewID Create unique id
func NewShortID() string {
	return shortIDFunc()
}

func InitConfig() context.Context {
	config.MustLoad(configFile)
	config.C.RunMode = "test"
	config.C.Log.Level = 2
	config.C.Gorm.Debug = false
	ctx := logger.NewTraceIDContext(context.Background(), "test")
	return ctx
}
