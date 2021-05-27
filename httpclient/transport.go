package httpclient

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/airbloc/logger"
	"github.com/pkg/errors"
)

type logTransport struct {
	transport http.RoundTripper
	schemes   map[string]bool
	logger    logger.Logger
}

func (t logTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if ok := t.schemes[req.URL.Scheme]; !ok {
		return nil, errors.Errorf("unacceptable scheme %s", req.URL.Scheme)
	}

	timer := time.Now()
	resp, err := t.transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}
	t.logger.Debug(
		"[{}] request to {} with method {} ({})",
		resp.StatusCode, req.URL.String(), req.Method, time.Since(timer),
	)

	if resp.StatusCode >= 400 {
		var buf bytes.Buffer
		resp.Body = ioutil.NopCloser(io.TeeReader(resp.Body, &buf))

		rawBody, err := ioutil.ReadAll(resp.Body)
		resp.Body = ioutil.NopCloser(&buf)
		if err != nil {
			t.logger.Error("failed to read response body. err={}", err)
		} else {
			return nil, errors.New(string(rawBody))
		}
	}

	return resp, err
}
