package main

import (
	"net/http"
	"os"

	appHtpp "github.com/Napas/wait-for/http"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	cli.VersionFlag = &cli.StringFlag{
		Name:    "version",
		Aliases: []string{"V"},
	}

	app := &cli.App{
		Name: "Wait for",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
			},
		},
		Before: func(ctx *cli.Context) error {
			if ctx.Bool("verbose") {
				logger.SetLevel(logrus.DebugLevel)
			}

			return nil
		},
		Commands: []*cli.Command{
			appHtpp.NewCmd(http.DefaultClient, logger),
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
