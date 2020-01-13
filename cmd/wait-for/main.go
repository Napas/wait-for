package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Napas/wait-for/grpc"

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
			grpc.NewCmd(logger),
		},
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
