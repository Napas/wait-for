package grpc

import (
	"github.com/Napas/wait-for/runner"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func NewCmd(logger logrus.FieldLogger) *cli.Command {
	return &cli.Command{
		Name:        "grpc",
		Usage:       "grpc command to run with arguments",
		Description: "Waits for grpc request to succeed and then runs given command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "service",
				Aliases:  []string{"s"},
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			waiter := newGrpcWaiter(logger)
			err := waiter.Wait(ctx, grpcWaiterCfg{
				Url:     ctx.String("url"),
				Service: ctx.String("service"),
			})

			if err != nil {
				return err
			}

			return runner.Run(ctx, logger)
		},
	}
}
