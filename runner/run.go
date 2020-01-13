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
	stdOutReader, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	stdErrReader, err := cmd.StderrPipe()

	if err != nil {
		return err
	}

	stdOutScanner := bufio.NewScanner(stdOutReader)

	go func() {
		for stdOutScanner.Scan() {
			logger.Println(stdOutScanner.Text())
		}
	}()

	stdErrScanner := bufio.NewScanner(stdErrReader)

	go func() {
		for stdErrScanner.Scan() {
			logger.Println(stdErrScanner.Text())
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
