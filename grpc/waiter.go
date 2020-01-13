package grpc

import (
	"context"
	"time"

	"google.golang.org/grpc/health/grpc_health_v1"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
)

type grpcWaiter struct {
	logger logrus.FieldLogger
}

func newGrpcWaiter(
	logger logrus.FieldLogger,
) *grpcWaiter {
	return &grpcWaiter{
		logger: logger,
	}
}

func (w *grpcWaiter) Wait(ctx context.Context, cfg grpcWaiterCfg) (err error) {
	var conn *grpc.ClientConn

	for conn, err = grpc.DialContext(ctx, cfg.Url, grpc.WithInsecure()); err != nil; {
		w.logger.Debug(err)
		w.logger.Info("No luck, waiting ...")

		time.Sleep(time.Second)
	}

	client := grpc_health_v1.NewHealthClient(conn)

	for !w.isServing(ctx, client, cfg.Service) {
		w.logger.Info("No luck, waiting ...")

		time.Sleep(time.Second)
	}

	return nil
}

func (w *grpcWaiter) isServing(ctx context.Context, client grpc_health_v1.HealthClient, service string) bool {
	w.logger.Infof("Checking if %s is up", service)

	resp, err := client.Check(ctx, &grpc_health_v1.HealthCheckRequest{
		Service: service,
	})

	if err != nil {
		w.logger.Debug(err)

		return false
	}

	return resp.Status == grpc_health_v1.HealthCheckResponse_SERVING
}

type grpcWaiterCfg struct {
	Url     string
	Service string
}
