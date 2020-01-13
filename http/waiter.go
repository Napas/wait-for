package http

import (
	"net/http"
	url "net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HttpWaiter struct {
	httpClient HttpClient
	debug      bool
	logger     logrus.FieldLogger
}

func newHttpWaiter(
	httpClient HttpClient,
	debug bool,
	logger logrus.FieldLogger,
) *HttpWaiter {
	return &HttpWaiter{
		httpClient: httpClient,
		debug:      debug,
		logger:     logger,
	}
}

func (w *HttpWaiter) Wait(cfg httpWaiterCfg) error {
	reqUrl, err := url.Parse(cfg.Url)

	if err != nil {
		return err
	}

	for err := w.doRequest(reqUrl, cfg.ExpectedHttpCode, cfg.Method); err != nil; {
		if w.debug {
			w.logger.Error(err)
		}

		w.logger.Info("No luck, waiting ...")

		time.Sleep(time.Second)
	}

	return nil
}

func (w *HttpWaiter) doRequest(reqUrl *url.URL, expectedStatusCode int, method string) error {
	w.logger.Infof("Making request %s %s", method, reqUrl.String())

	resp, err := w.httpClient.Do(&http.Request{
		Method: method,
		URL:    reqUrl,
	})

	if err != nil {
		return err
	}

	if resp.StatusCode == expectedStatusCode {
		return nil
	}

	return errors.New("unexpected status code")
}

type httpWaiterCfg struct {
	Url              string
	Method           string
	ExpectedHttpCode int
}

func (cfg *httpWaiterCfg) GetMethod() string {
	if cfg.Method != "" {
		return cfg.Method
	}

	return http.MethodGet
}

func (cfg *httpWaiterCfg) GetExpectedStatusCode() int {
	if cfg.ExpectedHttpCode == 0 {
		return cfg.ExpectedHttpCode
	}

	return http.StatusOK
}
