package main

import (
	"context"
	"os"

	"app_module__/internal/app"
	"app_module__/pkg/logger"

	"github.com/urfave/cli/v2"
)

// go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "0.0.1"

func main() {
	app := cli.NewApp()

	app.Name = "AppNameTitleCase__"
	app.Version = VERSION
	app.Usage = "AppNameTitleCase__ gRPC microservice"

	logger.SetVersion(VERSION)
	ctx := logger.NewTraceIDContext(context.Background(), "main")
	app.Commands = []*cli.Command{
		newServiceCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.Errorf(ctx, "%s", err)
	}
}

func newServiceCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "app_name__ start",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "config file(.json,.yaml,.toml)",
				Required: false,
				Value:    "./configs/config.toml",
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")),
				app.SetVersion(VERSION))
		},
	}
}
