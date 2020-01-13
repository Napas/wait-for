package http

import (
	"github.com/Napas/wait-for/runner"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func NewCmd(httpClient HttpClient, logger logrus.FieldLogger) *cli.Command {
	return &cli.Command{
		Name:        "http",
		Usage:       "http command to run with arguments",
		Description: "Waits for http request to succeed and then runs given command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "method",
				Aliases: []string{"m"},
				Value:   "GET",
			},
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Required: true,
			},
			&cli.IntFlag{
				Name:    "expected-status-code",
				Aliases: []string{"e"},
				Value:   200,
			},
		},
		Action: func(ctx *cli.Context) error {
			httpWaiter := newHttpWaiter(
				httpClient,
				ctx.Bool("verbose"),
				logger,
			)
			cfg := httpWaiterCfg{
				Url:              ctx.String("url"),
				Method:           ctx.String("method"),
				ExpectedHttpCode: ctx.Int("expected-status-code"),
			}

			if err := httpWaiter.Wait(ctx.Context, cfg); err != nil {
				return err
			}

			return runner.Run(ctx, logger)
		},
	}
}
