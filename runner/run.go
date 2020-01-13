package runner

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Run(ctx *cli.Context, logger logrus.FieldLogger) error {
	logger.Infof("Running %s %s", ctx.Args().First(), strings.Join(ctx.Args().Tail(), " "))
	cmd := exec.CommandContext(ctx.Context, ctx.Args().First(), ctx.Args().Tail()...)
	reader, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(reader)

	go func() {
		for scanner.Scan() {
			logger.Println(scanner.Text())
		}
	}()

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
